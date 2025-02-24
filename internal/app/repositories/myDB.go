package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

func InitTables() error {
	db, err := sql.Open("sqlite3", "./server.sqlite")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return errors.New("error")
	}
	defer db.Close()
	query := `
        CREATE TABLE IF NOT EXISTS tasks (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Task TEXT,
			Status TEXT DEFAULT 'not ready',
			Result float DEFAULT ''
        );
    `
	_, err = db.Exec(query)
	return err
}

var mutex sync.Mutex

func AddRecord(task string) (int64, error) {
	mutex.Lock()
	db, err := sql.Open("sqlite3", "./server.sqlite")
	defer mutex.Unlock()
	defer db.Close()
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return 0, errors.New("error")
	}
	query := `
        INSERT INTO tasks (Task)
        VALUES (?)
    `

	result, err := db.Exec(query, task)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	return lastInsertId, err
}

type Record struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Result string `json:"result"`
	Status string `json:"status"`
}

func GetRecords() ([]Record, error) {
	mutex.Lock()
	db, err := sql.Open("sqlite3", "./server.sqlite")
	rows, err := db.Query("SELECT Id, Task, Result, Status FROM tasks")
	mutex.Unlock()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.Id, &record.Task, &record.Result, &record.Status); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func GetRecord(id int64) (Record, error) {
	mutex.Lock()
	db, err := sql.Open("sqlite3", "./server.sqlite")
	task := new(Record)
	err = db.QueryRow("SELECT Id, Task, Result, Status FROM tasks WHERE Id = ?", id).Scan(&task.Id, &task.Task, &task.Status, &task.Result)

	mutex.Unlock()

	return *task, err

}

func UpdateTask(id int64, result float64) {

	mutex.Lock()
	db, _ := sql.Open("sqlite3", "./server.sqlite")
	query := `
        UPDATE tasks
        SET Status = "ready", result = ?
        WHERE id = ?
	`

	_, _ = db.Exec(query, result)
	mutex.Unlock()
}
