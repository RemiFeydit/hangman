package utils 

func ContainsSliceInt(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}