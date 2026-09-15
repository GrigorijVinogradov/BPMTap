package structuremode

import (
	"fmt"
)

type structureState struct {
	count, bars int
	isFirstRun bool
	lastCharacter string
}

func (st *structureState) reset() {
	st.count = 0
	st.bars = 0
	st.lastCharacter = ""
	st.isFirstRun = true
}

func (st *structureState) resetAndPrint() {
	st.reset()

	fmt.Println()
	fmt.Println()

	PrintStructureKeybinds()
}

