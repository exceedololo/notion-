package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/exceedololo/notion-/internal/models"
	"github.com/exceedololo/notion-/internal/usercases/incrementer"
)

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	var req models.KeyValueIncrementRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the incrementer service
	result := incrementer.Increment(req)

	response := models.KeyValueIncrementResponse{
		Value: result,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
