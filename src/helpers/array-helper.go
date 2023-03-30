package helpers

func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func IndexOf(array []int, number int) int {
	for key, value := range array {
		if value == number {
			return key
		}
	}
	return 0
}

func Contains(array []int, number int) bool {
	for _, value := range array {
		if value == number {
			return true
		}
	}

	return false
}
