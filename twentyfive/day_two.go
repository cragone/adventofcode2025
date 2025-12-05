package twentyfive

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func InvalidIdCount(data string) (count int) {
	var targets = []int{2, 3, 5, 7}
	fields := strings.Split(data, ",")
	for _, field := range fields {
		numbers := strings.Split(field, "-")
		for _, target := range targets {
			num := ReturnGreatestSlice(numbers[0], target)
			start, _ := strconv.Atoi(numbers[0])
			final, _ := strconv.Atoi(numbers[1])
			count = FindRepeatingNumberSequences(target, num, start, final, count)
		}
		// for each of these need to split the number more and set the klowest to number
	}

	return
}

func ReturnGreatestSlice(number string, target int) int {
	var nums []int
	switch target {
	case 2:
		num1 := number[0:(len(number) / 2)]
		num2 := number[(len(number) / 2):]
		one, _ := strconv.Atoi(num1)
		two, _ := strconv.Atoi(num2)
		nums = append(nums, one)
		nums = append(nums, two)
	case 3:
		num1 := number[0:(len(number) / 3)]
		num2 := number[(len(number) / 3):((len(number) / 3) * 2)]
		num3 := number[((len(number) / 3) * 2):]
		one, _ := strconv.Atoi(num1)
		two, _ := strconv.Atoi(num2)
		three, _ := strconv.Atoi(num3)
		nums = append(nums, one)
		nums = append(nums, two)
		nums = append(nums, three)
	case 5:
		num1 := number[0:(len(number) / 5)]
		num2 := number[(len(number) / 5):((len(number) / 5) * 2)]
		num3 := number[((len(number) / 5) * 2):((len(number) / 5) * 3)]
		num4 := number[((len(number) / 5) * 3):((len(number) / 5) * 4)]
		num5 := number[((len(number) / 5) * 4):]
		one, _ := strconv.Atoi(num1)
		two, _ := strconv.Atoi(num2)
		three, _ := strconv.Atoi(num3)
		four, _ := strconv.Atoi(num4)
		five, _ := strconv.Atoi(num5)
		nums = append(nums, one)
		nums = append(nums, two)
		nums = append(nums, three)
		nums = append(nums, four)
		nums = append(nums, five)
	case 7:
		num1 := number[0:(len(number) / 7)]
		num2 := number[(len(number) / 7):((len(number) / 7) * 2)]
		num3 := number[((len(number) / 7) * 2):((len(number) / 7) * 3)]
		num4 := number[((len(number) / 7) * 3):((len(number) / 7) * 4)]
		num5 := number[((len(number) / 7) * 4):((len(number) / 7) * 5)]
		num6 := number[((len(number) / 7) * 5):((len(number) / 7) * 6)]
		num7 := number[((len(number) / 7) * 6):]
		one, _ := strconv.Atoi(num1)
		two, _ := strconv.Atoi(num2)
		three, _ := strconv.Atoi(num3)
		four, _ := strconv.Atoi(num4)
		five, _ := strconv.Atoi(num5)
		six, _ := strconv.Atoi(num6)
		seven, _ := strconv.Atoi(num7)
		nums = append(nums, one)
		nums = append(nums, two)
		nums = append(nums, three)
		nums = append(nums, four)
		nums = append(nums, five)
		nums = append(nums, six)
		nums = append(nums, seven)
	default:
		num1 := number[0:(len(number) / 2)]
		num2 := number[(len(number) / 2):]
		one, _ := strconv.Atoi(num1)
		two, _ := strconv.Atoi(num2)
		nums = append(nums, one)
		nums = append(nums, two)

	}

	greatest := nums[0]
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

// 43880368548

func FindRepeatingNumberSequences(target, number, start, final, count int) int {
	var currentValue []int
	for i := 0; i < target; i++ {
		currentValue = append(currentValue, number)
	}
	if ConvertNumbersToANumber(currentValue) < start {
		number++
		return FindRepeatingNumberSequences(target, number, start, final, count)
	}
	if ConvertNumbersToANumber(currentValue) >= final {
		return count
	}

	count += ConvertNumbersToANumber(currentValue)
	fmt.Println("current:", ConvertNumbersToANumber(currentValue))
	number++
	return FindRepeatingNumberSequences(target, number, start, final, count)
}

func ConvertNumbersToANumber(numbers []int) (result int) {
	result = numbers[0]
	for i := 1; i < len(numbers); i++ {
		digits := int(math.Log10(float64(numbers[i]))) + 1

		result = result*int(math.Pow10(digits)) + numbers[i]

	}
	return result
}

//now i have to split to two check
// split to three check
// split to five check
// split to seven check
