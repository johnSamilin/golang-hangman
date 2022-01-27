package utils

import (
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"strconv"
	"strings"
)

func ParseArgs(args []string) (workersCount int, cmdSource string, commands []string) {
	workersCount = 3
	cmdSource = constants.SOURCE_FILES
	commands = []string{}
	for _, arg := range args {
		argParsed := strings.Split(arg, "=")
		if argParsed[0] == "--workers" {
			count, err := strconv.Atoi(argParsed[1])
			if err == nil && count > 0 {
				workersCount = count
			}
			continue
		}
		if argParsed[0] == "--src" && argParsed[1] == constants.SOURCE_ARGS {
			cmdSource = constants.SOURCE_ARGS
			continue
		}
		if argParsed[0] == "--src" {
			cmdSource = constants.SOURCE_FILES
			continue
		}
		commands = append(commands, arg)
	}
	return
}
