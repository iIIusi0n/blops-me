package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type FileContent struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func extractContent(path, name string) (string, error) {
	scriptPath := "scripts/extract.py"

	cmd := exec.Command("/opt/venv/bin/python", scriptPath, path)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing script: %v, output: %s", err, out.String())
	}

	fileContent := FileContent{
		Filename: name,
		Content:  out.String(),
	}

	jsonOutput, err := json.Marshal(fileContent)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %v", err)
	}

	return string(jsonOutput), nil
}
