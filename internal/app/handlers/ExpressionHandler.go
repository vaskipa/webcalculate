package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webcalculate/internal/app/repositories"
)

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "error: id is not valid", http.StatusUnprocessableEntity)
	}
	// плохая практика импортировать сразу из repositories, но пусть будет так
	data, err := repositories.GetRecord(int64(idInt))

	jsonResponse, err := json.Marshal(map[string]interface{}{"expression": data})

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
