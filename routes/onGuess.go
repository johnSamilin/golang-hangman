package routes

import (
	i "johnSamilin/golang-hangman/interfaces"
	"johnSamilin/golang-hangman/utils"
	"net/http"
	"strings"
)

// handler for guessing. request body must contain { letter: 'j' } structure
func MakeGuess(state *i.GameState, req *http.Request) (int, string) {
	var clientId = req.Header.Get("X-Client-Id")

	jsonResult := i.GuessRequest{}
	var err = utils.GetDataFromBody(req.Body, jsonResult)
	if err != nil {
		return 500, "Unable to parse request"
	}

	var word = state.Clients[clientId].Word
	var letter = strings.ToLower(jsonResult.Letter)
	var guesses = state.Clients[clientId].Guesses
	var attemptsLeft = len(word) - state.Clients[clientId].Attempts
	var isGuessCorrect = strings.Contains(strings.ToLower(word), letter)

	var response i.GuessResponse
	guesses[letter] = 1
	var maskedWord = utils.GetMaskedWord(word, guesses)

	if attemptsLeft == 0 {
		delete(state.Clients, clientId)
		state.Words <- word

		response = i.GuessResponse{
			LetterExist:  isGuessCorrect,
			AttemptsLeft: 0,
			Word:         maskedWord,
			Won:          false,
		}
	}
	if attemptsLeft != 0 {
		var newAttemptsCount = state.Clients[clientId].Attempts
		if !isGuessCorrect {
			newAttemptsCount = newAttemptsCount + 1
		}
		state.Clients[clientId] = i.Client{
			Word:     word,
			Attempts: newAttemptsCount,
			Guesses:  guesses,
		}
		response = i.GuessResponse{
			LetterExist:  isGuessCorrect,
			AttemptsLeft: attemptsLeft,
			Word:         maskedWord,
			Won:          !strings.Contains(maskedWord, "_"),
		}
	}

	return utils.GetJsonEncodedResponse(response)
}
