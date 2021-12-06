/************************************
 *
 *	Day 1: Sonar Sweep
 *
 ************************************/

package problem1

import (
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func RunDay1() {
	filename := "sonar_sweep_depths.txt"
	data := GetFileContents(filename)
	numIncreases := CountSingleMeasurementIncreases(data)

	print("\nday 1, part 1: Sonar Sweep || The answer is: ", numIncreases, "\n")
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
