package main

import (
	"bpmtap/bpmmode"
	"bpmtap/printing"
	"bpmtap/structuremode"
	"bpmtap/util"
)

func main() {
	restorefunc := util.SetTerminalToRawInputMode()
	defer restorefunc()
	var savedBpmAvgs []int64
	var savedSongParts []structuremode.SongPart

	for {
		printing.PrintMainMenu()
		b := util.ReadByte()

		if util.IsKeyPressSpecificLetter(b, "b") {
			savedBpmAvgs = bpmmode.BpmMode(savedBpmAvgs)
			continue
		}

		if util.IsKeyPressSpecificLetter(b, "s") {
			savedSongParts = structuremode.StructureMode(savedSongParts)
			continue
		}

		if util.IsKeyPressSpecificLetter(b, "q") {
			printing.ClearScreen()
			break
		}
	}
}
