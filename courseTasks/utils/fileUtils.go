package utils

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/executor/constants"
	"os"
	"regexp"
	"strings"
)

func ReadConfigs(source string) []string {
	file, err := os.Open(source)
	configs := []string{}
	if err != nil {
		panic("failed opening directory: " + err.Error())
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)
	for _, name := range list {
		matched, _ := regexp.MatchString(constants.CONFIG_FILENAME_PATTERN, name)
		if matched {
			configs = append(configs, name)
		}
	}

	return configs
}

func ReadFile(folder string, name string) (contents []string, exists bool) {
	file, err := os.ReadFile(folder + "/" + name)
	contents = []string{}
	exists = true
	if err != nil {
		fmt.Println("failed opening config file: " + err.Error() + " (" + name + ")")
		exists = false
		return
	}

	contents = strings.Split(string(file), "\r\n")
	return
}
