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
	cExit := make(chan bool, 1)
	if cmdSource == constants.SOURCE_ARGS {
		e = executor.CreateCmdArgsExecutor(cExit, workersCount)
		e.SetCommands(commands)
	} else {
		e = executor.CreateConfigFilesExecutor(cExit, workersCount, constants.CONFIGS_FOLDER)
	}

	go func() {
		for {
			select {
			case <-cExit:
				os.Exit(2)
			}
		}
	}()

	var wg sync.WaitGroup
	e.StartExecution(&wg)
	wg.Wait()
	fmt.Println("-----------")
	fmt.Printf("execution of %s has finished. The most common command was %s (%d times)\n", cmdSource, e.GetStats().TopWord, e.GetStats().TopWordOccurrences)
	fmt.Println("-----------")
}
