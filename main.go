package main

import (
	"bpmtap/printing"
	"os"
	"os/exec"
	"bpmtap/bpmmode"
	"bpmtap/structuremode"
)


func isKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}

func setupTerminalInput() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func main() {
	setupTerminalInput()
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
