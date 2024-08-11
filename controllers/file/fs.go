package file

import (
	"database/sql"
	"log"
	"math/rand"
	"os"

	"blops-me/data"
	"blops-me/internal/gemini"
)

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func newFileName() string {
	path := "uploads/"
	fileName := randomString(64)
	return path + fileName
}

func saveNewFiles(db *sql.DB, geminiClient *gemini.ClientQueue, files []gemini.FileRequest, storage data.Storage) {
	if paths, err := GetFoldersFullPath(db, storage.ID); err == nil {
		geminiClient.MakeRequest(files, storage.Name, paths)
	} else {
		log.Printf("Error getting folders full path: %v\n", err)
	}
}

func deleteFile(db *sql.DB, fileID int) error {
	file, err := data.GetFile(db, fileID)
	if err != nil {
		return err
	}

	err = data.DeleteFile(db, fileID)
	if err != nil {
		return err
	}

	if file.Type == "DIR" {
		files, err := data.GetFilesInFolder(db, fileID)
		if err != nil {
			return err
		}
		for _, file := range files {
			err := deleteFile(db, file.ID)
			if err != nil {
				return err
			}
		}
	} else {
		err = os.Remove(file.Path)
		if err != nil {
			return err
		}
	}

	return nil
}
