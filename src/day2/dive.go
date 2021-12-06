/************************************
 *
 *	Day 2: Dive!
 *
 ************************************/

package dive

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func RunDay2() {
	filename := "submarine_movement_data.txt"

	data := ReadData(filename)
	fmt.Print(data)

	fmt.Print("\nday 2, part 1: Dive! || The answer is: ", 0)
	fmt.Print("\nday 2, part 2: Dive! || The answer is: ", 0, "\n")
}

func ReadData(filename string) []string {
	var data string

	// Text file is converted to Golang code and is executed as binary when program is executed.
	box := packr.New("fileBox", "./")
	data, err := box.FindString(filename)
	if err != nil {
		print(err)
	}

	movementData := strings.Split(data, "\n")
	return movementData
}

/************************************
 *	Day 2, Part 1
 ************************************/
func AdjustMovement(newPosition int) int {
	return 0
}
