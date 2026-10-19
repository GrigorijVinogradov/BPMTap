package bpmmode

import (
	"bpmtap/components"
	"bpmtap/printing"
	"bpmtap/util"
)

type bpmState struct {
	count      int
	isFirstRun bool
	allBpms    []int64
	avgBpms    int64
	bpm        int64
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

func (st *bpmState) updateBpmStats(tapIntervals []int64, ts components.TimeSignature) {
	avg := util.Average(tapIntervals[1:ts.BeatsPerMeasure])
	st.bpm = 60000 / avg

	st.allBpms = append(st.allBpms, st.bpm)
	st.avgBpms = util.Average(st.allBpms)
}
