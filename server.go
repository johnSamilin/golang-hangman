package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/pborman/uuid"
)

type Client struct {
	word     string         // the word that associated with the client
	attempts int            // number of guess attempts that already done
	guesses  map[string]int // map with the letters that were guessed
}

type GameState struct {
	words   chan string       // words buffer
	clients map[string]Client // currently playing clients
}

type NewClientResponse struct {
	Id         string `json:"id"`         // unique id
	WordLength int    `json:"wordLength"` // how many letters in the word to guess
}

type GuessRequest struct {
	Letter string `json:"letter"`
}

type GuessResponse struct {
	LetterExist  bool   `json:"letterExist"`  // is the suggested letter in the word
	AttemptsLeft int    `json:"attemptsLeft"` // how many attempts left
	Word         string `json:"word"`         // masked word. unknown letters are represented as _
	Won          bool   `json:"won"`          // game result
}

const MAX_WORDS = 5
const MIN_WORDS_COUNT = 2
const WORDS_SOURCE = "https://random-word-api.herokuapp.com/word?number="

// checks if words count less than certain amount and requests some more
var isRequestingMoreWords = false

func checkForWordsQueue(state *GameState) {
	if len(state.words) < MIN_WORDS_COUNT && !isRequestingMoreWords {
		isRequestingMoreWords = true
		fmt.Println("requesting more words...")
		var resp, err = http.Get(WORDS_SOURCE + strconv.Itoa(MAX_WORDS-len(state.words)))
		if err != nil {
			state.words <- "Unknown"
			return
		}

		defer resp.Body.Close()
		var data, er = ioutil.ReadAll(resp.Body)
		if er != nil {
			state.words <- "Unknown"
			return
		}
		var newWords []string
		json.Unmarshal(data, &newWords)
		for _, w := range newWords {
			state.words <- w
		}

		isRequestingMoreWords = false
	}
}

// handler for new client. assings unique id and a word
func newClient(state *GameState) (int, string) {
	var clientId = uuid.New()
	var word string
	select {
	// in case there are no words, wait until they are downloaded
	case word = <-state.words:
	case <-time.After(time.Second * 10):
		word = "Unknown"
	}
	// asyncronously load words if nessesary
	go checkForWordsQueue(state)

	state.clients[clientId] = Client{
		word:     word,
		attempts: 0,
		guesses:  make(map[string]int),
	}
	var resp = &NewClientResponse{
		Id:         clientId,
		WordLength: len(word),
	}
	var data, err = json.Marshal(resp)
	if err != nil {
		fmt.Println("ERR", err)
		return 500, err.Error()
	}

	return 200, string(data)
}

// reveal letters that were guessed correctly
func getMaskedWord(wrd string, triedLetters map[string]int) string {
	return strings.Map(func(r rune) rune {
		var _, exists = triedLetters[strings.ToLower(string(r))]
		if exists {
			return r
		}

		return rune('_')
	}, wrd)
}

// handler for guessing. request body must contain { letter: 'j' } structure
func makeGuess(state *GameState, req *http.Request) (int, string) {
	var clientId = req.Header.Get("X-Client-Id")

	defer req.Body.Close()
	var data, _ = ioutil.ReadAll(req.Body)
	jsonResult := GuessRequest{}
	var err = json.Unmarshal(data, &jsonResult)
	if err != nil {
		return 500, "Unable to parse request"
	}

	var word = state.clients[clientId].word
	var letter = strings.ToLower(jsonResult.Letter)
	var guesses = state.clients[clientId].guesses
	var attemptsLeft = len(word)/2 - state.clients[clientId].attempts
	var isGuessCorrect = strings.Contains(strings.ToLower(word), letter)

	var response []byte
	guesses[letter] = 1
	var maskedWord = getMaskedWord(word, guesses)

	if attemptsLeft == 0 {
		delete(state.clients, clientId)
		// return word into stack
		state.words <- word

		response, _ = json.Marshal(GuessResponse{
			LetterExist:  isGuessCorrect,
			AttemptsLeft: 0,
			Word:         maskedWord,
			Won:          false,
		})
	}
	// else
	if attemptsLeft != 0 {
		var newAttemptsCount = state.clients[clientId].attempts
		if !isGuessCorrect {
			// successful guess does not increment attempts count
			// thus, you may make up to len(word)/2 mistakes
			newAttemptsCount = newAttemptsCount + 1
		}
		state.clients[clientId] = Client{
			word:     word,
			attempts: newAttemptsCount,
			guesses:  guesses,
		}
		response, _ = json.Marshal(GuessResponse{
			LetterExist:  isGuessCorrect,
			AttemptsLeft: attemptsLeft,
			Word:         maskedWord,
			Won:          !strings.Contains(maskedWord, "_"),
		})
	}

	return 200, string(response)
}

// every call to makeGuess must have auth header (X-Client-Id)
func checkAuth(state *GameState, res http.ResponseWriter, req *http.Request) {
	var clientId = req.Header.Get("X-Client-Id")
	var _, clientIdExists = state.clients[clientId]
	if !clientIdExists {
		res.WriteHeader(http.StatusForbidden)
	}
}

func main() {
	var server = martini.Classic()
	var state = &GameState{
		words:   make(chan string, MAX_WORDS),
		clients: make(map[string]Client),
	}

	// default words
	state.words <- "Golang"
	state.words <- "JavaScript"
	state.words <- "React"
	state.words <- "WebGL"
	state.words <- "Friendship"
	// go checkForWordsQueue(state)

	// Injection
	server.Map(state)

	server.Get("/start", newClient)
	server.Post("/guess", checkAuth, makeGuess)

	// UI
	server.Use(martini.Static("public"))

	server.Run()
}
