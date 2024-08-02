package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create the uploads directory if it doesn't exist
		uploadsDir := "_files"
		if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
			err = os.Mkdir(uploadsDir, 0755)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}

		// Save the file to the uploads directory
		filename := file.Filename
		savePath := filepath.Join(uploadsDir, filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		absPath, err := filepath.Abs(savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
			"path":    absPath,
		})
	}
}
