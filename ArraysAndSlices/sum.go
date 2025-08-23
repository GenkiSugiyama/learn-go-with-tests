package main

func Sum(numbers [5]int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
