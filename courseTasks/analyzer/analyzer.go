package analyzer

import (
	"errors"
	"johnSamilin/golang-hangman/courseTasks/utils"
	"regexp"
	"strings"
)

func (a *Analyzer) AddLine(line string) {
	a.text = append(a.text, line)
	a.analyze(line)
}

func (a *Analyzer) GetFromFile(folder string, name string) error {
	var content, exists = utils.ReadFile(folder, name)
	for _, line := range content {
		a.AddLine(line)
	}

	if !exists {
		return errors.New("no such file")
	}

	return nil
}

func (a *Analyzer) analyze(text string) {
	stripMultipleWS := regexp.MustCompile(`(\s)+`)
	textModified := stripMultipleWS.ReplaceAllString(text, " ")

	words := strings.Split(textModified, " ")
	a.stats.WhitespacesCount = len(words) - 1
	for _, word := range words {
		prevCount, exists := a.words[word]
		if !exists {
			a.words[word] = 1
		}
		a.words[word] = prevCount + 1
		if prevCount+1 > a.stats.TopWordOccurrences {
			a.stats.TopWordOccurrences = prevCount + 1
			a.stats.TopWord = word
		}

		a.stats.CharCount = a.stats.CharCount + len(word)
		for _, letter := range word {
			prevLetterCount, exists := a.chars[letter]
			if !exists {
				a.chars[letter] = 1
			}
			a.chars[letter] = prevLetterCount + 1
			if prevLetterCount+1 > a.stats.TopSymbolOccurrences {
				a.stats.TopSymbolOccurrences = prevLetterCount + 1
				a.stats.TopSymbol = letter
			}
			_, exist := NON_LETTER_CHARS[letter]
			if exist {
				a.stats.NonLettersCount = a.stats.NonLettersCount + 1
			}
		}
	}
}

func (a *Analyzer) Restart() {
	a.stats = Stats{
		CharCount:            0,
		TopWord:              "",
		TopWordOccurrences:   0,
		NonLettersCount:      0,
		WhitespacesCount:     0,
		TopSymbol:            ' ',
		TopSymbolOccurrences: 0,
	}
	a.text = []string{}
	a.words = make(map[string]int)
	a.chars = make(map[rune]int)
}

func (a *Analyzer) GetText() string {
	return strings.Join(a.text, "\n")
}

func (a *Analyzer) GetStats() Stats {
	return a.stats
}

func New() Analyzer {
	var instance = Analyzer{}
	instance.Restart()

	return instance
}
