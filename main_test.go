package calculator

import "testing"

type testpair struct {
	denominations []int
	value         int
	result        map[int]int
	expectedError error
}

var tests = []testpair{
	{[]int{1, 2, 5, 10}, 88, map[int]int{10: 8, 5: 1, 2: 1, 1: 1}, nil},
	{[]int{1, 2, 5, 10}, 13, map[int]int{10: 1, 5: 0, 2: 1, 1: 1}, nil},
	{[]int{2, 5, 10}, 88, map[int]int{10: 8, 5: 0, 2: 4}, nil},
}

func TestCalculator_Calculate(t *testing.T) {
	for _, pair := range tests {
		c := NewCalculator(pair.denominations)

		result, err := c.Calculate(pair.value)

		if pair.expectedError != err {
			t.Error(
				"For", pair.denominations, pair.value,
				"expected error", pair.expectedError,
				"got", err,
			)
		}

		if len(pair.result) != len(result) {
			t.Error(
				"For", pair.denominations, pair.value,
				"expected", pair.value,
				"got", result,
			)
		}

		for denomination, count := range pair.result {
			if result[denomination] != count {
				t.Error(
					"For", pair.denominations, pair.value,
					"expected", pair.value,
					"got", result,
				)
			}
		}
	}
}
