package game

import "hangman.com/notitou/hangman/utils"

func InitWord(word string) []string {
	res := []string{}
	randomLetters := RandomLetterIndex(word)
	for i := 0; i < len(word); i++ {
		if utils.ContainsSliceInt(randomLetters, i){
			res = append(res, string(word[i]))
		} else {
			res = append(res, "_")
		}
	}
	return res
}
