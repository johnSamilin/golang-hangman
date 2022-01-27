package executor

import (
	"fmt"
	"strconv"
	"sync"
)

func CreateCmdArgsExecutor(workers int) IExecutor {
	ex := &CmdArgsExecutor{
		Executor: Executor{
			workersCount: workers,
			workersPool:  make(chan Worker, workers),
		},
	}
	for i := 0; i < workers; i++ {
		ex.workersPool <- Worker{Id: i, Lock: &sync.Mutex{}, IsActive: true}
	}

	return ex
}

func (e *CmdArgsExecutor) SetCommands(cmds []string) {
	e.commands = cmds
}

func (e *CmdArgsExecutor) StartExecution(wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Args mode")
	fmt.Println(e)
	e.stats = start(0, e.workersPool, e.commands, wg)
}

func (e *CmdArgsExecutor) GetStats() string {
	return "The most common command was " + e.stats.TopWord + " (" + strconv.Itoa(e.stats.TopWordOccurrences) + " times)"
}
