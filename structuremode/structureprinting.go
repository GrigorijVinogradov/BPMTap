package structuremode

import (
	"bpmtap/printing"
)

func Reprint() {
	printing.ClearScreen()
	PrintInitialStructureDisplay()
}

func PrintInitialStructureDisplay() {
	printing.PrintRawLn("Tap any key to start tapping segment length!")
	printing.PrintRawLn()
}

func PrintStructureKeybinds() {
	printing.PrintRawLn("press q to quit")
	printing.PrintRawLn("press r to reset current segment")
	printing.PrintRawLn("press d to delete last segment")
	printing.PrintRawLn("press t to toggle between 4 or 6 beats per measure")
	printing.PrintRawLn("press spacebar to finish segment without starting a new one")

	printing.PrintRawLn()
}
