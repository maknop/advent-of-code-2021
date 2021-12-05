package problem1

import (
	"bufio"
	"fmt"
	"os"
)

func RunProblem1() {
	filename := "sonar_sweep_depths.txt"
	data := GetFileContents(filename)

	print(data)
}

func GetFileContents(filename string) []string {
	var data []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: File does not exist.")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
