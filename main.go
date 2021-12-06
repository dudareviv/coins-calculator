package calculator

import (
	"errors"
	"sort"
)

type Calculator struct {
	// Desc order
	denominations []int
}

func NewCalculator(denominations []int) *Calculator {
	c := new(Calculator)

	c.denominations = denominations
	sort.Ints(c.denominations)
	c.denominations = reverse(c.denominations)

	return c
}

func (c *Calculator) Calculate(value int) (map[int]int, error) {
	result := make(map[int]int)

	// Prepare
	for _, denomination := range c.denominations {
		result[denomination] = 0
	}

	// Calculate
	denominationsLength := len(c.denominations)
	strategy := 0
	i := -1
	cycles := 0

	for {
		if value == 0 {
			break
		}

		if cycles > 255 {
			message := "cycles max reached"
			return nil, errors.New(message)
		}

		cycles++

		i++

		if i >= denominationsLength {
			i = 0
			strategy = 1
		}

		denomination := c.denominations[i]

		switch strategy {
		case 0:
			result[denomination] += value / denomination
			value = value % denomination

			break
		case 1:

			if result[denomination] == 0 {
				break
			}

			result[denomination]--
			value += denomination

			strategy = 0

			break
		}
	}

	// Optimize
	denominations := reverse(c.denominations)

	for i, denomination := range denominations {
		if i == len(denominations)-1 {
			break
		}

		nextDenomination := denominations[i+1]

		mcd := getMinCommonDivision(denomination, nextDenomination)

		c := result[denomination] * denomination

		if c < mcd {
			continue
		}

		result[denomination] -= c / denomination
		result[nextDenomination] += c / nextDenomination
	}

	return result, nil
}

func getMinCommonDivision(a, b int) int {
	if b%a == 0 {
		return b
	}

	return a * b
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
