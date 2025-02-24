package handlers

import (
	"encoding/json"
	"net/http"
	"webcalculate/internal/app/services"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	request := new(ExpressionJson)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	lastInsertId, err := services.NewCalculateTask(request.Expression)

	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
		// пока не будем отвечать по этой ошибке
		//http.Error(w, "Проблема с записью в бд", http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(CalculatorJson{lastInsertId})

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
