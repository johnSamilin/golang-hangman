package executor

import (
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"johnSamilin/golang-hangman/courseTasks/utils"
	"os"
	"strings"
	"time"
)

func (t Task) String() string {
	return t.command
}

func (t *TaskTime) Execute() (bool, string) {
	return true, time.Now().String()
}

func (t *TaskTimeWait) Execute() (bool, string) {
	time.Sleep(time.Second)
	return true, time.Now().String()
}

func (t *TaskShutDown) Execute() (bool, string) {
	return false, ""
}

func (t *TaskExit) Execute() (bool, string) {
	os.Exit(2)
	return false, ""
}

func (t *TaskRead) Execute() (bool, string) {
	var fileContents []string
	var exists bool
	fileContents, exists = utils.ReadFile(constants.OUTPUT_FOLDER, t.filename)
	if !exists {
		err := os.WriteFile(constants.OUTPUT_FOLDER+"/"+t.filename, []byte{}, 0644)
		if err != nil {
			fileContents = append(fileContents, "FAILURE")
		}
	}
	return true, strings.Join(fileContents, " ")
}

func (t *TaskWrite) Execute() (bool, string) {
	var fileContents []string
	fileContents, _ = utils.ReadFile(constants.OUTPUT_FOLDER, t.filename)
	fileContents = append(fileContents, t.content)
	var content []byte
	for lineNum, line := range fileContents {
		for _, letter := range line {
			content = append(content, byte(letter))
		}
		if lineNum < len(fileContents)-1 {
			content = append(content, byte('\n'))
		}
	}
	err := os.WriteFile(constants.OUTPUT_FOLDER+"/"+t.filename, content, 0644)
	if err != nil {
		return true, "Failed to write to file: " + err.Error()
	}
	if err == nil {
		return true, "Done"
	}

	return true, ""
}
