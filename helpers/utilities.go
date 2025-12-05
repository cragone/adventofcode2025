package helpers

import (
	"strconv"
)

func ConverStringsToNumbers(stringOne string, stringTwo string) (int, int) {

	one, _ := strconv.Atoi(stringOne)
	two, _ := strconv.Atoi(stringTwo)

	return one, two
}
