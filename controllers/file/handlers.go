package file

import (
	"blops-me/data"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func ListFilesHandler(c *gin.Context) {
	storageID := c.Param("id")
	parsedStorageID, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid storage ID"})
		return
	}

	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storageOwner, err := data.GetStorageOwner(db, parsedStorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	isStorage := false
	var parsedPathID int
	pathID := c.Query("path")
	if pathID == "" {
		isStorage = true
	} else {
		parsedPathID, err = strconv.Atoi(pathID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid path ID"})
			return
		}
	}

	var files []data.File
	if isStorage {
		files, err = data.GetFilesInStorage(db, parsedStorageID)
	} else {
		files, err = data.GetFilesInFolder(db, parsedPathID)
	}
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, files)
}

func UploadFilesHandler(c *gin.Context) {
	storageID := c.Param("id")
	parsedStorageID, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid storage ID"})
		return
	}

	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storageOwner, err := data.GetStorageOwner(db, parsedStorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}

	files := form.File["files"]
	for _, file := range files {
		filePath := newFileName()
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		go saveNewFile(db, file.Filename, parsedStorageID, file.Size, filePath)
	}

	c.JSON(200, gin.H{"message": "Files uploaded"})
}

func DeleteFileHandler(c *gin.Context) {
	storageID := c.Param("id")
	parsedStorageID, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid storage ID"})
		return
	}

	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storageOwner, err := data.GetStorageOwner(db, parsedStorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	fileID := c.Query("file")
	if fileID == "" {
		c.JSON(400, gin.H{"error": "Invalid file ID"})
		return
	}

	parsedFileID, err := strconv.Atoi(fileID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid file ID"})
		return
	}

	err = deleteFile(db, parsedFileID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "File deleted"})
}

func GetFileHandler(c *gin.Context) {
	fileID := c.Param("id")
	parsedFileID, err := strconv.Atoi(fileID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid file ID"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	file, err := data.GetFile(db, parsedFileID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	storageOwner, err := data.GetStorageOwner(db, file.StorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	userID := c.GetString("user")
	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	c.FileAttachment(file.Path, file.Name)
}

func GetPathHandler(c *gin.Context) {
	storageID := c.Param("id")
	parsedStorageID, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid storage ID"})
		return
	}

	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storageOwner, err := data.GetStorageOwner(db, parsedStorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	pathID := c.Param("pathID")
	parsedPathID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid path ID"})
		return
	}

	path, err := GetFullPath(db, parsedPathID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, path)
}

func GetParentHandler(c *gin.Context) {
	storageID := c.Param("id")
	parsedStorageID, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid storage ID"})
		return
	}

	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storageOwner, err := data.GetStorageOwner(db, parsedStorageID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	if storageOwner != userID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	pathID := c.Param("pathID")
	parsedPathID, err := strconv.Atoi(pathID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid path ID"})
		return
	}

	id, err := data.GetParentID(db, parsedPathID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, id)
}
