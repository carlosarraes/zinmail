package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func JsonWriter(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := Response{Message: message}
	jsonErr, _ := json.Marshal(err)
	i, _ := w.Write(jsonErr)
	fmt.Println(i)
}
