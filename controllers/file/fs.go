package file

import (
	"blops-me/data"
	"database/sql"
	"log"
	"math/rand"
	"os"
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

func saveNewFile(db *sql.DB, originalName string, storageID int, size int64, savedPath string) {
	log.Printf("Saving file %s to database. Storage ID: %d, Size: %d, Path: %s", originalName, storageID, size, savedPath)
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
