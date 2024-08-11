package gemini

import (
	"fmt"
	"os"
	"strings"
)

func generatePrompt(storageName string, existingFolders []string) string {
	template, _ := os.ReadFile("assets/prompt.txt")

	pathString := ""
	for _, folder := range existingFolders {
		pathString += fmt.Sprintf(`"%s", `, folder)
	}
	pathString = strings.TrimSuffix(pathString, ", ")

	prompt := string(template)
	prompt = strings.ReplaceAll(prompt, "<STORAGE_NAME>", storageName)
	prompt = strings.ReplaceAll(prompt, "<EXISTING_FOLDERS>", pathString)

	return prompt
}
