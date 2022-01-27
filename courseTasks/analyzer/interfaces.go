package analyzer

type iAnalyzer interface {
	GetFromFile()
	AddLine()
	Restart()
	GetText()
	GetStats()
}
