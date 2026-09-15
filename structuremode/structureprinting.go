package structuremode

import (
	"fmt"
)

func PrintInitialStructureDisplay() {
	fmt.Println("Tap any key to start tapping segment length!")
	fmt.Println()
}

func PrintStructureKeybinds() {
	fmt.Println("press q to quit")
	fmt.Println("press r to reset current segment")
	fmt.Println("press d to delete last segment")
	fmt.Println("press t to toggle between 4 or 6 beats per measure")
	fmt.Println("press spacebar to finish segment without starting a new one")

	fmt.Println()
}
