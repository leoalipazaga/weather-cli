package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Forecast struct {
	Id          int32  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Weather struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float32 `json:"lat"`
		Long    float32 `json:"long"`
	} `json:"location"`
	Current struct {
		TempC      float32 `json:"temp_c"`
		TempF      float32 `json:"temp_f"`
		WindKph    float32 `json:"wind_kph"`
		WindDegree float32 `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		Humidity   float32 `json:"humidity"`
		Condition  struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int32  `json:"code"`
		} `json:"condition"`
	} `json:"current"`
}

func createColorTemp(temp float32) func(a ...interface{}) string {
	if temp <= -0 {
		return color.New(color.FgHiBlue).SprintFunc()
	}

	if temp <= 13 {
		return color.New(color.FgBlue).SprintFunc()
	}

	if temp <= 21 {
		return color.New(color.FgHiGreen).SprintFunc()
	}

	if temp <= 27 {
		return color.New(color.FgGreen).SprintFunc()
	}

	if temp <= 33 {
		return color.New(color.FgHiRed).SprintFunc()
	}

	return color.New(color.FgRed).SprintFunc()
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	city := "London"
	if len(os.Args) >= 2 {
		city = os.Args[1]
	}
	parseURL, _ := url.Parse(os.Getenv("RAPID_API_URL"))
	query := parseURL.Query()
	query.Set("q", city)
	parseURL.RawQuery = query.Encode()
	req, _ := http.NewRequest("GET", parseURL.String(), nil)
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPID_API_HOST"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(res.Body)
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		panic(err)
	}
	fmt.Printf("%s, %s\n", weather.Location.Name, weather.Location.Country)
	fmt.Printf("%s\n", weather.Current.Condition.Text)
	fmt.Printf("Wind: %.0fC %s\n", weather.Current.WindKph, weather.Current.WindDir)
	fmt.Printf("Centigrades: %s°C\n", createColorTemp(weather.Current.TempC)(weather.Current.TempC))
	fmt.Printf("Farenheit: %.0fF\n", weather.Current.TempF)
}
