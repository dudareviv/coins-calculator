package calculator

import (
	"errors"
	"fmt"
	"sort"
)

type Calculator struct {
	denominations []int
}

func NewCalculator(denominations []int) *Calculator {
	c := new(Calculator)

	c.denominations = denominations
	sort.Ints(c.denominations)

	return c
}

func (c *Calculator) Calculate(value int) (map[int]int, error) {
	result := make(map[int]int)

	reversedDenominations := reverse(c.denominations)

	for _, denomination := range reversedDenominations {
		var denominationCount = value / denomination
		value -= denominationCount * denomination

		result[denomination] = denominationCount
	}

	if value > 0 {
		message := fmt.Sprintf("Сумма не делится на цело. Остаток: %d", value)

		return result, errors.New(message)
	}

	return result, nil
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
