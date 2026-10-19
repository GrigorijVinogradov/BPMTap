package structuremode

import (
	"bpmtap/components"
	"bpmtap/printing"
	"bpmtap/util"
	"fmt"
	"strings"
)

func StructureMode(SongParts []SongPart) []SongPart {
	printing.ClearScreen()

	state := structureState{
		count:         0,
		bars:          0,
		isFirstRun:    true,
		lastCharacter: "",
	}

	ts := components.TimeSignature{
		BeatsPerMeasure: 4,
	}

	PrintInitialStructureDisplay()
	PrintStructureKeybinds()

	for {
		printSongParts(SongParts)

		b := util.ReadByte()
		printing.ClearScreen()

		//quitting
		if util.IsKeyPressSpecificLetter(b, "q") {
			break
		}

		//all other specific keybinds
		if handleKeyPress(b, &SongParts, &state, &ts) {
			continue
		}

		//switch songpart if key is switched
		if string(state.lastCharacter) != string(b) {
			if state.isFirstRun {
				state.lastCharacter = string(b)
				state.isFirstRun = false
			} else {
				SongParts = append(SongParts, SongPart{state.lastCharacter, state.bars})
				state.count = 0
				state.bars = 0
				state.lastCharacter = string(b)
			}
		}

		//print current counting bar
		fmt.Print(string(b), ": ")
		fmt.Print(strings.Repeat("*", state.count+1), strings.Repeat(" ", ts.BeatsPerMeasure-state.count))

		//increment bar count if beats per measure is reached
		if state.count == ts.BeatsPerMeasure-1 {
			state.count = 0
			state.bars += 1
		} else {
			state.count += 1
		}

		//finish printing the bar
		fmt.Print(state.bars, " Bars")

		printing.PrintRawLn()
		printing.PrintRawLn()

		PrintStructureKeybinds()
	}

	return SongParts
}

func handleKeyPress(b []byte, SongParts *[]SongPart, state *structureState, ts *components.TimeSignature) bool {
	if util.IsKeyPressSpecificLetter(b, " ") {
		*SongParts = append(*SongParts, SongPart{state.lastCharacter, state.bars})
		state.resetAndPrint()

		return true
	}

	if util.IsKeyPressSpecificLetter(b, "d") {
		if len(*SongParts) > 0 {
			*SongParts = (*SongParts)[:len(*SongParts)-1]
		}

		state.resetAndPrint()
		return true
	}

	if util.IsKeyPressSpecificLetter(b, "r") {
		state.resetAndPrint()

		return true
	}

	if util.IsKeyPressSpecificLetter(b, "t") {
		ts.Toggle()

		state.resetAndPrint()

		return true
	}

	return false
}
