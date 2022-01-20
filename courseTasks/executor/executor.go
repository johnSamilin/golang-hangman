package executor

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/analyzer"
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"johnSamilin/golang-hangman/courseTasks/utils"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func startExecution(configId int, workers chan Worker, commands []string, exit chan bool, wg *sync.WaitGroup) analyzer.Stats {
	var tasksPool = make(chan string, len(commands))
	collector := analyzer.New()
	for _, cmd := range commands {
		tasksPool <- cmd
		collector.AddLine(cmd)
	}
	for {
		freeWorker := <-workers
		if len(tasksPool) > 0 {
			shouldStop := processCommand(freeWorker.Id, freeWorker.Lock, configId, <-tasksPool, exit)
			if !shouldStop {
				workers <- freeWorker
			}
			if shouldStop {
				break
			}
		}
		if len(tasksPool) == 0 {
			close(tasksPool)
			wg.Done()
			break
		}
	}
	return collector.GetStats()
}

func processCommand(idx int, lock *sync.Mutex, configId int, cmd string, exit chan bool) (shouldStop bool) {
	shouldStop = false
	var args = strings.SplitN(cmd, " ", 3)
	lock.Lock()
	commonInfo := "Worker #" + strconv.Itoa(idx) + " (task from cmd pool #" + strconv.Itoa(configId) + ")"
	switch args[0] {
	case constants.Time:
		fmt.Println(commonInfo, "Time:", time.Now())
		lock.Unlock()
	case constants.TimeWait:
		go func(info string, lock *sync.Mutex) {
			fmt.Println(commonInfo, "Time (wait):", time.Now())
			lock.Unlock()
		}(commonInfo, lock)
	case constants.ShutDown:
		fmt.Println(commonInfo, "Stopped")
		shouldStop = true
		lock.Unlock()
	case constants.Exit:
		fmt.Println(commonInfo, "Exits an application")
		shouldStop = true
		lock.Unlock()
		exit <- true
	case constants.Read:
		var fileContents []string
		var exists bool
		fileContents, exists = utils.ReadFile(constants.OUTPUT_FOLDER, args[1])
		if !exists {
			err := os.WriteFile(constants.OUTPUT_FOLDER+"/"+args[1], []byte{}, 0644)
			if err != nil {
				fileContents = append(fileContents, "FAILURE")
			}
		}
		fmt.Println(commonInfo, "Read file", args[1], ":", fileContents)
		lock.Unlock()
	case constants.Write:
		var fileContents []string
		fileContents, _ = utils.ReadFile(constants.OUTPUT_FOLDER, args[1])
		fileContents = append(fileContents, args[2])
		var content []byte
		for lineNum, line := range fileContents {
			for _, letter := range line {
				content = append(content, byte(letter))
			}
			if lineNum < len(fileContents)-1 {
				content = append(content, byte('\n'))
			}
		}
		err := os.WriteFile(constants.OUTPUT_FOLDER+"/"+args[1], content, 0644)
		if err != nil {
			fmt.Println(commonInfo, "Failed to write to file: "+err.Error())
		}
		if err == nil {
			fmt.Println(commonInfo, "Write to file", args[1], "\"", args[2], "\"")
		}
		lock.Unlock()
	default:
		lock.Unlock()
	}

	return
}
