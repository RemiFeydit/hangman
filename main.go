package main

import (
	"fmt"

	"hangman.com/notitou/hangman/functions"
)

func main() {
	wordToSearch := functions.RandomWord()
	initWord := functions.InitWord(wordToSearch)
	attempt := 10
	userAnswer := ""
	for true {
		if attempt == 0 {
			fmt.Println("Game Over!!!")
			return
		} else {
			userAnswer = functions.AskUser(attempt)
			attempt--
		}
	}

	fmt.Println(userAnswer)
	fmt.Printf("%v", wordToSearch)
	fmt.Printf("%v", initWord)
}
