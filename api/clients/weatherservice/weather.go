package weather

import (
	"encoding/json"
	"net/http"
	"time"
)

//ClientInterface ...
type ClientInterface interface {
	GetCurrentWeather() (WeatherResponse, error)
}

type WeatherServiceClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewWeatherServiceClient(baseURL, apiKey string) (*WeatherServiceClient, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
		Timeout: 3 * time.Second,
	}

	return &WeatherServiceClient{baseURL: baseURL, apiKey: apiKey, client: client}, nil
}

type MainResponse struct {
	Temp float32 `json:"temp"`
}
type WeatherResponse struct {
	Name string       `json:"name"`
	Main MainResponse `json:"main"`
}

// This lang and long are of New York city
const (
	LAT  = "40.71"
	LONG = "-74.00"
)

// GetCurrentWeather ...
func (c *WeatherServiceClient) GetCurrentWeather() (WeatherResponse, error) {
	url := c.baseURL + "/data/2.5/weather?lat=" + LAT + "&lon=" + LONG + "&appid=" + c.apiKey + "&units=metric"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)

	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()
	var r WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if resp.StatusCode != 200 {
		return WeatherResponse{}, err
	}
	return r, nil
}
