package gemini

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func uploadToGemini(ctx context.Context, client *genai.Client, path, name string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	options := genai.UploadFileOptions{
		DisplayName: name,
	}
	fileData, err := client.UploadFile(ctx, "", file, &options)
	if err != nil {
		return "", err
	}

	log.Printf("Uploaded file %s as: %s\n", fileData.DisplayName, fileData.URI)
	return fileData.URI, nil
}
