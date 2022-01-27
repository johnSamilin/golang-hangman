package executor

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/utils"
	"strconv"
	"sync"
)

func CreateConfigFilesExecutor(workers int, filesSource string) IConfigFilesExecutor {
	ex := &ConfigFilesExecutor{
		Executor: Executor{
			workersCount: workers,
			workersPool:  make(chan Worker, workers),
		},
		filesSource: filesSource,
	}
	for i := 0; i < workers; i++ {
		ex.workersPool <- Worker{Id: i, Lock: &sync.Mutex{}, IsActive: true}
	}

	return ex
}

func (e *ConfigFilesExecutor) Read() {
	e.files = utils.ReadConfigs(e.filesSource)
}

func (e *ConfigFilesExecutor) SetCommands(cmds []string) {
	e.commands = cmds
}

func (e *ConfigFilesExecutor) StartExecution(wg *sync.WaitGroup) {
	fmt.Println("Files mode (each config is being executed by number of workers)")
	e.Read()
	for idx, name := range e.files {
		wg.Add(1)
		var commands, _ = utils.ReadFile(e.filesSource, name)
		e.SetCommands(commands)
		stats := start(idx, e.workersPool, commands, wg)
		if e.stats.TopWordOccurrences < stats.TopWordOccurrences {
			e.stats.TopWordOccurrences = stats.TopWordOccurrences
			e.stats.TopWord = stats.TopWord
		}
		if e.stats.TopSymbolOccurrences < stats.TopSymbolOccurrences {
			e.stats.TopSymbolOccurrences = stats.TopSymbolOccurrences
			e.stats.TopSymbol = stats.TopSymbol
		}
		e.stats.CharCount = e.stats.CharCount + stats.CharCount
		e.stats.NonLettersCount = e.stats.NonLettersCount + stats.NonLettersCount
		e.stats.WhitespacesCount = e.stats.WhitespacesCount + stats.WhitespacesCount
	}
}

func (e *ConfigFilesExecutor) GetStats() string {
	return "The most common command was " + e.stats.TopWord + " (" + strconv.Itoa(e.stats.TopWordOccurrences) + " times)"
}
