package main

import (
	"bpmtap/bpmmode"
	"bpmtap/printing"
	"bpmtap/structuremode"
	"os"

	"golang.org/x/term"
)

func isKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}

func setupTerminalInput() func() {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	return func() {
		term.Restore(fd, oldState)
	}
}

func main() {
	restorefunc := setupTerminalInput()
	defer restorefunc()
	var savedBpmAvgs []int64
	var savedSongParts []structuremode.SongPart

	var b []byte = make([]byte, 1)
	for {
		printing.PrintMainMenu()
		os.Stdin.Read(b)

		if isKeyPressSpecificLetter(b, "b") {
			savedBpmAvgs = bpmmode.BpmMode(savedBpmAvgs)
		}

		if isKeyPressSpecificLetter(b, "s") {
			savedSongParts = structuremode.StructureMode(savedSongParts)
		}

		if isKeyPressSpecificLetter(b, "q") {
			printing.ClearScreen()
			break
		}
	}
}
