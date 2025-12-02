package twentyfive

import (
	"fmt"
	"log"
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
	for command := range strings.FieldsSeq(data) {
		fmt.Printf("current is: %d, command is: %s ", current, command)
		switch {
		case strings.Contains(command, "L"):
			current = MoveDialLeft(current, command)
		default:
			current = MoveDialRight(current, command)
		}

		if current > 99 || current < 0 {
			fmt.Printf("command that caused problem:%s, current: %d\n", command, current)
			break
		}
		fmt.Printf("cur: %d\n", current)
		if current == 0 {
			password++
			fmt.Println("added to password")
		}
	}

	return password
}

// double check math
func MoveDialLeft(current int, command string) int {
	command = strings.ReplaceAll(command, "L", "")
	num, err := strconv.Atoi(command)
	if err != nil {
		log.Fatal(err)
	}
	return (((current - num) % 100) + 100) % 100
}

func MoveDialRight(current int, command string) int {
	command = strings.ReplaceAll(command, "R", "")
	num, err := strconv.Atoi(command)
	if err != nil {
		log.Fatal(err)
	}
	return (current + num) % 100
}
