package printing

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintMainMenu() {
	ClearScreen()
	fmt.Println("Press b for BPM Tap mode")
	fmt.Println("Press s for Structure Tap mode")
	fmt.Println("Press q to Quit")
}
