package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("expected %d but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// reflect.DeepEqualは型安全ではない
	// = 比較対象の2つの値が同じ型ではなくてもコンパイルが通ってしまう
	// そのため、型が異なる値を比較してしまうと、意図しない結果を招く可能性がある
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}
