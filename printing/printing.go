package printing

import (
	"fmt"
	"strings"
)

func PrintRawLn(params ...any) {
	stringifiedParams := make([]string, len(params))
	for i, v := range params {
		stringifiedParams[i] = fmt.Sprint(v)
	}

	fmt.Print(strings.Join(stringifiedParams, " "), "\r\n")
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintMainMenu() {
	ClearScreen()
	PrintRawLn("Press b for BPM Tap mode")
	PrintRawLn("Press s for Structure Tap mode")
	PrintRawLn("Press q to Quit")
}
