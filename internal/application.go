package internal

import (
	"encoding/json"
	"github.com/vaskipa/calculator/calculator"
	"net/http"
	"strconv"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	request := new(ExpressionJson)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	polishNotation, err := calculator.ToPolishNotation(request.Expression)
	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)

		return
	}
	size := len(polishNotation) - 1
	_ = calculator.GenerateAST(polishNotation, &size)
	/*
		if err != nil {
			http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
			return
		}
	*/
	lastInsertId, err := addRecord(request.Expression)

	if err != nil {
		http.Error(w, "Проблема с записью в бд", http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(CalculatorJson{lastInsertId})

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func ExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getRecords()
	if err != nil {
		http.Error(w, "error: internal server error", http.StatusInternalServerError)
	}

	jsonResponse, err := json.Marshal(map[string]interface{}{"expressions": data})

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "error: id is not valid", http.StatusUnprocessableEntity)
	}
	data, err := getRecord(int64(idInt))

	jsonResponse, err := json.Marshal(map[string]interface{}{"expression": data})

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Записываем JSON в ответ
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (a *Application) RunServer() error {
	createTable()
	http.HandleFunc("/api/v1/calculate", CalculateHandler)
	http.HandleFunc("/api/v1/expressions/{id}", ExpressionHandler)
	http.HandleFunc("/api/v1/expressions", ExpressionsHandler)
	return http.ListenAndServe("0.0.0.0:80", nil)
}
