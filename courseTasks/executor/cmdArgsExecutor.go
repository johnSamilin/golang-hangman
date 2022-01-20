package executor

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/analyzer"
	"sync"
)

func CreateCmdArgsExecutor(exitSig chan bool, workers int) IExecutor {
	ex := &CmdArgsExecutor{
		Executor: Executor{
			workersCount: workers,
			workersPool:  make(chan Worker, workers),
			exitSignal:   exitSig,
		},
	}
	for i := 0; i < workers; i++ {
		ex.workersPool <- Worker{Id: i, Lock: &sync.Mutex{}}
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
	e.stats = startExecution(0, e.workersPool, e.commands, e.exitSignal, wg)
}

func (e *CmdArgsExecutor) GetStats() analyzer.Stats {
	return e.stats
}
