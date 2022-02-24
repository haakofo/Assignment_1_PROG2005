package handlers_and_structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Neighbour_unis_Handler(w http.ResponseWriter, r *http.Request) {
	initial_url := "http://universities.hipolabs.com/"
	secondary_url := "https://restcountries.com/"
	string_parts := strings.Split(r.URL.Path, "/")
	uni_name := string_parts[5]
	country_name := string_parts[4]
	var isoCode string

	// Create header with content type
	w.Header().Add("content-type", "application/json")

	// Instantiate the api_client
	api_client := &http.Client{}

	// Create new request
	req, err := http.NewRequest(http.MethodGet, initial_url+"search?name="+uni_name+"&country="+country_name, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Issue request
	res, err := api_client.Do(req)
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	// Instantiate json_decoder
	json_decoder := json.NewDecoder(res.Body)

	_, err = json_decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	for json_decoder.More() {
		var uniinfo Complete_Unifinfo
		err = json_decoder.Decode(&uniinfo)
		if err != nil {
			log.Fatal(err)
		}

		isoCode = uniinfo.Isocode

		// Create new request
		req, err := http.NewRequest(http.MethodGet, secondary_url+"v3.1/alpha/"+isoCode+"?fields=maps,languages", nil)
		if err != nil {
			fmt.Errorf("Here is the error in creating the request:", err.Error())
		}

		// Issue request
		res, err = api_client.Do(req)
		if err != nil {
			fmt.Errorf("Here is the error in the response:", err.Error())
		}

		countryData, _ := ioutil.ReadAll(res.Body)

		err = json.Unmarshal(countryData, &uniinfo)
		if err != nil {
			log.Println(err.Error())
		}

		// Instantiate encoder
		encoder := json.NewEncoder(w)

		// Encode specific content
		err = encoder.Encode(uniinfo)
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
		}

	}

	_, err = json_decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	// Create new request
	req, err = http.NewRequest(http.MethodGet, secondary_url+"v3.1/alpha/"+isoCode+"?fields=borders", nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Issue request
	res, err = api_client.Do(req)
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	var borders Borders

	country_borders, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(country_borders, &borders)
	if err != nil {
		log.Println(err.Error())
	}

	for i := range borders.Isocodes {
		// Get full country name from RESTcountries API
		// Create new request
		req, err = http.NewRequest(http.MethodGet, secondary_url+"v3.1/alpha/"+borders.Isocodes[i]+"?fields=name", nil)
		if err != nil {
			fmt.Errorf("Error in creating request:", err.Error())
		}

		// Issue request
		res, err = api_client.Do(req)
		if err != nil {
			fmt.Errorf("Error in response:", err.Error())
		}

		var country Country
		county_byte_data, _ := ioutil.ReadAll(res.Body)

		err = json.Unmarshal(county_byte_data, &country)
		if err != nil {
			log.Println(err.Error())
		}

		// Create new request
		req, err = http.NewRequest(http.MethodGet, initial_url+"search?name="+uni_name+"&country="+country.CountryName.Common, nil)
		if err != nil {
			fmt.Errorf("Here is the error in creating the request:", err.Error())
		}

		// Issue request
		res, err = api_client.Do(req)
		if err != nil {
			fmt.Errorf("Here is the error in the response:", err.Error())
		}

		json_decoder = json.NewDecoder(res.Body)

		_, err = json_decoder.Token()
		if err != nil {
			log.Fatal(err)
		}

		for json_decoder.More() {
			var uniinfo Complete_Unifinfo
			err = json_decoder.Decode(&uniinfo)
			if err != nil {
				log.Fatal(err)
			}

			// Create new request
			req, err = http.NewRequest(http.MethodGet, "https://restcountries.com/v3.1/alpha/"+borders.Isocodes[i]+"?fields=maps,languages", nil)
			if err != nil {
				fmt.Errorf("There was an error in creating request:", err.Error())
			}

			// Issue request

			res, err = api_client.Do(req)
			if err != nil {
				fmt.Errorf("There was an error in response:", err.Error())
			}

			current_country_data, _ := ioutil.ReadAll(res.Body)

			err = json.Unmarshal(current_country_data, &uniinfo)
			if err != nil {
				log.Println(err.Error())
			}

			// Instantiate json_encoder
			json_encoder := json.NewEncoder(w)

			// Encode specific content --> Alternative: "err := json.NewEncoder(w).Encode(location)"
			err = json_encoder.Encode(uniinfo)
			if err != nil {
				http.Error(w, "UH OH! There was an error during encoding", http.StatusInternalServerError)
				return
			}

		}

		_, err = json_decoder.Token()
		if err != nil {
			log.Fatal(err)
		}
	}

}
