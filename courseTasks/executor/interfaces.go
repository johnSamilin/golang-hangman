package executor

import (
	"fmt"
	"sync"
)

type IExecutor interface {
	StartExecution(wg *sync.WaitGroup)
	SetCommands([]string)
	GetStats() string
}

type IConfigFilesExecutor interface {
	IExecutor
	Read()
}

type ITask interface {
	fmt.Stringer
	Execute() (bool, string)
}
