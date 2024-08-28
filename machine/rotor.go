package machine

import (
	"strings"
)

type rotor struct {
	offset       uint
	charsForward [26]uint
	charsReverse [26]uint
	notches      map[uint]bool
}

func newRotor(chars string, notches string) rotor {
	var rotorForwardPath [26]uint
	for idx, char := range strings.ToLower(chars) {
		rotorForwardPath[idx] = uint(char) - uint('a')
	}
	var rotorReversePath [26]uint
	// Find reverse paths through the rotor
	for i := 0; i < 26; i++ {
		// Find which char is mapped in the forward direction
		fwd := rotorForwardPath[i]
		// Then map the out wire (fwd) back to the index that did this
		// mapping
		rotorReversePath[fwd] = uint(i)
	}
	rotorNotches := make(map[uint]bool)
	for _, char := range strings.ToLower(notches) {
		// +1 because of the way the notches are normally described
		// i.e. it steps the next rotor when it steps to this rotor
		rotorNotches[(uint(char)-uint('a')+1)%26] = true
	}
	return rotor{
		charsForward: rotorForwardPath,
		charsReverse: rotorReversePath,
		notches:      rotorNotches,
		offset:       0,
	}
}

func (r *rotor) set(setting rune) {
	r.offset = uint(setting) - uint('a')
}

func (r *rotor) step() bool {
	r.offset = (r.offset + 1) % 26
	if r.notches[r.offset] {
		return true
	}
	return false
}

func (r *rotor) translateForward(input rune) rune {
	inputVal := uint(input) - uint('a')
	outputVal := (inputVal + r.offset) % 26
	return rune(wrapSub(r.charsForward[outputVal], r.offset, 26) + uint('a'))
}

func (r *rotor) translateReverse(input rune) rune {
	inputVal := uint(input) - uint('a')
	outputVal := (inputVal + r.offset) % 26
	return rune(wrapSub(r.charsReverse[outputVal], r.offset, 26) + uint('a'))
}

func wrapSub(lhs uint, rhs uint, wrap uint) uint {
	if rhs <= lhs {
		return lhs - rhs
	} else {
		return wrap - (rhs - lhs)
	}
}

func getRotors(rotorSpecification string) [3]rotor {
	var rotors [3]rotor
	for idx, r := range strings.Split(rotorSpecification, ",") {
		switch r {
		case "I":
			rotors[idx] = newI()
		case "II":
			rotors[idx] = newII()
		case "III":
			rotors[idx] = newIII()
		case "IV":
			rotors[idx] = newIV()
		case "V":
			rotors[idx] = newV()
		case "VI":
			rotors[idx] = newVI()
		case "VII":
			rotors[idx] = newVII()
		case "VIII":
			rotors[idx] = newVIII()
		default:
			rotors[idx] = newI()
		}
	}
	return rotors
}

// region specific wheels

func newI() rotor {
	return newRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q")
}

func newII() rotor {
	return newRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "E")
}

func newIII() rotor {
	return newRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "V")
}

func newIV() rotor {
	return newRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J")
}

func newV() rotor {
	return newRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "Z")
}

func newVI() rotor {
	return newRotor("JPGVOUMFYQBENHZRDKASXLICTW", "ZM")
}

func newVII() rotor {
	return newRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "ZM")
}

func newVIII() rotor {
	return newRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "ZM")
}

// endregion specific wheels
