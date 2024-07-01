package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type IPAPIResponse struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

const (
	ipifyURL       = "https://api.ipify.org?format=json"
	ipAPIURL       = "http://ip-api.com/json/"
	openWeatherURL = "http://api.openweathermap.org/data/2.5/weather?units=metric&appid="
)

var openWeatherAPIKey string
var port string

func init() {
	if err := godotenv.Load(); err != nil {
	    log.Println("No .env file found, using system environment variables")
	}
	
	openWeatherAPIKey = os.Getenv("OPENWEATHER_API_KEY")
	if openWeatherAPIKey == "" {
	    log.Fatal("OPENWEATHER_API_KEY environment variable not set")
	}
 
	port = os.Getenv("PORT")
	if port == "" {
	    port = "8080"
	    log.Println("PORT not set, using default:", port)
	}
}

func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getClientIP() (string, error) {
	var ipResp struct {
		IP string `json:"ip"`
	}
	err := getJSON(ipifyURL, &ipResp)
	return ipResp.IP, err
}

func getLocation(ip string) (IPAPIResponse, error) {
	var location IPAPIResponse
	err := getJSON(ipAPIURL+ip, &location)
	return location, err
}

func getWeather(lat, lon float64) (float64, error) {
	url := fmt.Sprintf("%s%s&lat=%f&lon=%f", openWeatherURL, openWeatherAPIKey, lat, lon)
	var weather WeatherResponse
	err := getJSON(url, &weather)
	return weather.Main.Temp, err
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	visitorName := r.URL.Query().Get("visitor_name")
	if visitorName == "" {
		visitorName = "Guest"
	}

	var (
		clientIP string
		err      error
	)
	clientIP, err = getClientIP()
	if err != nil {
		http.Error(w, "Error getting client IP", http.StatusInternalServerError)
		return
	}

	location, err := getLocation(clientIP)
	if err != nil {
		http.Error(w, "Error getting location", http.StatusInternalServerError)
		return
	}

	temp, err := getWeather(location.Lat, location.Lon)
	if err != nil {
		http.Error(w, "Error getting weather", http.StatusInternalServerError)
		return
	}

	response := Response{
		ClientIP: clientIP,
		Location: fmt.Sprintf("%s, %s", location.City, location.Country),
		Greeting: fmt.Sprintf("Hello, %s! The temperature is %.1f degrees Celsius in %s, %s",
			visitorName, temp, location.City, location.Country),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Printf("running...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
