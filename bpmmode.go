package main

import (
	"os"
	"bpmtap/printing"
	"strings"
	"fmt"
	"time"
)

type bpmState struct {
	count int
	isFirstRun bool
	allBpms []int64
	avgBpms int64
	bpm int64
}

func (st *bpmState) reset() {
	st.count = 0
	st.allBpms = []int64{} 
	st.isFirstRun = true
}

func (st *bpmState) updateBpmStats(tapIntervals []int64) {
	avg := average(tapIntervals[1:bpmBeatsPerMeasure])
	st.bpm = 60000 / avg

	st.allBpms = append(st.allBpms, st.bpm)
	st.avgBpms = average(st.allBpms)
}

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

func reset(st *bpmState) {
	printing.ClearScreen()
	st.reset()
	PrintInitialDisplay()
}

func getMillisecondInterval(previous time.Time) (ms int64, newPrevious time.Time) {
	tapTime := time.Now()
	interval := tapTime.Sub(previous)
	return interval.Milliseconds(), tapTime
}

func bpmMode(savedAvgs []int64) []int64 {	
	state := bpmState {
		count: 0,
		isFirstRun: true,
		allBpms: []int64{},
		avgBpms: 0,
		bpm: 0,
	}

	tapIntervals := make([]int64, 9)
	previous := time.Now()

	printing.ClearScreen()

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

		if(state.count == 0) {
			previous = time.Now()
		}

		fmt.Print(strings.Repeat("*", state.count+1))

		if(state.count == bpmBeatsPerMeasure-1) {
			tapIntervals[state.count], previous = getMillisecondInterval(previous)

			state.updateBpmStats(tapIntervals)
			
			state.isFirstRun = false
			state.count = 0
		} else {
			tapIntervals[state.count], previous = getMillisecondInterval(previous)

			state.count += 1
		}

		PrintInfos(state.isFirstRun, state.bpm, state.avgBpms)
		
		handleSpecificKeys(b, &state, &savedAvgs)

		PrintKeyBinds()
 	 	PrintSavedAvgs(savedAvgs)
	}
	
	return savedAvgs
}

func handleSpecificKeys(b []byte, state *bpmState, savedAvgs *[]int64) {
		if isKeyPressSpecificLetter(b, "r") {
			reset(state)
		}

		if isKeyPressSpecificLetter(b, "t") {
			reset(state)
			if bpmBeatsPerMeasure == 4 {
				bpmBeatsPerMeasure = 6
			} else {
				bpmBeatsPerMeasure = 4
			}
		}

		if isKeyPressSpecificLetter(b, "p") {
			reset(state)
			*savedAvgs = append(*savedAvgs, state.avgBpms)
		}

}
