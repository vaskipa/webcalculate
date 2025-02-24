package handlers

import (
	"encoding/json"
	"net/http"
	"webcalculate/internal/app/repositories"
)

func ExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := repositories.GetRecords()
	if err != nil {
		http.Error(w, "error: internal server error", http.StatusInternalServerError)
	}

	jsonResponse, err := json.Marshal(map[string]interface{}{"expressions": data})

	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
