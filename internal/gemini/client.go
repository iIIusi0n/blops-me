package gemini

import (
	"context"
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

	safety := []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(0.5)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(16 * 1024)
	model.ResponseMIMEType = "application/json"
	model.SafetySettings = safety

	return &ClientQueue{
		ctx:        ctx,
		rpmLimiter: rpmLimiter,
		rpdLimiter: rpdLimiter,
		client:     client,
		model:      model,
	}
}

func (cq *ClientQueue) MakeRequest(fileRequests []FileRequest, storageName string, existingFolders []string) {
	fileUris := make([]string, 0)
	for _, fileRequest := range fileRequests {
		uri, err := uploadToGemini(cq.ctx, cq.client, fileRequest.Path, fileRequest.Name)
		if err != nil {
			log.Printf("Error uploading file: %v\n", err)
			continue
		}

		fileUris = append(fileUris, uri)
	}

	fileParts := make([]genai.Part, 0)
	for _, uri := range fileUris {
		fileParts = append(fileParts, genai.FileData{URI: uri})
	}

	session := cq.model.StartChat()
	session.History = []*genai.Content{
		{
			Role:  "user",
			Parts: fileParts,
		},
	}

	ctx := context.Background()
	if err := cq.rpmLimiter.Wait(ctx); err != nil {
		log.Printf("RPM limit exceeded: %v\n", err)
		return
	}
	if err := cq.rpdLimiter.Wait(ctx); err != nil {
		log.Printf("RPD limit exceeded: %v\n", err)
		return
	}

	generatedPrompt := generatePrompt(storageName, existingFolders)
	resp, err := session.SendMessage(cq.ctx, genai.Text(generatedPrompt))
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
		return
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}

func (cq *ClientQueue) Close() {
	cq.client.Close()
}
