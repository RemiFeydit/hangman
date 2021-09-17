package main

import (
	"hangman.com/notitou/hangman/game"
)

func main() {
	wordToSearch := game.RandomWord()
	initWord := game.InitWord(wordToSearch)
	attempt := 10
	var userAnswer string
	for true {
		game.PrintWord(initWord)
		game.AskUser(&attempt, &wordToSearch, &initWord, &userAnswer)
		if game.WinCon(&userAnswer, &wordToSearch, &initWord) || game.GameOverCon(&wordToSearch, &attempt){
			return
		}
	}
}
