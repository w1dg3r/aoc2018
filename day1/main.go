package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getData(fileName string) []int {
	//Get data from file and append it to data array
	f, _ := os.Open(fileName)
	defer f.Close()

	fScanner := bufio.NewScanner(f)

	data := []int{}

	for fScanner.Scan() {
		lineValue, _ := strconv.Atoi(fScanner.Text())
		data = append(data, lineValue)
	}
	return data
}

func main() {
	//read data
	data := getData("input.txt")

	//Part 1 --------------------
	freq := 0
	for _, row := range data {
		freq += row
	}

	//Part 2 --------------------
	currentFreq := 0
	//Map unique datapoints
	dataFound := map[int]bool{
		0: true,
	}

	maxLoops := 1000 //Max number of loops to find freq
	fFound := false
	var i, jumps int = 0, 0

	for i = 0; i < maxLoops; i++ {
		for _, row := range data {
			jumps++
			currentFreq = currentFreq + row
			if dataFound[currentFreq] {
				fFound = true
				break
			} else {
				dataFound[currentFreq] = true
			}
		}

		if fFound {
			break
		}
	}

	//Print result
	fmt.Printf("Part 1: Value is %d\n", freq)
	fmt.Printf("Part 2: First freq was %d after %d loops and a total of %d jumps.\n", currentFreq, i, jumps)
}
