package util

import (
	"os"

	"golang.org/x/term"
)

var b []byte = make([]byte, 1)

func SetTerminalToRawInputMode() func() {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	return func() {
		term.Restore(fd, oldState)
	}
}

func ReadByte() []byte {
	os.Stdin.Read(b)
	return b
}
