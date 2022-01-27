package executor

import (
	"johnSamilin/golang-hangman/courseTasks/analyzer"
	"sync"
)

type Worker struct {
	Id       int
	Lock     *sync.Mutex
	IsActive bool
}

type Task struct {
	ITask
	command string
	result  chan string
}

type TaskTime struct {
	Task
}

type TaskTimeWait struct {
	Task
}

type TaskShutDown struct {
	Task
}

type TaskRead struct {
	Task
	filename string
}

type TaskWrite struct {
	Task
	filename string
	content  string
}

type TaskExit struct {
	Task
}

type Executor struct {
	workersCount int
	workersPool  chan Worker
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
