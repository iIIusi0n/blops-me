package file

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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

func createFolder(db *sql.DB, folderName string, storageID int) (int, error) {
	folderName = strings.Trim(folderName, "/")
	folders := strings.Split(folderName, "/")

	var parentID interface{}
	var err error
	for i, folder := range folders {
		if i == 0 {
			parentID = nil
		}

		var files []data.File
		if parentID == nil {
			files, err = data.GetFilesInStorage(db, storageID)
		} else {
			files, err = data.GetFilesInFolder(db, parentID.(int))
		}
		if err != nil {
			return 0, err
		}

		found := false
		for _, file := range files {
			if file.Name == folder && file.Type == "DIR" {
				parentID = file.ID
				found = true
			}
		}

		if found {
			continue
		}

		var file data.File
		if parentID == nil {
			file = data.File{
				Name:         folder,
				Type:         "DIR",
				StorageID:    storageID,
				LastModified: time.Now().Format("2006-01-02"),
			}
		} else {
			file = data.File{
				Name:         folder,
				Type:         "DIR",
				StorageID:    storageID,
				ParentID:     parentID.(int),
				LastModified: time.Now().Format("2006-01-02"),
			}
		}

		var id int
		if parentID == nil {
			id, err = data.AddFile(db, file, false)
		} else {
			id, err = data.AddFile(db, file, true)
		}
		if err != nil {
			return 0, err
		}

		parentID = id
	}

	return parentID.(int), nil
}

func findFile(files []gemini.FileRequest, name string) (gemini.FileRequest, bool) {
	for _, file := range files {
		if file.Name == name {
			return file, true
		}
	}
	return gemini.FileRequest{}, false
}

func saveNewFiles(db *sql.DB, geminiClient *gemini.ClientQueue, files []gemini.FileRequest, storage data.Storage) {
	paths, err := GetFoldersFullPath(db, storage.ID)
	if err != nil {
		log.Printf("Error getting folders full path: %v\n", err)
		return
	}

	resp, err := geminiClient.MakeRequest(files, storage.Name, paths)
	if err != nil {
		log.Printf("Error making request: %v\n", err)
		return
	}

	for _, file := range resp.Files {
		idx := len(strings.Split(file.FullPath, "/")) - 1
		filename := strings.Split(file.FullPath, "/")[idx]
		folder := strings.Join(strings.Split(file.FullPath, "/")[:idx], "/")
		folder = "/" + folder

		parentID, err := createFolder(db, folder, storage.ID)
		if err != nil {
			log.Printf("Error creating folder: %v\n", err)
			return
		}

		originalFile, found := findFile(files, file.OriginalFilename)
		if !found {
			continue
		}

		var fileType string
		if strings.Contains(filename, ".") {
			fileType = strings.ToTitle(strings.Split(filename, ".")[1])
		} else {
			fileType = "FILE"
		}

		newFile := data.File{
			Name:         filename,
			Type:         fileType,
			LastModified: time.Now().Format("2006-01-02"),
			Size:         originalFile.Size,
			Path:         originalFile.Path,
			StorageID:    storage.ID,
			ParentID:     parentID,
		}

		_, err = data.AddFile(db, newFile, true)
		if err != nil {
			log.Printf("Error adding file: %v\n", err)
			return
		}
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
