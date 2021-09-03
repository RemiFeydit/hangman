package functions

import (
	"fmt"
)

func AskUser(numberAttempt int) string {
	var answer string
	fmt.Printf("You have %d attempt left\n", numberAttempt)
	fmt.Scanln(&answer)
	return answer
}
