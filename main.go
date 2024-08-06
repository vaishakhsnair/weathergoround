package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
)

// GroqResponse represents the structure of the sassy weather response.
type GroqResponse struct {
	Recommendation string `json:"recommendation"`
	CurrentTime    string `json:"current_time"`
}

// getWeather fetches the current weather data from the Open-Meteo API.
func getWeather() (string, error) {
	latitude := 9.9581
	longitude := 76.3634
	client := resty.New()

	// Construct the request URL for the Open-Meteo API.
	requestURL := "https://api.open-meteo.com/v1/forecast?latitude=" + fmt.Sprintf("%f", latitude) + "&longitude=" + fmt.Sprintf("%f", longitude) + "&current_weather=true"
	response, err := client.R().Get(requestURL)

	if err != nil {
		return "", err
	}

	return string(response.Body()), nil
}

// getSassyResponse sends the weather data to the Groq API and returns a sassy weather recommendation.
func getGPTResponse(weather string, prompt string) (GroqResponse, error) {
	client := resty.New()
	apiKey := os.Getenv("GROQ_API_KEY")

	//fmt.Println(string(prompt))

	current_time := time.Now().Local()

	// Construct the request payload for the Groq API.
	requestBody := map[string]interface{}{
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": string(prompt),
			},
			{
				"role":    "user",
				"content": weather + "\n Current time: " + current_time.String(),
			},
		},
		"model":       "llama3-8b-8192",
		"temperature": 1.0,
		"max_tokens":  1024,
		"top_p":       1.0,
		"stream":      false,
		"stop":        nil,
	}

	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+apiKey).
		SetBody(requestBody).
		Post("https://api.groq.com/openai/v1/chat/completions")

	if err != nil {
		return GroqResponse{}, err
	}

	// Print the raw response for debugging purposes
	//fmt.Println(string(response.Body()))

	// Parse the response from the Groq API
	var parsedResponse struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	err = json.Unmarshal(response.Body(), &parsedResponse)
	if err != nil {
		return GroqResponse{}, err
	}

	// Extract the recommendation and current time from the response content
	var recommendation GroqResponse
	err = json.Unmarshal([]byte(parsedResponse.Choices[0].Message.Content), &recommendation)
	if err != nil {
		return GroqResponse{}, err
	}

	return recommendation, nil
}

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gofr.New()

	// Define a handler for the root endpoint
	app.GET("/sassy", func(context *gofr.Context) (interface{}, error) {
		weather, err := getWeather()
		if err != nil {
			context.Error(http.StatusInternalServerError, err.Error())
			return nil, err
		}

		//read the prompt from the file

		prompt, err := os.ReadFile("prompts/sassy.txt")
		if err != nil {
			context.Error(http.StatusInternalServerError, err.Error())
			return nil, err
		}

		response, err := getGPTResponse(weather, string(prompt))

		if err != nil {
			context.Error(http.StatusInternalServerError, err.Error())
			return nil, err
		}

		return response, nil
	})

	app.GET("/", func(context *gofr.Context) (interface{}, error) {
		weather, err := getWeather()
		if err != nil {
			context.Error(http.StatusInternalServerError, err.Error())
			return nil, err
		}

		return weather, nil
	})

	// Start the application
	app.Run()
}
