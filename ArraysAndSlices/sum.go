package main

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// 引数の型を ...[]int とすることで任意の数のスライスを受け取ることができる
// numbersToSum内の各スライスの先頭要素を除いたスライスの合計値を新しいスライスの要素として追加する
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
