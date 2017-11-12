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
	"io"
	"net/http"
	"strconv"
)

type WeatherRiskData struct {
	Status      int    `json:"status"`
	RiskScore   int64  `json:"riskScore"`
	Description string `json:"description"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	println("Executing Compute Weather Risk...")

	var historicalData HistoricalData
	err := json.NewDecoder(r.Body).Decode(&historicalData)
	if err == io.EOF || err != nil {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if valid data returned
	if len(historicalData.History.DailySummary) == 0 {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	var weatherRiskData = WeatherRiskData{}
	maxWindSpeed, err := strconv.Atoi(historicalData.History.DailySummary[0].Maxwspdm)

	if maxWindSpeed > 20 && maxWindSpeed <= 40 {
		weatherRiskData.RiskScore = 60
		weatherRiskData.Description = "Possibly stormy weather"
	} else if maxWindSpeed > 50 {
		weatherRiskData.RiskScore = 80
		weatherRiskData.Description = "Stormy Weather identified"
	} else {
		weatherRiskData.RiskScore = 20
		weatherRiskData.Description = "Very less likelyhood of Storm"
	}

	//marshal to JSON
	weatherRiskData.Status = 200
	weatherRiskDataJSON, err := json.Marshal(weatherRiskData)
	if err != nil {
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	println(string(weatherRiskDataJSON))
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(weatherRiskDataJSON)))

}

func createErrorResponse(w http.ResponseWriter, message string, status int) {
	errorJSON, _ := json.Marshal(&Error{
		Status:  status,
		Message: message})
	//Send custom error message to caller
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(errorJSON))
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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
