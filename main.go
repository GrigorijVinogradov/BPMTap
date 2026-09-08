package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"bpmtap/printing"
)

func isKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}

func average(collection []int64) int64 {
	var total int64 = 0
	for _, v := range collection{
		total += v
	}

	return total / int64(len(collection))
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

	printing.ClearScreen()

	var bpm int64
	var avgBpms int64

	var savedAvgs []int64

	printing.PrintInitialDisplay()
	printing.PrintKeyBinds()

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

		printing.PrintInfos(isFirstRun, bpm, avgBpms)


		if isKeyPressSpecificLetter(b, "r") {
			printing.ClearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			printing.PrintInitialDisplay()
		}

		if isKeyPressSpecificLetter(b, "p") {
			printing.ClearScreen()
			count = 0
			allBpms = []int64{} 
			isFirstRun = true
			savedAvgs = append(savedAvgs, avgBpms)
			printing.PrintInitialDisplay()
		}

		printing.PrintKeyBinds()
		printing.PrintSavedAvgs(savedAvgs)
	}
}
