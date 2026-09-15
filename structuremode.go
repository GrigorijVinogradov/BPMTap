package main

import (
	"os"
	"bpmtap/printing"
	"strings"
	"fmt"
)

var structureBeatsPerMeasure int = 4

type structureState struct {
	count, bars int
	isFirstRun bool
	lastCharacter string
}

func (st *structureState) reset() {
	st.count = 0
	st.bars = 0
	st.lastCharacter = ""
	st.isFirstRun = true
}

func (st *structureState) resetAndPrint() {
	st.reset()

	fmt.Println()
	fmt.Println()

	PrintStructureKeybinds()
}

type songPart struct {
	character string
	bars int
}

func (sp songPart) printSongPart() {
	fmt.Print(string(sp.character), ": ")
	fmt.Print(sp.bars, " Bars")
	fmt.Println()
}

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

func structureMode(songParts []songPart) []songPart {	
	printing.ClearScreen()

	state := structureState {
		count: 0,
		bars: 0,
		isFirstRun: true,
		lastCharacter: "",
	}

	PrintInitialStructureDisplay()
	PrintStructureKeybinds()

	var b []byte = make([]byte, 1)
	for {
		for _, v := range songParts {
			v.printSongPart()
		}

		os.Stdin.Read(b)
		printing.ClearScreen()


		if isKeyPressSpecificLetter(b, " ") {
			songParts = append(songParts, songPart{ state.lastCharacter, state.bars })
			state.resetAndPrint()

			continue
		}

		if isKeyPressSpecificLetter(b, "d") {
			if len(songParts) > 0 {
				songParts = songParts[:len(songParts)-1]
			}

			state.resetAndPrint()
			continue
		}

		if isKeyPressSpecificLetter(b, "r") {
			state.resetAndPrint()

			continue
		}

		if isKeyPressSpecificLetter(b, "t") {
			if structureBeatsPerMeasure == 4 {
				structureBeatsPerMeasure = 6
			} else {
				structureBeatsPerMeasure = 4
			}

			state.resetAndPrint()

			continue
		}


		if string(state.lastCharacter) != string(b) {
			if state.isFirstRun {
				state.lastCharacter = string(b)
				state.isFirstRun = false
			} else {
				songParts = append(songParts, songPart{ state.lastCharacter, state.bars })
				state.count = 0
				state.bars = 0
				state.lastCharacter = string(b)
			}
		}

		fmt.Print(string(b), ": ")
		fmt.Print(strings.Repeat("*", state.count+1), strings.Repeat(" ", structureBeatsPerMeasure-state.count))

		if state.count == structureBeatsPerMeasure-1 {
			state.count = 0
			state.bars += 1
		} else {
			state.count += 1
		}

		fmt.Print(state.bars, " Bars")
		fmt.Println()
		fmt.Println()

		PrintStructureKeybinds()

		if isKeyPressSpecificLetter(b, "q") {
			break
		}
	}

	return songParts
}
