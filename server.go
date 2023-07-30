package main

import (
	i "johnSamilin/golang-hangman/interfaces"
	"johnSamilin/golang-hangman/middlewares"
	"johnSamilin/golang-hangman/routes"
	"johnSamilin/golang-hangman/utils"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
)

func checkForWordsQueue(state *i.GameState) {
	if len(state.Words) < MIN_WORDS_COUNT {
		var newWords = utils.GetWords(MAX_WORDS - len(state.Words))
		for _, w := range newWords {
			state.Words <- w
		}
	}
}

func main() {
	var server = martini.Classic()
	var state = &i.GameState{
		Words:   make(chan string, MAX_WORDS),
		Clients: make(map[string]i.Client),
	}

	go func() {
		var ticker = time.NewTicker(time.Second * WORDS_QUEUE_POLL_INTERVAL)
		checkForWordsQueue(state)
		for range ticker.C {
			checkForWordsQueue(state)
		}
	}()

	server.Map(state)
	server.Use(cors.Allow(CORS_OPTIONS))

	server.Get("/start", routes.OnStart)
	server.Post("/guess", middlewares.CheckAuth, routes.MakeGuess)
	server.Run()
}
