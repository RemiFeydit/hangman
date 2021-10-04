package game

import (
	"fmt"
	"unicode"

	"hangman.com/notitou/hangman/utils"
)

func AskUser(numberAttempt *int, wordToSearch *string, initWord *[]string, userAnswer *string, submittedLetter *[]string) {
	var answer string
	var accentChar []string
	PrintWord(*initWord)
	if *numberAttempt == 10 {
		fmt.Printf("You have %d attempt left\n", *numberAttempt)
	}
	fmt.Print("Choose : ")
	fmt.Scanln(&answer)
	fmt.Print("\033[H\033[2J")
	*userAnswer = answer
	answerToRune := []rune(answer)
	if len(answer) == 0 {
		fmt.Println("You can't submit an empty answer")
		AskUser(numberAttempt, wordToSearch, initWord, userAnswer, submittedLetter)
		return
	}
	if utils.ContainsSliceStr(*submittedLetter, string(unicode.ToLower(rune(answer[0])))) && utils.Len(answer) == 1 {
		fmt.Printf("You already submitted the letter %v, try again\n", answer)
		ReadFile(9 - *numberAttempt)
	} else if !(utils.ContainsStr(*wordToSearch, string(unicode.ToLower(rune(answer[0]))))) && utils.Len(answer) == 1 {
		*numberAttempt--
		fmt.Printf("Not present in the word, %v attempts remaining\n", *numberAttempt)
		ReadFile(9 - *numberAttempt)
	} else if utils.Len(answer) > 1 && answer != *wordToSearch {
		*numberAttempt -= 2
		if *numberAttempt < 0 {
			*numberAttempt = 0
		}
		fmt.Printf("%v is not the right word, %v attempts remaining\n", answer, *numberAttempt)
	} else {
		wordToSearchRune := []rune(*wordToSearch)
		for i, char := range wordToSearchRune {
			if len(answerToRune) == 1 {
				switch unicode.ToUpper(answerToRune[0]) {
				case 'E':
					accentChar = []string{"é", "è", "ê"}
				case 'A':
					accentChar = []string{"à", "â", "ä"}
				case 'I':
					accentChar = []string{"i", "ï", "ï"}
				case 'O':
					accentChar = []string{"ö", "ô"}
				case 'U':
					accentChar = []string{"û", "ü", "ù"}
				}
			}

			if utils.ContainsSliceStr(accentChar, string(char)) {
				(*initWord)[i] = string(unicode.ToUpper(char))
			}

		}
		fmt.Printf("Congrats, the letter %v was in the word !\n", answer)
		ReadFile(9 - *numberAttempt)
	}
	if !(utils.ContainsSliceStr(*submittedLetter, answer)) {
		*submittedLetter = append(*submittedLetter, answer)
		*submittedLetter = append(*submittedLetter, accentChar...)
	}
}
