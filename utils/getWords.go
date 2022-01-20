package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetWords(n int) []string {
	var resp, err = http.Get(WORDS_SOURCE + strconv.Itoa(n))
	var newWords = []string{}

	if err != nil {
		fmt.Println("Error while requesting more words", err.Error())
		return []string{"Unknown"}
	}
	var er = GetDataFromBody(resp.Body, &newWords)
	if er != nil {
		fmt.Println("Error while parsing more words", er.Error())
		return []string{"Unknown"}
	}

	return newWords
}
