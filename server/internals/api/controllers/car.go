package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"server/internals/dal"
	"server/internals/utils"
)

func UploadCarImage(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}

	carType, color, make, model := utils.ProcessImage(fileBytes)
	caption := utils.GenerateCaption(carType, color, make, model)

	err = dal.SaveCarData(fileBytes, carType, color, make, model, caption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type":    carType,
		"color":   color,
		"make":    make,
		"model":   model,
		"caption": caption,
	})
}
