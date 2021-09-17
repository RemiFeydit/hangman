package game

import (
	"hangman.com/notitou/hangman/utils"
	"math/rand"
	"time"
	"fmt"
)

func RandomLetterIndex(word string) []int{
	n := len(word) / 2 - 1
	randomIndexSlice := []int{}
	var randWordIndex int
	rand.Seed(time.Now().UnixNano())
	for i:= 0; i <= n; i++ {
		randWordIndex = rand.Intn(len(word))
		if utils.ContainsSliceInt(randomIndexSlice, randWordIndex) {
			i--
		}else{
			randomIndexSlice = append(randomIndexSlice, randWordIndex)
		}
	}
	fmt.Printf("%v", randomIndexSlice)
	return randomIndexSlice
}