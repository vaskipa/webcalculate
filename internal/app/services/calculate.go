package services

import (
	"github.com/vaskipa/calculator/calculator"
	"webcalculate/internal/app/repositories"
)

func NewCalculateTask(expression string) (int64, error) {
	polishNotation, err := calculator.ToPolishNotation(expression)
	if err != nil {
		return 0, err
	}

	size := len(polishNotation) - 1
	calcNode := calculator.GenerateAST(polishNotation, &size)
	lastInsertId, err := repositories.AddRecord(expression)

	if err != nil {
		return 0, err
	}

	err = distributedFunc(lastInsertId, calcNode, repositories.GlobalCh)
	return lastInsertId, err
}
