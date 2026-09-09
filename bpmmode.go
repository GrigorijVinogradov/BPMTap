package main

import (
	"os"
	"bpmtap/printing"
	"strings"
	"fmt"
	"time"
)

var bpmBeatsPerMeasure int = 4

func PrintInitialDisplay() {
	fmt.Println("Press Any Key to start BPM Tapping!")
	fmt.Println()
	fmt.Println()
}

func PrintInfos(isFirstRun bool, bpm int64, avgBpms int64) {
	fmt.Println()
	if(!isFirstRun) {
		fmt.Println(bpm, "Average:", avgBpms)
	} else {
		fmt.Println()
	}
	fmt.Println()
}

func PrintSavedAvgs(savedAvgs []int64) {
	for i, v := range savedAvgs {
		percentage := 100
		if i > 0 {
			percentage = int(float64(v) / float64(savedAvgs[i-1])*100)
		}
		fmt.Println("Tempo", i+1, "-", v, "-", percentage, "%")
	}
}

func PrintKeyBinds() {
	fmt.Println("press q to quit")
	fmt.Println("press p to save current average")
	fmt.Println("press r to reset current average")
	fmt.Println("press t to toggle between 4 or 6 beats per measure")

	fmt.Println()
}

func bpmMode(savedAvgs []int64) []int64 {	
	count := 0

	tapIntervals := make([]int64, 9)
	previous := time.Now()

	isFirstRun := true

	var allBpms []int64 

	printing.ClearScreen()

	var bpm int64
	var avgBpms int64

	PrintInitialDisplay()
	PrintKeyBinds()
	PrintSavedAvgs(savedAvgs)

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		printing.ClearScreen()

		if isKeyPressSpecificLetter(b, "q") {
			break
		}

		if(count == 0) {
			previous = time.Now()
		}

		fmt.Print(strings.Repeat("*", count+1))

		if(count == bpmBeatsPerMeasure-1) {
			tapTime := time.Now()
			interval := tapTime.Sub(previous)
			previous = tapTime
			tapIntervals[count] = interval.Milliseconds()

			avg := average(tapIntervals[1:bpmBeatsPerMeasure])
			bpm = 60000 / avg

			allBpms = append(allBpms, bpm)
			avgBpms = average(allBpms)

			isFirstRun = false
			count = 0
		} else {
			tapTime := time.Now()
			interval := tapTime.Sub(previous)
			previous = tapTime
			tapIntervals[count] = interval.Milliseconds()

			count += 1
		}

		PrintInfos(isFirstRun, bpm, avgBpms)

		if isKeyPressSpecificLetter(b, "r") {
			printing.ClearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			PrintInitialDisplay()
		}

		if isKeyPressSpecificLetter(b, "t") {
			printing.ClearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			if bpmBeatsPerMeasure == 4 {
				bpmBeatsPerMeasure = 6
			} else {
				bpmBeatsPerMeasure = 4
			}
			PrintInitialDisplay()
		}

		if isKeyPressSpecificLetter(b, "p") {
			printing.ClearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			savedAvgs = append(savedAvgs, avgBpms)
			PrintInitialDisplay()
		}

		PrintKeyBinds()
 	 	PrintSavedAvgs(savedAvgs)
	}
	
	return savedAvgs
}
