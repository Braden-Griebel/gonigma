package machine

import "strings"

type Enigma struct {
	rotors    [3]rotor
	reflector reflector
	plugboard plugboard
}

func NewMachine(rotorSpecification string,
	rotorSettings string,
	reflectorSpecification string,
	reflectorSetting string,
	plugboardSpecification string,
) Enigma {
	newRotors := getRotors(rotorSpecification)
	newRefl := getReflector(reflectorSpecification)
	for idx, setting := range strings.Split(rotorSettings, ",") {
		newRotors[idx].set(rune(setting[0]))
	}
	newRefl.set(rune(reflectorSetting[0]))
	newPlug := getPlugboard(plugboardSpecification)
	return Enigma{
		rotors:    newRotors,
		reflector: newRefl,
		plugboard: newPlug,
	}
}

func (e *Enigma) step() {
	stepNext := true
	for i := 2; i >= 0; i-- {
		if stepNext {
			stepNext = e.rotors[i].step()
		}
	}
}

func (e *Enigma) translateChar(char rune) rune {
	// Through the plugboard
	char = e.plugboard.translateChar(char)
	// Through the rotors
	for _, r := range e.rotors {
		char = r.translateForward(char)
	}
	// Through the reflector
	char = e.reflector.translateChar(char)
	// Back through the rotor
	for i := 2; i >= 0; i-- {
		char = e.rotors[i].translateReverse(char)
	}
	// Back through the plugboard
	char = e.plugboard.translateChar(char)
	// Step the machine
	e.step()
	return char
}

func (e *Enigma) Translate(input string) string {
	output := make([]rune, len(input))
	for idx, char := range input {
		switch {
		case 'a' <= char && char <= 'z':
			output[idx] = e.translateChar(char)
		case 'A' <= char && char <= 'Z':
			// Uppercase should go back to uppercase
			output[idx] = rune(uint(e.translateChar(rune(uint(char)-uint('A')+uint('a')))) - 'a' + uint('A'))
		default:
			output[idx] = char
		}
	}
	return string(output)
}
