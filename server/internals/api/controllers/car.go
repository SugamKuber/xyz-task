package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"server/internals/dal"
	"server/internals/utils"

	"github.com/gin-gonic/gin"
)

func UploadCarImage(c *gin.Context) {
	log.Println("Starting file upload...")

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		log.Printf("Error getting form file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to upload image: %v", err)})
		return
	}
	defer file.Close()

	log.Printf("File received: %s, size: %d", header.Filename, header.Size)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image: %v", err)})
		return
	}

	log.Printf("File read successfully, size: %d bytes", len(fileBytes))

	apiKey := os.Getenv("GEMINI_API")
	if apiKey == "" {
		log.Fatal("GEMINI_API not set in environment")
	}
	geminiClient, err := utils.NewClient(apiKey)
	if err != nil {
		log.Printf("Error creating Gemini client: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to initialize Gemini client: %v", err)})
		return
	}

	carType, color, make, model, caption, err := geminiClient.AnalyzeCarImage(fileBytes)
	if err != nil {
		log.Printf("Error processing image with Gemini: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to process image: %v", err)})
		return
	}

	log.Printf("Image processed successfully. Type: %s, Color: %s, Make: %s, Model: %s", carType, color, make, model)

	err = dal.SaveCarData(fileBytes, carType, color, make, model, caption)
	if err != nil {
		log.Printf("Error saving car data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save data: %v", err)})
		return
	}

	log.Println("Car data saved successfully")

	c.JSON(http.StatusOK, gin.H{
		"type":    carType,
		"color":   color,
		"make":    make,
		"model":   model,
		"caption": caption,
	})

	log.Println("Response sent successfully")
}
