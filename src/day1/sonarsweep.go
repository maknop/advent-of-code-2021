/************************************
 *
 *	Day 1: Sonar Sweep
 *
 ************************************/

package problem1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func RunDay1() {
	part1Filename := "part1_sonar_sweep_depths.txt"
	part2Filename := "part2_sonar_sweep_depths.txt"

	part1Data := GetFileContents(part1Filename)
	part2Data := GetFileContents(part2Filename)

	// Part 1
	numIncreasesPart1 := CountSingleMeasurementIncreases(part1Data)

	// Part 2
	numIncreasesPart2 := CountSingleMeasurementIncreases(part2Data)

	fmt.Print("\nday 1, part 1: Sonar Sweep || The answer is: ", numIncreasesPart1)
	fmt.Print("\nday 1, part 2: Sonar Sweep || The answer is: ", numIncreasesPart2, "\n")
}

func GetFileContents(filename string) []string {
	var data string

	// Text file is converted to Golang code and is executed as binary when program is executed.
	box := packr.New("fileBox", "./")
	data, err := box.FindString(filename)
	if err != nil {
		print(err)
	}

	DataArray := strings.Split(data, "\n")
	return DataArray
}

/************************************
 *	Day 1, Part 1
 ************************************/
func CountSingleMeasurementIncreases(DataArray []string) int {
	var count int

	for i := 0; i < len(DataArray)-1; i += 1 {
		next, err := strconv.Atoi(DataArray[i+1])
		curr, err := strconv.Atoi(DataArray[i])
		if err != nil {
			print(err)
		}

		if next > curr {
			count += 1
		}
	}

	return count
}

/************************************
 *	Day 1, Part 2
 ************************************/
func CountSumsOfThreeIncreases(DataArray []string) int {
	var count int

	for i := 0; i < len(DataArray)-2; i += 1 {
		nextSum, err := strconv.Atoi(DataArray[i] + DataArray[i+1] + DataArray[i+2])
		currSum, err := strconv.Atoi(DataArray[i-1] + DataArray[i] + DataArray[i+1])
		if err != nil {
			print(err)
		}

		if nextSum > currSum {
			count += 1
		}
	}

	return count
}
