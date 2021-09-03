package functions

func InitWord(word string) []string {
	res := []string{}
	for i := 0; i < len(word); i++ {
		res = append(res, "_")
	}
	return res
}
