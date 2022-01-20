package analyzer

var NON_LETTER_CHARS = map[rune]bool{
	rune('.'):  true,
	rune(','):  true,
	rune('('):  true,
	rune(')'):  true,
	rune('!'):  true,
	rune('\\'): true,
	rune('/'):  true,
	rune('@'):  true,
	rune('\''): true,
	rune('"'):  true,
}
