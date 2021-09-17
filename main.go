package main

import (
	"hangman.com/notitou/hangman/game"
)

func main() {
	wordToSearch := game.RandomWord()
	initWord := game.InitWord(wordToSearch)
	attempt := 10
	submittedLetter := []string{}
	var userAnswer string
	for true {
		game.AskUser(&attempt, &wordToSearch, &initWord, &userAnswer, &submittedLetter)
		if game.WinCon(&userAnswer, &wordToSearch, &initWord) || game.GameOverCon(&wordToSearch, &attempt){
			return
		}
	}
}
