package bpmmode

import (
	"os"
	"bpmtap/printing"
	"bpmtap/util"
	"bpmtap/timesignature"
	"strings"
	"fmt"
	"time"
)

func average(collection []int64) int64 {
	var total int64 = 0
	for _, v := range collection{
		total += v
	}

	return total / int64(len(collection))
}

func getMillisecondInterval(previous time.Time) (ms int64, newPrevious time.Time) {
	tapTime := time.Now()
	interval := tapTime.Sub(previous)
	return interval.Milliseconds(), tapTime
}

func BpmMode(savedAvgs []int64) []int64 {	
	state := bpmState {
		count: 0,
		isFirstRun: true,
		allBpms: []int64{},
		avgBpms: 0,
		bpm: 0,
	}

	ts := timesignature.TimeSignature {
		BeatsPerMeasure: 4,
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

		if util.IsKeyPressSpecificLetter(b, "q") {
			break
		}

		if(state.count == 0) {
			previous = time.Now()
		}

		fmt.Print(strings.Repeat("*", state.count+1))

		if(state.count == ts.BeatsPerMeasure-1) {
			tapIntervals[state.count], previous = getMillisecondInterval(previous)

			state.updateBpmStats(tapIntervals, ts)
			
			state.isFirstRun = false
			state.count = 0
		} else {
			tapIntervals[state.count], previous = getMillisecondInterval(previous)

			state.count += 1
		}

		PrintInfos(state.isFirstRun, state.bpm, state.avgBpms)
		
		handleSpecificKeys(b, &state, &savedAvgs, &ts)

		PrintKeyBinds()
 	 	PrintSavedAvgs(savedAvgs)
	}
	
	return savedAvgs
}

func handleSpecificKeys(b []byte, state *bpmState, savedAvgs *[]int64, ts *timesignature.TimeSignature) {
		if util.IsKeyPressSpecificLetter(b, "r") {
			state.resetAndReprint()
		}

		if util.IsKeyPressSpecificLetter(b, "t") {
			state.resetAndReprint()
			ts.Toggle()
		}

		if util.IsKeyPressSpecificLetter(b, "p") {
			state.resetAndReprint()
			*savedAvgs = append(*savedAvgs, state.avgBpms)
		}

}
