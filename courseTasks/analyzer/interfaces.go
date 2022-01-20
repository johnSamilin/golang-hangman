package analyzer

type Stats struct {
	CharCount            int
	TopWord              string
	TopWordOccurrences   int
	TopSymbol            rune
	TopSymbolOccurrences int
	NonLettersCount      int
	WhitespacesCount     int
}

type Analyzer struct {
	iAnalyzer
	stats Stats
	text  []string
	words map[string]int
	chars map[rune]int
}

type iAnalyzer interface {
	GetFromFile()
	AddLine()
	Restart()
	GetText()
	GetStats()
}
