package game 

import "fmt"

func PrintWord(initWord []string){
	for i, char := range initWord {
		if i == len(initWord) - 1{
			fmt.Println(char)
		} else {
			print(char)
			print(" ")
		}
	}
}