package handlers

import (
	"webcalculate/internal/app/repositories"
)

type ExpressionJson struct {
	Expression string `json:"expression"`
}

type CalculatorJson struct {
	Id int64 `json:"KeyId"`
}

type ExpressionTaskResponse struct {
	Id int64 `json:"id"`

	Arg1 float64 `json:"arg1"`

	Arg2 float64 `json:"arg2"`

	Operation     string `json:"operation"`
	OperationTime int    `json:"operation_time"`
}

func getRespFromTask(task repositories.ExpressionTask) ExpressionTaskResponse {
	return ExpressionTaskResponse{Id: task.KeyId, Arg1: task.ArgLeft, Arg2: task.Arg2, Operation: task.Operation, OperationTime: task.OperationTime}
}

type TaskResult struct {
	Id     int64   `json:"id"`
	Result float64 `json:"result"`
}
