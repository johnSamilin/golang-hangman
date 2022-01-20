package routes

import (
	i "johnSamilin/golang-hangman/interfaces"
	"johnSamilin/golang-hangman/utils"

	"github.com/google/uuid"
)

// handler for new client. assings unique id and a word
func OnStart(state *i.GameState) (int, string) {
	var clientId = uuid.New().String()
	var word = <-state.Words

	state.Clients[clientId] = i.Client{
		Word:     word,
		Attempts: 1,
		Guesses:  make(map[string]int),
	}
	var resp = &i.NewClientResponse{
		Id:         clientId,
		WordLength: len(word),
	}

	return utils.GetJsonEncodedResponse(resp)
}
