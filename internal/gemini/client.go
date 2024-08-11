package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/time/rate"
	"google.golang.org/api/option"
)

type FileRequest struct {
	Name      string
	Size      int64
	Path      string
	StorageID int
	IsStorage bool
	PathID    int
}

type FileResponse struct {
	Files []struct {
		FullPath         string   `json:"full_path"`
		NewFolders       []string `json:"new_folders"`
		OriginalFilename string   `json:"original_filename"`
	} `json:"files"`
}

func NewFileRequest(name string, size int64, path string, storageID int, isStorage bool, pathID int) FileRequest {
	return FileRequest{
		Name:      name,
		Size:      size,
		Path:      path,
		StorageID: storageID,
		IsStorage: isStorage,
		PathID:    pathID,
	}
}

type ClientQueue struct {
	ctx context.Context

	rpmLimiter *rate.Limiter
	rpdLimiter *rate.Limiter

	client *genai.Client
	model  *genai.GenerativeModel
}

func NewClientQueue(apiKey string) *ClientQueue {
	requestsPerMinute := 15
	requestsPerDay := 1_500

	rpmLimiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(requestsPerMinute)), requestsPerMinute)
	rpdLimiter := rate.NewLimiter(rate.Every(24*time.Hour/time.Duration(requestsPerDay)), requestsPerDay)

	ctx := context.Background()
	clientOption := option.WithAPIKey(apiKey)

	client, err := genai.NewClient(ctx, clientOption)
	if err != nil {
		log.Fatalf("Error creating client: %v\n", err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(0.5)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(16 * 1024)
	model.ResponseMIMEType = "application/json"
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	return &ClientQueue{
		ctx:        ctx,
		rpmLimiter: rpmLimiter,
		rpdLimiter: rpdLimiter,
		client:     client,
		model:      model,
	}
}

func (cq *ClientQueue) MakeRequest(fileRequests []FileRequest, storageName string, existingFolders []string) (FileResponse, error) {
	parts := make([]genai.Part, 0)
	generatedPrompt := generatePrompt(storageName, existingFolders)
	parts = append(parts, genai.Text(generatedPrompt))
	for _, fileRequest := range fileRequests {
		payload, err := extractContent(fileRequest.Path, fileRequest.Name)
		if err != nil {
			log.Printf("Error extracting content: %v\n", err)
			continue
		}
		parts = append(parts, genai.Text(payload))
	}

	ctx := context.Background()
	if err := cq.rpmLimiter.Wait(ctx); err != nil {
		log.Printf("RPM limit exceeded: %v\n", err)
		return FileResponse{}, err
	}
	if err := cq.rpdLimiter.Wait(ctx); err != nil {
		log.Printf("RPD limit exceeded: %v\n", err)
		return FileResponse{}, err
	}

	resp, err := cq.model.GenerateContent(cq.ctx, parts...)
	if err != nil {
		log.Printf("Error generating content: %v\n", err)
		return FileResponse{}, err
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		partString := fmt.Sprintf("%v", part)
		var fileResponse FileResponse
		err := json.Unmarshal([]byte(partString), &fileResponse)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v\n", err)
			return FileResponse{}, err
		}

		return fileResponse, nil
	}

	return FileResponse{}, fmt.Errorf("no response from model")
}

func (cq *ClientQueue) Close() {
	cq.client.Close()
}
