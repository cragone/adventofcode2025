package twentyfive

import (
	"fmt"
	"strconv"
	"strings"
)

//L or left towards lower numbers
// R or right towards higher numbers

// a distance for how many dials rotated in that direction

//0-99 on the dial so if 99 +1 it goes to 0

//The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

//dial starts at 50
//
//

// need to update to count anytime the dial clicks 0
func GetPassword(current int, data string) (password int) {
	for _, command := range strings.Fields(data) {
		var count int
		fmt.Printf("current is: %d, command is: %s ", current, command)

		if strings.Contains(command, "L") {
			current, count = MoveDialLeft(current, command)
		} else {
			current, count = MoveDialRight(current, command)
		}

		password += count

		if current > 99 || current < 0 {
			fmt.Printf("command that caused problem:%s, current: %d\n", command, current)
			break
		}

		fmt.Printf("cur: %d, password: %d, added: %d\n", current, password, count)
	}
	return password
}

// double check math
func MoveDialLeft(current int, command string) (int, int) {
	command = strings.TrimPrefix(command, "L")
	move, _ := strconv.Atoi(command)

	full := move / 100
	partial := move % 100

	end := (current - partial + 100) % 100

	touches := full

	if partial > 0 && current > 0 && partial >= current {
		touches++
	}

	return end, touches
}

func MoveDialRight(current int, command string) (int, int) {
	command = strings.TrimPrefix(command, "R")
	move, _ := strconv.Atoi(command)

	full := move / 100
	partial := move % 100

	end := (current + partial) % 100

	touches := full

	if partial > 0 && current+partial >= 100 {
		touches++
	}

	return end, touches
}
