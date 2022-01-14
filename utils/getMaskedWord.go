package utils

import "strings"

// reveal letters that were guessed correctly
func GetMaskedWord(wrd string, triedLetters map[string]int) string {
	return strings.Map(func(r rune) rune {
		var _, exists = triedLetters[strings.ToLower(string(r))]
		if exists {
			return r
		}

		return rune('_')
	}, wrd)
}
