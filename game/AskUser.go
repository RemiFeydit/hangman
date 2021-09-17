package game

import (
	"fmt"
	"hangman.com/notitou/hangman/utils"
)

func AskUser(numberAttempt *int, wordToSearch *string, initWord *[]string, userAnswer *string) {
	var answer string
	fmt.Printf("You have %d attempt left\n", *numberAttempt)
	fmt.Scanln(&answer)
	*userAnswer = answer
	if utils.ContainsStr(*wordToSearch, answer){
		for i, char := range *wordToSearch {
			if string(char) == answer {
				(*initWord)[i] = string(char)
			}
		}
	}else{
		*numberAttempt--
	}
}
