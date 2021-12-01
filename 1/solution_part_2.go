package main

import (  
    "fmt"
    "io/ioutil"
	"strings"
	"strconv"
)

func main() {  
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	inputData := string(data)
	seaDepthArray := strings.Split(inputData, "\n")
	numMeasurements := len(seaDepthArray)
	increasingSteps := 0

	// Sum of 1,2 & 3 
	nextTripple := 0
	a, _ := strconv.Atoi(seaDepthArray[0])
	b, _ := strconv.Atoi(seaDepthArray[1])
	c, _ := strconv.Atoi(seaDepthArray[2])
	currentTripple := a + b + c 

	for idx, _ := range seaDepthArray {
		if idx + 3 == numMeasurements {
			break
		}
		if idx > 0 {
			currentTripple = nextTripple
		}
		a, _ := strconv.Atoi(seaDepthArray[idx + 1])
		b, _ := strconv.Atoi(seaDepthArray[idx + 2])
		c, _ := strconv.Atoi(seaDepthArray[idx + 3])
		nextTripple = a + b + c 

		
    	if currentTripple < nextTripple {
			increasingSteps += 1
		}
	}

	fmt.Println("Num Increasing Steps:", increasingSteps)
}