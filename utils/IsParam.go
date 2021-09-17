package utils

import (
	"os"
)

func IsParam() bool {
	if len(os.Args[1:]) != 1 {
		return false
	}
	return true
}