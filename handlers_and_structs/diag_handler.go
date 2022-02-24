package handlers_and_structs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var Elapsed_Time time.Time

func Diag_Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		initial_url := "http://universities.hipolabs.com/"
		secondary_url := "https://restcountries.com/"

		// Create new requests
		req, err := http.NewRequest(http.MethodGet, initial_url, nil)
		if err != nil {
			fmt.Errorf("Error in creating request:", err.Error())
		}

		// Instantiate the client
		client := &http.Client{}

		req2, err := http.NewRequest(http.MethodGet, secondary_url, nil)
		if err != nil {
			fmt.Errorf("Error in creating request:", err.Error())
		}

		// Issue request
		res, err := client.Do(req)
		if err != nil {
			fmt.Errorf("Error in response:", err.Error())
		}

		res2, err := client.Do(req2)
		if err != nil {
			fmt.Errorf("Error in response:", err.Error())
		}

		// Instantiate json_encoder
		json_encoder := json.NewEncoder(w)

		api_diag := Direct_diag{
			Version:         "v1",
			UniversitiesAPI: res.Status,
			CountriesSPI:    res2.Status,
			Uptime:          time.Duration(time.Since(Elapsed_Time).Seconds())}

		w.Header().Add("content-type", "application/json")

		// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
		err = json_encoder.Encode(api_diag)
		if err != nil {
			http.Error(w, "WOAH! There was an Error during encoding", http.StatusInternalServerError)
			return
		}
	}
}
