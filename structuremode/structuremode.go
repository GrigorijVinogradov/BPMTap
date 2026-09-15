package structuremode

import (
	"bpmtap/printing"
	"bpmtap/timesignature"
	"bpmtap/util"
	"fmt"
	"os"
	"strings"
)

func StructureMode(SongParts []SongPart) []SongPart {	
	printing.ClearScreen()

	state := structureState {
		count: 0,
		bars: 0,
		isFirstRun: true,
		lastCharacter: "",
	}

	ts := timesignature.TimeSignature {
		BeatsPerMeasure: 4,
	}

	PrintInitialStructureDisplay()
	PrintStructureKeybinds()

	var b []byte = make([]byte, 1)
	for {
		for _, v := range SongParts {
			v.printSongPart()
		}

		os.Stdin.Read(b)
		printing.ClearScreen()


		if util.IsKeyPressSpecificLetter(b, " ") {
			SongParts = append(SongParts, SongPart{ state.lastCharacter, state.bars })
			state.resetAndPrint()

			continue
		}

		if util.IsKeyPressSpecificLetter(b, "d") {
			if len(SongParts) > 0 {
				SongParts = SongParts[:len(SongParts)-1]
			}

			state.resetAndPrint()
			continue
		}

		if util.IsKeyPressSpecificLetter(b, "r") {
			state.resetAndPrint()

			continue
		}

		if util.IsKeyPressSpecificLetter(b, "t") {
			ts.Toggle()

			state.resetAndPrint()

			continue
		}


		if string(state.lastCharacter) != string(b) {
			if state.isFirstRun {
				state.lastCharacter = string(b)
				state.isFirstRun = false
			} else {
				SongParts = append(SongParts, SongPart{ state.lastCharacter, state.bars })
				state.count = 0
				state.bars = 0
				state.lastCharacter = string(b)
			}
		}

		fmt.Print(string(b), ": ")
		fmt.Print(strings.Repeat("*", state.count+1), strings.Repeat(" ", ts.BeatsPerMeasure-state.count))

		if state.count == ts.BeatsPerMeasure-1 {
			state.count = 0
			state.bars += 1
		} else {
			state.count += 1
		}

		fmt.Print(state.bars, " Bars")
		fmt.Println()
		fmt.Println()

		PrintStructureKeybinds()

		if util.IsKeyPressSpecificLetter(b, "q") {
			break
		}
	}

	return SongParts
}
