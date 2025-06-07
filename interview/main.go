package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/random-json", handleRandomJson)
	http.ListenAndServe(":8080", nil)
}
func handleRandomJson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//handling error
		http.Error(w, "Only post allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		//handle error
	}
	//get query parameters
	queryParams := r.URL.Query().Get("queryParams")
	var data any
	marshalErr := json.Unmarshal(body, &data)
	if marshalErr != nil {
		http.Error(w, "Only post allowed", http.StatusMethodNotAllowed)
		return
	}

	collectedData := collectData(data)
	// value := getValuefromMap(data, queryParams)
	value := getValueByPath(data, queryParams)
	fmt.Println("data is ", value)
	fmt.Println("collected data is:", collectedData)
	// Respond with the collected data
	w.Header().Set("Content-Type", "application/json")
	x := map[string]interface{}{
		"value":          value,
		"collected_data": collectedData,
	}
	response, err := json.Marshal(x)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

}

func getValueByPath(data interface{}, path string) interface{} {
	re := regexp.MustCompile(`[^.\[\]]+`) // matches keys and array indexes
	keys := re.FindAllString(path, -1)

	var current interface{} = data
	for _, key := range keys {
		switch curr := current.(type) {
		case map[string]interface{}:
			current = curr[key]
		case []interface{}:
			index, err := strconv.Atoi(key)
			if err != nil || index < 0 || index >= len(curr) {
				return nil
			}
			current = curr[index]
		default:
			return nil
		}
	}
	return current
}

func getValuefromMap(collectedData any, queryParams string) interface{} {

	keys := strings.Split(queryParams, ".")

	currentValue := any(collectedData)

	for i, key := range keys {
		if data, ok := currentValue.(map[string]interface{}); ok {
			if val, ok := data[key]; ok {
				currentValue = val
			} else {
				return nil
			}
		} else {
			return nil
		}
		if i == len(keys)-1 {
			return currentValue
		}
	}
	return nil
}

func collectData(data any) map[string]interface{} {
	// Initialize a map to collect data
	collectedData := make(map[string]interface{})
	// Process the collected data
	for key, value := range data.(map[string]interface{}) {

		if nestedMap, ok := value.(map[string]interface{}); ok {
			for nestedKey, nestedValue := range nestedMap {
				// Process nested key-value pairs
				collectedData[nestedKey] = nestedValue
			}
		} else {
			// Process top-level key-value pairs
			collectedData[key] = value
		}
	}
	return collectedData
}
