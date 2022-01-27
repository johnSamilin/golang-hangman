package executor

import (
	"johnSamilin/golang-hangman/courseTasks/analyzer"
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"strings"
	"sync"
)

func createTask(command string) ITask {
	var args = strings.SplitN(command, " ", 3)
	switch args[0] {
	case constants.Time:
		return &TaskTime{Task: Task{
			command: command,
		}}
	case constants.TimeWait:
		return &TaskTimeWait{Task: Task{
			command: command,
		}}
	case constants.ShutDown:
		return &TaskShutDown{Task: Task{
			command: command,
		}}
	case constants.Exit:
		return &TaskExit{Task: Task{
			command: command,
		}}
	case constants.Read:
		return &TaskRead{Task: Task{
			command: command,
		},
			filename: args[1],
		}
	case constants.Write:
		return &TaskWrite{Task: Task{
			command: command,
		},
			filename: args[1],
			content:  args[2],
		}
	default:
		panic("Unknown instruction " + command)
	}
}

func start(configId int, workers chan Worker, commands []string, wg *sync.WaitGroup) analyzer.Stats {
	var tasksPool = make(chan ITask, len(commands))
	collector := analyzer.New()
	for _, cmd := range commands {
		tasksPool <- createTask(cmd)
		collector.AddLine(cmd)
	}
	for {
		freeWorker := <-workers
		if len(tasksPool) > 0 {
			go func(f Worker, w chan Worker) {
				freeWorker.Work(configId, <-tasksPool)
				if freeWorker.IsActive {
					workers <- freeWorker
				}
			}(freeWorker, workers)
		}
		if len(tasksPool) == 0 {
			close(tasksPool)
			wg.Done()
			break
		}
	}
	return collector.Stats
}
