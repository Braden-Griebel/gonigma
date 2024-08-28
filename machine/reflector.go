package machine

import (
	"strings"
)

type reflector struct {
	configuration [26]uint
	offset        uint
}

func newReflector(configuration string) reflector {
	configuration = strings.ToLower(configuration)
	var newConfig [26]uint
	for idx, char := range configuration {
		newConfig[idx] = uint(char) - uint('a')
	}
	return reflector{
		configuration: newConfig,
		offset:        0,
	}
}

func getReflector(reflectorSpecification string) reflector {
	reflectorSpecification = strings.ToLower(reflectorSpecification)
	switch reflectorSpecification {
	case "a":
		return newA()
	case "b":
		return newB()
	case "c":
		return newC()
	default:
		return newA()
	}
}

func (r *reflector) translateChar(char rune) rune {
	charVal := uint(char) - uint('a')
	outVal := (charVal + r.offset) % 26
	return rune(wrapSub(r.configuration[outVal], r.offset, 26) + uint('a'))
}

func (r *reflector) set(setting rune) {
	r.offset = uint(setting) - uint('a')
}

// region specific wheels

func newA() reflector {
	return newReflector("EJMZALYXVBWFCRQUONTSPIKHGD")
}

func newB() reflector {
	return newReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
}

func newC() reflector {
	return newReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL")
}

// endregion specific wheels
