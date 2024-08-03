package storage

import (
	"database/sql"
	"log"
	"sort"
	"strconv"
	"strings"

	"blops-me/data"
	"github.com/gin-gonic/gin"
)

func ListStorageHandler(c *gin.Context) {
	userID := c.GetString("user")
	db := c.MustGet("db").(*sql.DB)
	storages, err := data.GetStorages(db, userID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if storages == nil {
		storages = []data.Storage{}
	}

	sort.Slice(storages, func(i, j int) bool {
		return storages[i].Name < storages[j].Name
	})

	for i := range storages {
		storages[i].Name = strings.ToTitle(storages[i].Name)
	}

	c.JSON(200, gin.H{"storages": storages})
}

func CreateStorageHandler(c *gin.Context) {
	userID := c.GetString("user")
	var resp struct {
		StorageName string `json:"storage_name"`
	}
	err := c.BindJSON(&resp)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	if len(resp.StorageName) < 1 && len(resp.StorageName) > 16 {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	resp.StorageName = strings.ToTitle(resp.StorageName)
	db := c.MustGet("db").(*sql.DB)
	err = data.AddNewStorage(db, resp.StorageName, userID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{"message": "Storage created"})
}

func DeleteStorageHandler(c *gin.Context) {
	userID := c.GetString("user")
	storageIdRaw := c.GetHeader("storage-id")
	storageID, err := strconv.Atoi(storageIdRaw)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	err = data.DeleteStorage(db, storageID, userID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Storage deleted"})
}
