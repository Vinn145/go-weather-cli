package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city,country>")
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ API key not found. Please set OPENWEATHER_API_KEY")
		return
	}

	city := os.Args[1]
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=id", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("API returned status:", resp.Status)
		return
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Printf("🌍 Cuaca di %s\n", weather.Name)
	fmt.Printf("🌡️ Suhu: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("💧 Kelembapan: %d%%\n", weather.Main.Humidity)
	if len(weather.Weather) > 0 {
		fmt.Printf("☁️ Kondisi: %s\n", weather.Weather[0].Description)
	}
}
