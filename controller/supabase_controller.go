package controller

import (
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadController struct {
	uploadService service.UploadService
}

func NewUploadController(uploadService service.UploadService) *UploadController {
	return &UploadController{uploadService: uploadService}
}

func (ctl *UploadController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	bucket := "uploads"
	url, err := ctl.uploadService.UploadFile(file, bucket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "File uploaded successfully",
		"file_name": file.Filename,
		"file_url":  url,
	})
}
