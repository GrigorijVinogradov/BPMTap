package bpmmode

import (
	"bpmtap/printing"
	"bpmtap/timesignature"
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

func (st *bpmState) resetAndReprint() {
	printing.ClearScreen()
	st.reset()
	PrintInitialDisplay()
}


func (st *bpmState) updateBpmStats(tapIntervals []int64, ts timesignature.TimeSignature) {
	avg := average(tapIntervals[1:ts.BeatsPerMeasure])
	st.bpm = 60000 / avg

	st.allBpms = append(st.allBpms, st.bpm)
	st.avgBpms = average(st.allBpms)
}

