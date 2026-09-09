package main

import (
	"bpmtap/printing"
	"os"
	"os/exec"
)

func isKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}

func average(collection []int64) int64 {
	var total int64 = 0
	for _, v := range collection{
		total += v
	}

	return total / int64(len(collection))
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
	var savedSongParts []songPart

	var b []byte = make([]byte, 1)
	for {
		printing.PrintMainMenu()
		os.Stdin.Read(b)

		if isKeyPressSpecificLetter(b, "b") { 
			savedBpmAvgs = bpmMode(savedBpmAvgs)
		}

		if isKeyPressSpecificLetter(b, "s") { 
			savedSongParts = structureMode(savedSongParts)
		}

		if isKeyPressSpecificLetter(b, "q") { 
			printing.ClearScreen()
			break
		}
	}
}
