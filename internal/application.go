package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webcalculate/calculate"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	result, err := calculate.Calc(request.Expression)
	if err != nil {
		http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
	} else {
		fmt.Fprintf(w, "result: %f", result)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalculateHandler)
	return http.ListenAndServe("0.0.0.0:8080", nil)
}
