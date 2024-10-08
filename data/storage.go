package data

import (
	"database/sql"

	"blops-me/utils"
)

type Storage struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func AddNewStorage(db *sql.DB, storageName string, userID string) error {
	storageName = utils.EncodeBase(storageName)

	exists := 0
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM storage WHERE storage_name = ? AND user_id = ?)", storageName, userID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists == 1 {
		return nil
	}

	_, err = db.Exec("INSERT INTO storage (storage_name, user_id) VALUES (?, ?)", storageName, userID)
	return err
}

func GetStorages(db *sql.DB, userID string) ([]Storage, error) {
	rows, err := db.Query("SELECT id, storage_name FROM storage WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storages []Storage
	for rows.Next() {
		var storage Storage
		err := rows.Scan(&storage.ID, &storage.Name)
		if err != nil {
			return nil, err
		}
		storage.Name = utils.DecodeBase(storage.Name)
		storages = append(storages, storage)
	}

	return storages, nil
}

func GetStorage(db *sql.DB, storageID int, userID string) (Storage, error) {
	var storage Storage
	err := db.QueryRow("SELECT id, storage_name FROM storage WHERE id = ? AND user_id = ?", storageID, userID).Scan(&storage.ID, &storage.Name)
	if err != nil {
		return storage, err
	}
	storage.Name = utils.DecodeBase(storage.Name)
	return storage, nil
}

func DeleteStorage(db *sql.DB, storageID int, userID string) error {
	_, err := db.Exec("DELETE FROM storage WHERE id = ? AND user_id = ?", storageID, userID)
	return err
}
