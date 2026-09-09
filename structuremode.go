package main

import (
	"os"
	"bpmtap/printing"
	"strings"
	"fmt"
)

var structureBeatsPerMeasure int = 4

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

	count := 0
	bars := 0

	isFirstRun := true

	PrintInitialStructureDisplay()
	PrintStructureKeybinds()

	var lastCharacter string

	var b []byte = make([]byte, 1)
	for {
		for _, v := range songParts {
			v.printSongPart()
		}

		os.Stdin.Read(b)
		printing.ClearScreen()


		if isKeyPressSpecificLetter(b, " ") {
			songParts = append(songParts, songPart{ lastCharacter, bars })
			count = 0
			bars = 0

			lastCharacter = ""
			isFirstRun = true

			fmt.Println()
			fmt.Println()

			PrintStructureKeybinds()

			continue
		}

		if isKeyPressSpecificLetter(b, "d") {
			if len(songParts) > 0 {
				songParts = songParts[:len(songParts)-1]
			}
			count = 0
			bars = 0
			isFirstRun = true

			lastCharacter = ""

			fmt.Println()
			fmt.Println()

			PrintStructureKeybinds()

			continue
		}

		if isKeyPressSpecificLetter(b, "r") {
			count = 0
			bars = 0

			lastCharacter = ""
			isFirstRun = true

			fmt.Println()
			fmt.Println()

			PrintStructureKeybinds()

			continue
		}

		if isKeyPressSpecificLetter(b, "t") {
			count = 0
			bars = 0
			
			if structureBeatsPerMeasure == 4 {
				structureBeatsPerMeasure = 6
			} else {
				structureBeatsPerMeasure = 4
			}

			lastCharacter = ""
			isFirstRun = true

			fmt.Println()
			fmt.Println()

			PrintStructureKeybinds()

			continue
		}


		if string(lastCharacter) != string(b) {
			if isFirstRun {
				lastCharacter = string(b)
				isFirstRun = false
			} else {
				songParts = append(songParts, songPart{ lastCharacter, bars })
				count = 0
				bars = 0
				lastCharacter = string(b)
			}
		}

		fmt.Print(string(b), ": ")
		fmt.Print(strings.Repeat("*", count+1), strings.Repeat(" ", structureBeatsPerMeasure-count))

		if count == structureBeatsPerMeasure-1 {
			count = 0
			bars += 1
		} else {
			count += 1
		}

		fmt.Print(bars, " Bars")
		fmt.Println()
		fmt.Println()

		PrintStructureKeybinds()

		if isKeyPressSpecificLetter(b, "q") {
			break
		}
	}

	return songParts
}
