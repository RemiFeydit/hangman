package game

import (
	"fmt"
	
	"hangman.com/notitou/hangman/utils"
)

func WinCon(userAnswer *string, wordToSearch *string, initWord *[]string) bool {
	if *userAnswer == *wordToSearch || !(utils.ContainsSliceStr(*initWord, "_")){
		PrintWord(*initWord)
		fmt.Println("GG")
		fmt.Printf("The word was \"%s\"\n", *wordToSearch)
		return true
	}
	return false
}