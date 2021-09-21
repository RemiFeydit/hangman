package utils

func Len(s string) int {
	runeSlice := []rune(s)
	return len(runeSlice)
}
