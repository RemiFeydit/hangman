package game 

import "fmt"

func PrintWord(initWord []string){
	fmt.Print("\033[H\033[2J")
	for i, char := range initWord {
		if i == len(initWord) - 1{
			fmt.Println(char)
		} else {
			print(char)
			print(" ")
		}
	}
}