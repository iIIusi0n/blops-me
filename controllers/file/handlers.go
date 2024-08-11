package file

import (
	"database/sql"
	"log"
	"strconv"

	"blops-me/data"
	"blops-me/internal/gemini"
	"github.com/gin-gonic/gin"
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
	_, err = data.GetStorage(db, parsedStorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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
	storage, err := data.GetStorage(db, parsedStorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}

	geminiClient := c.MustGet("geminiClient").(*gemini.ClientQueue)

	files := form.File["files"]
	fileRequests := make([]gemini.FileRequest, 0)
	for _, file := range files {
		filePath := newFileName()
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			log.Printf("Error saving file: %v\n", err)
			continue
		}

		fileRequests = append(fileRequests, gemini.NewFileRequest(file.Filename, file.Size, filePath, parsedStorageID, isStorage, parsedPathID))
	}
	go saveNewFiles(db, geminiClient, fileRequests, storage)

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
	_, err = data.GetStorage(db, parsedStorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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

	userID := c.GetString("user")
	_, err = data.GetStorage(db, file.StorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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
	_, err = data.GetStorage(db, parsedStorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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
	_, err = data.GetStorage(db, parsedStorageID, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
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
