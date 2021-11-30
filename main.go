package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Необходимо указать сумму")
		return
	}

	if len(os.Args) < 3 {
		println("Необходимо указать номиналы и сумму")
		return
	}

	// Получить номиналы
	var argumentNominals = strings.Split(os.Args[1], ",")
	var nominals = make([]int, len(argumentNominals))

	for i, nominal := range argumentNominals {
		var convertedNominal, err = strconv.Atoi(nominal)

		if err != nil {
			println("Некорректный номинал", nominal, err)
			return
		}

		nominals[i] = convertedNominal
	}

	sort.Ints(nominals)

	// Получить сумму
	var money, err = strconv.Atoi(os.Args[2])

	if err != nil {
		println("Некорректная сумма", os.Args[2], err)
		return
	}

	// Проверить деление на минимальный номинал
	var remainder = money % nominals[0]

	if remainder > 0 {
		println("Сумма не делится на самый минимальный номинал", nominals[0])
		return
	}

	// Почитать и вывести результат
	remainder = money
	nominals = reverse(nominals)

	for _, nominal := range nominals {
		var nominalCount = remainder / nominal
		remainder -= nominalCount * nominal

		println("Номинал:", nominal, "Количество:", nominalCount)
	}
}

func reverse(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
