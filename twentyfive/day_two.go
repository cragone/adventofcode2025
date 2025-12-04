package twentyfive

import (
	"fmt"
	"strconv"
	"strings"
)

func InvalidIdCount(data string) (count int) {
	fields := strings.Split(data, ",")
	for _, field := range fields {
		numbers := strings.Split(field, "-")
		num1 := numbers[0][0:(len(numbers[0]) / 2)]
		num2 := numbers[0][(len(numbers[0]) / 2):len(numbers[0])]
		start, _ := strconv.Atoi(numbers[0])
		final, _ := strconv.Atoi(numbers[1])
		one, two := ConverStringsToNumbers(num1, num2)
		if one > two {
			count = FindRepeatingNumberSequences(two, start, final, count)
		} else {
			count = FindRepeatingNumberSequences(one, start, final, count)
		}

	}

	return
}

func FindRepeatingNumberSequences(number, start, final, count int) int {
	if ConvertNumbersToANumber(number, number) < start {
		number++
		return FindRepeatingNumberSequences(number, start, final, count)
	}

	if ConvertNumbersToANumber(number, number) >= final {
		return count
	}

	count += ConvertNumbersToANumber(number, number)
	fmt.Println("current:", ConvertNumbersToANumber(number, number))
	number++
	return FindRepeatingNumberSequences(number, start, final, count)
}

func ConverStringsToNumbers(stringOne string, stringTwo string) (int, int) {

	one, _ := strconv.Atoi(stringOne)
	two, _ := strconv.Atoi(stringTwo)

	return one, two
}

func ConvertNumbersToANumber(one int, two int) int {
	num, _ := strconv.Atoi(fmt.Sprintf("%d%d", one, two))
	return num
}

// 49 - 86 has three repeating for both 1-36 and the ladder
