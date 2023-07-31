package slice

func FindIndexInt(slice []int, target int) int {
	for index := range slice {
		if target == slice[index] {
			return index
		}
	}
	return -1
}
