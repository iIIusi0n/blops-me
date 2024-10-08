package data

import "database/sql"

type File struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	LastModified string `json:"last_modified"`
	Size         int64  `json:"size"`
	Path         string `json:"path"`
	StorageID    int    `json:"storage_id"`
	ParentID     int    `json:"parent_id"`
}

func AddFile(db *sql.DB, file File, isThereParent bool) (int, error) {
	var err error
	if isThereParent {
		_, err = db.Exec("INSERT INTO file (name, type, last_modified, size, path, storage_id, parent_id) VALUES (?, ?, ?, ?, ?, ?, ?)", file.Name, file.Type, file.LastModified, file.Size, file.Path, file.StorageID, file.ParentID)
	} else {
		_, err = db.Exec("INSERT INTO file (name, type, last_modified, size, path, storage_id) VALUES (?, ?, ?, ?, ?, ?)", file.Name, file.Type, file.LastModified, file.Size, file.Path, file.StorageID)
	}
	if err != nil {
		return 0, err
	}

	var fileID int
	err = db.QueryRow("SELECT id FROM file WHERE name = ? AND path = ? AND storage_id = ?", file.Name, file.Path, file.StorageID).Scan(&fileID)
	if err != nil {
		return 0, err
	}

	return fileID, nil
}

func GetFile(db *sql.DB, fileID int) (File, error) {
	var file File
	err := db.QueryRow("SELECT id, name, type, last_modified, size, path, storage_id FROM file WHERE id = ?", fileID).Scan(&file.ID, &file.Name, &file.Type, &file.LastModified, &file.Size, &file.Path, &file.StorageID)
	if err != nil {
		return File{}, err
	}

	parentID, err := GetParentID(db, fileID)
	if err != nil {
		return File{}, err
	}
	file.ParentID = parentID

	return file, nil
}

func GetParentID(db *sql.DB, fileID int) (int, error) {
	var parentID interface{}
	err := db.QueryRow("SELECT parent_id FROM file WHERE id = ?", fileID).Scan(&parentID)
	if err != nil {
		return 0, err
	}

	if parentID == nil {
		return 0, nil
	} else {
		return int(parentID.(int64)), nil
	}
}

func GetFilesByType(db *sql.DB, storageID int, fileType string) ([]File, error) {
	rows, err := db.Query("SELECT id, name, type, last_modified, size, path, storage_id FROM file WHERE storage_id = ? AND type = ?", storageID, fileType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.Name, &file.Type, &file.LastModified, &file.Size, &file.Path, &file.StorageID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func GetFilesInStorage(db *sql.DB, storageID int) ([]File, error) {
	rows, err := db.Query("SELECT id, name, type, last_modified, size, path, storage_id FROM file WHERE storage_id = ? AND parent_id IS NULL", storageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.Name, &file.Type, &file.LastModified, &file.Size, &file.Path, &file.StorageID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func GetFilesInFolder(db *sql.DB, folderID int) ([]File, error) {
	rows, err := db.Query("SELECT id, name, type, last_modified, size, path, storage_id FROM file WHERE parent_id = ?", folderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(&file.ID, &file.Name, &file.Type, &file.LastModified, &file.Size, &file.Path, &file.StorageID)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func DeleteFile(db *sql.DB, fileID int) error {
	_, err := db.Exec("DELETE FROM file WHERE id = ?", fileID)
	if err != nil {
		return err
	}

	return nil
}
