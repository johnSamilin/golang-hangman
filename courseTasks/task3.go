package main

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/executor"
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"johnSamilin/golang-hangman/courseTasks/utils"
	"os"
	"sync"
)

func main() {
	workersCount, cmdSource, commands := utils.ParseArgs(os.Args[1:])
	var e executor.IExecutor
	if cmdSource == constants.SOURCE_ARGS {
		e = executor.CreateCmdArgsExecutor(workersCount)
		e.SetCommands(commands)
	} else {
		e = executor.CreateConfigFilesExecutor(workersCount, constants.CONFIGS_FOLDER)
	}

	var wg sync.WaitGroup
	e.StartExecution(&wg)
	wg.Wait()
	fmt.Println("-----------")
	stats := e.GetStats()
	fmt.Printf("execution of %s has finished. %s\n", cmdSource, stats)
	fmt.Println("-----------")
}
