package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func average(collection []int64) int64 {
	var total int64 = 0
	for _, v := range collection{
		total += v
	}

	return total / int64(len(collection))
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func isKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}

func printInitialDisplay() {
	fmt.Println("Press Any Key to start BPM Tapping!")
	fmt.Println()
	fmt.Println()
}

func printInfos(isFirstRun bool, bpm int64, avgBpms int64) {
	fmt.Println()
	if(!isFirstRun) {
		fmt.Println(bpm, "Average:", avgBpms)
	} else {
		fmt.Println()
	}
	fmt.Println()
}

func printSavedAvgs(savedAvgs []int64) {
	for i, v := range savedAvgs {
		percentage := 100
		if i > 0 {
			percentage = int(float64(v) / float64(savedAvgs[i-1])*100)
		}
		fmt.Println("Tempo", i+1, "-", v, "-", percentage, "%")
	}
}

func printKeyBinds() {
	fmt.Println("press q to quit")
	fmt.Println("press p to save current average")
	fmt.Println("press r to reset current average")

	fmt.Println()
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	count := 0

	tapIntervals := [4]int64 {}
	previous := time.Now()

	isFirstRun := true

	var allBpms []int64 

	clearScreen()

	var bpm int64
	var avgBpms int64

	var savedAvgs []int64

	printInitialDisplay()
	printKeyBinds()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		clearScreen()

		if isKeyPressSpecificLetter(b, "q") {
			break
		}

		if(count == 0) {
			previous = time.Now()
		}

		fmt.Print(strings.Repeat("*", count+1))

		if(count == 3) {
			tapTime := time.Now()
			interval := tapTime.Sub(previous)
			previous = tapTime
			tapIntervals[count] = interval.Milliseconds()

			avg := average(tapIntervals[1:])
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

		printInfos(isFirstRun, bpm, avgBpms)


		if isKeyPressSpecificLetter(b, "r") {
			clearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			printInitialDisplay()
		}

		if isKeyPressSpecificLetter(b, "p") {
			clearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			savedAvgs = append(savedAvgs, avgBpms)
			printInitialDisplay()
		}

		printKeyBinds()
		printSavedAvgs(savedAvgs)
	}
}
