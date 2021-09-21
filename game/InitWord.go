package game

import (
	"unicode"

	"hangman.com/notitou/hangman/utils"
)

func InitWord(worde string) []string {
	res := []string{}
	word := []rune(worde)
	randomLetters := RandomLetterIndex(worde)
	for i := 0; i < utils.Len(worde); i++ {
		if utils.ContainsSliceInt(randomLetters, i) || rune(word[i]) == '-' {
			res = append(res, string(unicode.ToUpper(rune(word[i]))))
		} else {
			res = append(res, "_")
		}
	}
	return res
}
