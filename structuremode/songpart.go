package structuremode

import (
	"bpmtap/printing"
	"fmt"
)

type SongPart struct {
	character string
	bars      int
}

func printSongParts(SongParts []SongPart) {
	for _, v := range SongParts {
		v.printSongPart()
	}
}

func (sp SongPart) printSongPart() {
	fmt.Print(string(sp.character), ": ")
	fmt.Print(sp.bars, " Bars")
	printing.PrintRawLn()
}
