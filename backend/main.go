package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Create upload directory if it doesn't exist
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Enable CORS (for frontend)
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go File Upload API Running",
		})
	})

	router.POST("/upload", uploadFiles)

	fmt.Println("Server Started")
	fmt.Println("http://localhost:8080")

	router.Run(":8080")
}

func uploadFiles(c *gin.Context) {

	oldFile, err := c.FormFile("oldFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Old file is required",
		})
		return
	}

	newFile, err := c.FormFile("newFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New file is required",
		})
		return
	}

	timestamp := time.Now().Format("20060102150405")

	oldFileName := timestamp + "_old_" + filepath.Base(oldFile.Filename)
	newFileName := timestamp + "_new_" + filepath.Base(newFile.Filename)

	oldPath := filepath.Join("uploads", oldFileName)
	newPath := filepath.Join("uploads", newFileName)

	if err := c.SaveUploadedFile(oldFile, oldPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save old file",
		})
		return
	}

	if err := c.SaveUploadedFile(newFile, newPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save new file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Files uploaded successfully",
		"oldFile": oldFileName,
		"newFile": newFileName,
		"oldPath": oldPath,
		"newPath": newPath,
	})
}