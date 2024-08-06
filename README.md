# WeatherGoRound

**WeatherGoRound** is a Go application designed to showcase the `gofr` framework at MITS DevSummit'24. This application fetches weather data and generates a sassy weather recommendation using the Open-Meteo and Groq APIs.

## Features

- Fetches current weather data using the Open-Meteo API.
- Sends the weather data to the Groq API for a humorous and sassy recommendation.
- Provides endpoints to access raw weather data and sassy recommendations.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- Groq API key (available from the [Groq website](https://groq.com/))
- Internet connection for API calls

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/vaishakhsnair/weathergoround.git
   cd weathergoround
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Set up environment variables:**

   Create a `.env` file in the root directory with your Groq API key:

   ```
   GROQ_API_KEY=your_groq_api_key
   ```

4. **Create the `sassy.txt` file:**

   In the `prompts` directory, create a `sassy.txt` file with the following content:

   ```
   You are a sassy weather forecaster. Given a weather report, provide a sassy recommendation for the upcoming weather (i.e., weather after the current time).

   Respond only with the following JSON format:

   {
     "recommendation": "<your recommendation>",
     "current_time": "<current time>"
   }
   ```

## Running the Application

1. **Start the application:**

   ```sh
   go run main.go
   ```

2. **Access the endpoints:**

   - **Sassy Weather Recommendation:** Navigate to `http://localhost:8080/sassy` to get a sassy weather recommendation based on the current weather.
   - **Raw Weather Data:** Navigate to `http://localhost:8080/` to get the current weather data.

## Code Overview

- `main.go`: Contains the core logic and setup for the application.
  - `getWeather()`: Retrieves current weather data from the Open-Meteo API.
  - `getGPTResponse()`: Sends weather data to the Groq API and receives a sassy recommendation.
  - `main()`: Initializes and starts the HTTP server using the `gofr` framework.

## Purpose

This project was specifically developed to demonstrate the capabilities of the `gofr` framework at MITS DevSummit'24. It highlights how `gofr` can be used to build a simple yet powerful application integrating external APIs.

## Troubleshooting

- Ensure that the `.env` file contains a valid Groq API key.
- Verify the existence and format of the `prompts/sassy.txt` file.
- Check network connectivity if API calls fail.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests for improvements or bug fixes.

