package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var attempts int = 5
	var num int = rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("guess number")
		var guess, _, _ = reader.ReadLine()
		var guessN, _ = strconv.Atoi(string(guess))
		if guessN == num {
			attempts = attempts - 1
		}
		if guessN == num {
			fmt.Println("You won")
			break
		}
		if attempts == 0 {
			fmt.Println("You're damn loser")
			break
		}
	}
}
