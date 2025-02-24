package repositories

import (
	"fmt"
	"sync"
)

type ExpressionTask struct {
	KeyId int64
	Id    int64 `json:"id"`

	IsHead bool `json:"isHead"`
	IsLeft bool `json:"isLeft"`

	HeadKeyId int64 `json:"headId"`

	// определяет, готов ли arg1
	IsArg1Ready bool    `json:"isArg1Ready"`
	ArgLeft     float64 `json:"arg1"`

	// определяет, готов ли arg2
	IsArg2Ready bool    `json:"isArg2Ready"`
	Arg2        float64 `json:"arg2"`

	isError bool

	Operation     string `json:"operation"`
	OperationTime int    `json:"operationTime"`
}

type GlobalData struct {
	// словарь для того, чтобы внутри хранить все-все задания
	Data map[int64]ExpressionTask
	// индекс счетчик, чтобы генерировать id
	Id int64
	// чтобы из разных потоков добавлять-сохранять
	sync.Mutex
}

var globalDataInstance *GlobalData = nil
var once sync.Once

func NewGlobalData() *GlobalData {
	return &GlobalData{make(map[int64]ExpressionTask), 0, sync.Mutex{}}
}

func GlobalDataGetInstance() *GlobalData {
	once.Do(func() {
		globalDataInstance = NewGlobalData()
	})
	return globalDataInstance
}

func (p *GlobalData) AddTask() *ExpressionTask {
	p.Lock()
	task := new(ExpressionTask)
	// тут нужно прокидывать из глобальынх переменных
	task.OperationTime = 1000
	task.KeyId = p.Id

	p.Data[task.KeyId] = *task
	p.Id++
	p.Unlock()
	fmt.Println("Успешно создана таска", p.Id)
	return task
}

func (p *GlobalData) GetTask(keyId int64) ExpressionTask {
	return p.Data[keyId]
}

// волшебная константа
var GlobalCh chan int64 = make(chan int64, 20)
