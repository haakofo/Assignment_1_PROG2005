package handlers_and_structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Uni_info_Handler(w http.ResponseWriter, r *http.Request) {
	URI := strings.Replace(r.RequestURI, UNIINFO_PATH, "", -1)
	initial_url := "http://universities.hipolabs.com/search?name=" + URI
	secondary_url := "https://restcountries.com/v3.1/alpha/"

	w.Header().Add("content-type", "application/json")

	// Create new request
	req, err := http.NewRequest(http.MethodGet, initial_url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Instantiate the client
	client := &http.Client{}

	// Issue request
	res, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	decoder := json.NewDecoder(res.Body)

	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	for decoder.More() {
		var uniinfo Complete_Unifinfo
		err := decoder.Decode(&uniinfo)
		if err != nil {
			log.Fatal(err)
		}

		isoCode := uniinfo.Isocode

		// Create new request
		req2, err := http.NewRequest(http.MethodGet, secondary_url+isoCode+"?fields=maps,languages", nil)
		if err != nil {
			fmt.Errorf("Error in creating request:", err.Error())
		}

		// Issue request

		res2, err := client.Do(req2)
		if err != nil {
			fmt.Errorf("Error in response:", err.Error())
		}

		retrived_country_data, _ := ioutil.ReadAll(res2.Body)

		err = json.Unmarshal(retrived_country_data, &uniinfo)
		if err != nil {
			log.Println(err.Error())
		}

		// Instantiate json_encoder
		json_encoder := json.NewEncoder(w)

		// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
		err = json_encoder.Encode(uniinfo)
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
		}

	}

	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

}
