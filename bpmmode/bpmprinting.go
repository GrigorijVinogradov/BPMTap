package bpmmode

import (
	"bpmtap/printing"
)

func PrintInitialDisplay() {
	printing.PrintRawLn("Press Any Key to start BPM Tapping!")
	printing.PrintRawLn()
	printing.PrintRawLn()
}

func PrintInfos(isFirstRun bool, bpm int64, avgBpms int64) {
	printing.PrintRawLn()
	if(!isFirstRun) {
		printing.PrintRawLn(bpm, "Average:", avgBpms)
	} else {
		printing.PrintRawLn()
	}
	printing.PrintRawLn()
}

func PrintSavedAvgs(savedAvgs []int64) {
	for i, v := range savedAvgs {
		percentage := 100
		if i > 0 {
			percentage = int(float64(v) / float64(savedAvgs[i-1])*100)
		}
		printing.PrintRawLn("Tempo", i+1, "-", v, "-", percentage, "%")
	}
}

func PrintKeyBinds() {
	printing.PrintRawLn("press q to quit")
	printing.PrintRawLn("press p to save current average")
	printing.PrintRawLn("press n to save current average and start next immediately")
	printing.PrintRawLn("press r to reset current average")
	printing.PrintRawLn("press t to toggle between 4 or 6 beats per measure")

	printing.PrintRawLn()
}

