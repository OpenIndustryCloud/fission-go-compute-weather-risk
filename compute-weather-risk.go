package main

/*
This API will compute Risk based on the Weather data provided

--- INPUT ---

JSON output data of Wundergroud History API

{
    "response": {
        "version": "0.1"
    },
    "history": {
        "dailysummary": [
            {
					 //
					 //
                "maxwspdm": "50",
                "minwspdm": "13"
            }
        ],
        "observations": [
            {

            }
        ]
    }
}
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

// Handler funtion compute the Rsik for Stormy weather based on Wind Speed.
func Handler(w http.ResponseWriter, r *http.Request) {

	println("Executing Compute Weather Risk...")

	//Process Post Data
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

	//Create Reposne Struct
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

	//write risk data to out stream
	println(string(weatherRiskDataJSON))
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(weatherRiskDataJSON)))

}

// createErrorResponse - this function forms a error reposne with
// error message and http code
func createErrorResponse(w http.ResponseWriter, message string, status int) {
	errorJSON, _ := json.Marshal(&Error{
		Status:  status,
		Message: message})
	//Send custom error message to caller
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(errorJSON))
}

// Error - error object
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Output Data Model
type WeatherRiskData struct {
	Status      int    `json:"status"`
	RiskScore   int64  `json:"riskScore"`
	Description string `json:"description"`
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
