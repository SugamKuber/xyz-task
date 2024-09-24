package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	APIKey      string
	genaiClient *genai.Client
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is not set")
	}

	ctx := context.Background()
	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %v", err)
	}

	return &Client{
		APIKey:      apiKey,
		genaiClient: genaiClient,
	}, nil
}

func (c *Client) AnalyzeCarImage(imageBytes []byte) (string, string, string, string, string, error) {
	if c.APIKey == "" {
		return "", "", "", "", "", fmt.Errorf("GEMINI_API_KEY is not set")
	}

	ctx := context.Background()

	prompt := `Analyze this car image and provide the following details in JSON format: 
    {
        "type": "car type (e.g., SUV, race car, mini)", 
        "color": "color of the car", 
        "make": "company of the car", 
        "model": "model of the car", 
        "caption": "a brief caption describing the car"
    }`

	model := c.genaiClient.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(0.2)
	model.SetTopK(32)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(256)

	model.ResponseMIMEType = "application/json"

	model.ResponseSchema = &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"type":    {Type: genai.TypeString},
			"color":   {Type: genai.TypeString},
			"make":    {Type: genai.TypeString},
			"model":   {Type: genai.TypeString},
			"caption": {Type: genai.TypeString},
		},
		Required: []string{"type", "color", "make", "model", "caption"},
	}

	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt),
		genai.ImageData("image/jpeg", imageBytes),
	)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("failed to generate content: %v", err)
	}

	fmt.Printf("Full Gemini API Response: %+v\n", resp)

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", "", "", "", "", fmt.Errorf("no valid response from Gemini")
	}

	var jsonResponse string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			jsonResponse = string(textPart)
			break
		}
	}

	if jsonResponse == "" {
		return "", "", "", "", "", fmt.Errorf("no text content in the response")
	}

	var result struct {
		Type    string `json:"type"`
		Color   string `json:"color"`
		Make    string `json:"make"`
		Model   string `json:"model"`
		Caption string `json:"caption"`
	}

	err = json.Unmarshal([]byte(jsonResponse), &result)
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return result.Type, result.Color, result.Make, result.Model, result.Caption, nil
}
