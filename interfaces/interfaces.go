package interfaces

type Client struct {
	Word     string         // the word that associated with the client
	Attempts int            // number of guess attempts that already done
	Guesses  map[string]int // map with the letters that were guessed
}

type GameState struct {
	Words   chan string       // words buffer
	Clients map[string]Client // currently playing clients
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
