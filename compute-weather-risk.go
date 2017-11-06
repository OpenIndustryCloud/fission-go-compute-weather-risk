package main

/*
This API will collect Weather Data by consuming
Wunderground API and return summary for given date and city

--- INPUT ---

Historical Weather Data received from Wunderground API
for any given date and city

--- OUTPUT ---
{
	"RiskScore" : 70
	"Description" : "Stormy weather identified"
}

*/
import (
	"encoding/json"
	"net/http"
	"strconv"
)

type WeatherRiskData struct {
	RiskScore   int64  `json:"riskScore"`
	Description string `json:"description"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Please send a valid JSON", 400)
		return
	}
	var historicalData HistoricalData
	err := json.NewDecoder(r.Body).Decode(&historicalData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//check if valid data returned
	if len(historicalData.History.DailySummary) == 0 {
		http.Error(w, "No results found", 400)
		return
	}
	var weatherRiskData = WeatherRiskData{}
	maxWindSpeed, err := strconv.Atoi(historicalData.History.DailySummary[0].Maxwspdm)

	if maxWindSpeed > 20 && maxWindSpeed <= 40 {
		weatherRiskData.RiskScore = 60
		weatherRiskData.Description = "Possibly stormy weather"
	} else if maxWindSpeed > 40 {
		weatherRiskData.RiskScore = 80
		weatherRiskData.Description = "Stormy Weather identified"
	} else {
		weatherRiskData.RiskScore = 20
		weatherRiskData.Description = "Very less likelyhood of Storm"
	}

	//marshal to JSON
	weatherRiskDataJSON, err := json.Marshal(weatherRiskData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(weatherRiskDataJSON)))

}

//Model for WeatherAPI
type HistoricalData struct {
	Response Response `json:"response"`
	History  History  `json:"history"`
}

type Response struct {
	Version string `json:"version"`
}

type History struct {
	DailySummary []DailySummary `json:"dailysummary"`
}

type DailySummary struct {
	Fog          string `json:"fog"`
	Rain         string `json:"rain"`
	Maxtempm     string `json:"maxtempm"`
	Mintempm     string `json:"mintempm"`
	Tornado      string `json:"tornado"`
	Maxpressurem string `json:"maxpressurem"`
	Minpressurem string `json:"minpressurem"`
	Maxwspdm     string `json:"maxwspdm"`
	Minwspdm     string `json:"minwspdm"`
}

// func main() {
// 	println("staritng app..")
// 	http.HandleFunc("/", Handler)
// 	http.ListenAndServe(":8088", nil)
// }
