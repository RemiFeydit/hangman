package game

import "fmt"

func GameOverCon(wordToSearch *string, attempt *int) bool {
	if *attempt <= 0 {
		fmt.Println("Game Over!!!")
		fmt.Printf("The word u were looking for was \"%s\"\n", *wordToSearch)
		return true
	}
	return false
}
