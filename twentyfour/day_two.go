package twentyfour

import (
	"strconv"
	"strings"
)

func CheckForIncreasing(nums string) bool {
	parts := strings.Fields(nums)
	if len(parts) == 0 {
		return false
	}

	prev, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	for _, s := range parts[1:] {
		current, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		diff := current - prev
		if diff <= 0 || diff > 3 {
			return false
		}

		prev = current
	}
	return true
}

func CheckForDecreasing(nums string) bool {
	parts := strings.Fields(nums)
	if len(parts) == 0 {
		return false
	}

	prev, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	for _, s := range parts[1:] {
		current, err := strconv.Atoi(s)
		if err != nil {
			return false
		}

		diff := prev - current
		if diff <= 0 || diff > 3 {
			return false
		}

		prev = current
	}
	return true
}

func CountSafeLevels(data string) int {
	var count int
	data = strings.TrimSpace(data)
	for line := range strings.SplitAfterSeq(data, "\n") {
		if line == "" {
			continue
		}
		if CheckForDecreasing(line) || CheckForIncreasing(line) {
			count++
			continue
		}

	}
	return count
}
