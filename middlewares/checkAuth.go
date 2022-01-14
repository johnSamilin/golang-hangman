package middlewares

import (
	i "johnSamilin/golang-hangman/interfaces"
	"net/http"
)

// every call to makeGuess must have auth header (X-Client-Id)
func CheckAuth(state *i.GameState, res http.ResponseWriter, req *http.Request) {
	var clientId = req.Header.Get("X-Client-Id")
	var _, clientIdExists = state.Clients[clientId]
	if !clientIdExists {
		res.WriteHeader(http.StatusForbidden)
	}
}
