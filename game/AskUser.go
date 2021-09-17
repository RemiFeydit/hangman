package game

import (
	"fmt"
	"hangman.com/notitou/hangman/utils"
)

func AskUser(numberAttempt *int, wordToSearch *string, initWord *[]string, userAnswer *string, submittedLetter *[]string) {
	var answer string
	PrintWord(*initWord)
	if *numberAttempt == 10 {
		fmt.Printf("You have %d attempt left\n", *numberAttempt)
	}
	fmt.Print("Choose : ")
	fmt.Scanln(&answer)
	fmt.Print("\033[H\033[2J")
	*userAnswer = answer
	if utils.ContainsSliceStr(*submittedLetter, answer){
		fmt.Printf("You already submitted the letter %v, try again\n", answer)
	} else if !(utils.ContainsStr(*wordToSearch, answer)) && *wordToSearch != answer{
		*numberAttempt--
		fmt.Printf("Not present in the word, %v attempts remaining\n", *numberAttempt)	
	} else {
		for i, char := range *wordToSearch {
			if string(char) == answer {
				(*initWord)[i] = string(char)
			}
		}
		fmt.Printf("Congrats, the letter %v was in the word !\n", answer)
	}
	if !(utils.ContainsSliceStr(*submittedLetter, answer)){
		*submittedLetter = append(*submittedLetter, answer)
	}
}
