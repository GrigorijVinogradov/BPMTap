package structuremode

import (
	"fmt"
)

type SongPart struct {
	character string
	bars int
}

func (sp SongPart) printSongPart() {
	fmt.Print(string(sp.character), ": ")
	fmt.Print(sp.bars, " Bars")
	fmt.Println()
}

