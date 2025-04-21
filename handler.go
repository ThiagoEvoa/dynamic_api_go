package main

import (
	"encoding/json"
	"net/http"
)

var obj map[string]interface{}

func handleHttpFunction(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		GetObj(responseWriter, request)
	case http.MethodPost:
		PostObj(responseWriter, request)
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetObj(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	response := obj
	json.NewEncoder(responseWriter).Encode(response)
}

func PostObj(responseWriter http.ResponseWriter, request *http.Request) {
	var payloadData map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&payloadData)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	defer request.Body.Close()
	obj = payloadData
	responseWriter.WriteHeader(http.StatusCreated)
}
