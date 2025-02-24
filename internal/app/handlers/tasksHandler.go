package handlers

import (
	"encoding/json"
	"net/http"
	"webcalculate/internal/app/repositories"
)

func Task(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		select {

		case val := <-repositories.GlobalCh:

			task := repositories.GlobalDataGetInstance().GetTask(val)
			data := getRespFromTask(task)

			jsonResponse, _ := json.Marshal(map[string]interface{}{"expressions": data})
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)

		default:
			http.Error(w, "error: task not found", http.StatusNotFound)
		}
	} else {
		request := new(TaskResult)
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "error: Expression is not valid", http.StatusUnprocessableEntity)
			return
		}
	}

}
