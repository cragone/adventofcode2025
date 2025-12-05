package twentyfive

import (
	"strconv"
	"strings"
)

func InvalidIdCount(data string) (sum int) {
	fields := strings.Split(data, ",")

	for _, field := range fields {
		parts := strings.Split(field, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for n := start; n <= end; n++ {
			if isRepeatedPattern(n) {
				sum += n
			}
		}
	}

	return
}

func isRepeatedPattern(n int) bool {
	s := strconv.Itoa(n)
	L := len(s)

	for size := 1; size*2 <= L; size++ {
		if L%size != 0 {
			continue
		}

		pattern := s[:size]
		repeats := L / size

		ok := true
		for i := 1; i < repeats; i++ {
			if s[i*size:(i+1)*size] != pattern {
				ok = false
				break
			}
		}

		if ok && repeats >= 2 {
			return true
		}
	}

	return false
}
