package main

import (  
    "fmt"
    "io/ioutil"
	"strings"
)

func main() {  
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	inputData := string(data)
	seaDepthArray := strings.Split(inputData, "\n")

	increasingSteps := 1
	for idx, currentSeaDepth := range seaDepthArray {
		if (idx == 0) {
			continue
		}
		prevSeaDepth := seaDepthArray[idx - 1]
		fmt.Println(idx)
		
    	if currentSeaDepth > prevSeaDepth {
			increasingSteps += 1
		}
	}

	fmt.Println("Num Increasing Steps:", increasingSteps)
}