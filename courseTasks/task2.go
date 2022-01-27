package main

import (
	"fmt"
	"johnSamilin/golang-hangman/courseTasks/analyzer"
)

func main() {
	collector := analyzer.New()
	collector.AddLine("Quick brown fox jumped over the lazy dog then jumped again back!")
	fmt.Println(collector.Stats)
}
