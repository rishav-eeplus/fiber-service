package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// for sending json response
type FiberRouteData map[string]interface{}

// handler for fetching data from api based on lat and lon
func FiberDataHandler(w http.ResponseWriter, r *http.Request) {
	FiberDataUrl, exists := os.LookupEnv("FiberDataUrl")
	if !exists {
		log.Fatal("API Url not present")
	}
	query := r.URL.Query()
	latString := query.Get("lat")
	lonString := query.Get("lon")
	siteDistanceString := query.Get("siteDistance")
	if latString == "" || lonString == "" || siteDistanceString == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	url := fmt.Sprintf("%s?lat=%s&lon=%s&siteDistance=%s", FiberDataUrl, latString, lonString, siteDistanceString)
	response, err := apiRequest(url)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	json.NewEncoder(w).Encode(response)

}

// function for handling api request
func apiRequest(url string) (FiberRouteData, error) {
	var response FiberRouteData
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}
	FiberDataApiKey, exists := os.LookupEnv("FiberDataApiKey")
	if !exists {
		log.Fatal("API key not present")
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", FiberDataApiKey)
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}
