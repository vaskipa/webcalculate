package services

import (
	"github.com/vaskipa/calculator/calculator"
	"webcalculate/internal/app/handlers"
	"webcalculate/internal/app/repositories"
)

func distributedFunc(id int64, calcNode *calculator.Node, chTaskId chan<- int64) error {

	gd := repositories.GlobalDataGetInstance()
	head := gd.AddTask()
	head.IsHead = true
	head.Id = id

	err := toCalcList(head, calcNode.Left, true, chTaskId)
	if err != nil {
		return err
	}
	err = toCalcList(head, calcNode.Right, false, chTaskId)
	if err != nil {
		return err
	}
	if head.IsArg1Ready && head.IsArg2Ready {
		chTaskId <- head.Id
	}

	return nil
}

func toCalcList(headTask *repositories.ExpressionTask, calcNode *calculator.Node, isLeft bool, chTaskId chan<- int64) error {
	if !calcNode.IsOperation {
		if isLeft {
			headTask.ArgLeft = float64(calcNode.Digit)
			headTask.IsArg1Ready = true
		} else {
			headTask.Arg2 = float64(calcNode.Digit)
			headTask.IsArg2Ready = true
		}
		return nil
	}

	gd := repositories.GlobalDataGetInstance()
	node := gd.AddTask()
	node.HeadKeyId = headTask.KeyId

	err := toCalcList(node, calcNode.Left, true, chTaskId)
	if err != nil {
		return err
	}
	err = toCalcList(node, calcNode.Right, false, chTaskId)
	if err != nil {
		return err
	}

	if node.IsArg1Ready && node.IsArg2Ready {
		chTaskId <- node.Id
	}
	return nil
}

func upDateTasks(result handlers.TaskResult) error {
	gd := repositories.GlobalDataGetInstance()

	task := gd.GetTask(result.Id)

	if task.IsHead {
		repositories.UpdateTask(result.Id, result.Result)
		return nil
	}

	head := gd.GetTask(task.HeadKeyId)
	if task.IsLeft {
		head.ArgLeft = result.Result
		head.IsArg1Ready = true
	} else {
		head.Arg2 = result.Result
		head.IsArg2Ready = true
	}
	// не знаю, нужно ли обкладывать мутексом
	if head.IsArg1Ready && head.IsArg2Ready {
		repositories.GlobalCh <- head.KeyId
	}

	return nil
}
