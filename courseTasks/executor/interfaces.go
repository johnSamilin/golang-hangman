package executor

import (
	"johnSamilin/golang-hangman/courseTasks/analyzer"
	"sync"
)

type Worker struct {
	Id   int
	Lock *sync.Mutex
}

type IExecutor interface {
	StartExecution(wg *sync.WaitGroup)
	SetCommands([]string)
	GetStats() analyzer.Stats
	SetWorkersCount(int)
}

type IConfigFilesExecutor interface {
	IExecutor
	Read()
}

type Executor struct {
	workersCount int
	workersPool  chan Worker
	exitSignal   chan bool
	commands     []string
	stats        analyzer.Stats
}

type CmdArgsExecutor struct {
	Executor
}

type ConfigFilesExecutor struct {
	Executor
	filesSource string
	files       []string
}
