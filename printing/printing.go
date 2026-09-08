package printing

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

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

	fmt.Println()
}
