package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosarraes/mimemail/server/models"
	"github.com/carlosarraes/mimemail/server/utils"
)

type JsonBody struct {
	SearchTerm string `json:"searchTerm"`
}

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	var body JsonBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.JsonWriter(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	data, err := models.GetEmails(body.SearchTerm)
	if err != nil {
		utils.JsonWriter(w, http.StatusInternalServerError, "Error getting emails")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
