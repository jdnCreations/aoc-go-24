package utils

func GetDistance(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -diff
	}
	return diff
}