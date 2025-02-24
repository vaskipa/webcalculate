package internal

import (
	"net/http"
	"webcalculate/internal/app/handlers"
	"webcalculate/internal/app/repositories"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) RunServer() error {
	repositories.InitTables()

	http.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	http.HandleFunc("/api/v1/expressions/{id}", handlers.ExpressionHandler)
	http.HandleFunc("/api/v1/expressions", handlers.ExpressionsHandler)
	return http.ListenAndServe("0.0.0.0:80", nil)
}
