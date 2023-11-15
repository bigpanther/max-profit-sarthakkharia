package main

import (
	"fmt"
	"testing"
)

func Test_MaxProfit(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 6},
		{[]int{6, 5, 4, 3, 2, 1}, 0},
		{[]int{5, 10, 15, 11, 19, 7, 3, 1}, 14},
		{[]int{300, 500, 1001, 20, 40, 50, 100, 1100}, 1080},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := maxProfit(test.input)
			if test.expectedOutput != actual {
				t.Fail()
			}
		})
	}
}
