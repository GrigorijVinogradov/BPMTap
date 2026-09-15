package util

func IsKeyPressSpecificLetter(b []byte, letter string) bool {
	if string(b) == letter {
		return true
	}
	return false
}
