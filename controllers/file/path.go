package file

import (
	"blops-me/data"
	"database/sql"
	"fmt"
)

func GetFullPath(db *sql.DB, fileID int) (string, error) {
	fullPath := ""
	for {
		file, err := data.GetFile(db, fileID)
		if err != nil {
			return "", err
		}
		fullPath = fmt.Sprintf("/%s%s", file.Name, fullPath)
		if file.ParentID == 0 {
			break
		}
		fileID = file.ParentID
	}
	return fullPath, nil
}

func GetFoldersFullPath(db *sql.DB, storageID int) ([]string, error) {
	var foldersFullPath []string
	folders, err := data.GetFilesByType(db, storageID, "DIR")
	if err != nil {
		return nil, err
	}
	for _, folder := range folders {
		fullPath, err := GetFullPath(db, folder.ID)
		if err != nil {
			return nil, err
		}
		foldersFullPath = append(foldersFullPath, fullPath)
	}
	return foldersFullPath, nil
}
