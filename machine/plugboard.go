package machine

import (
	"strings"
)

type plugboard struct {
	wires [26]uint
}

func newPlugboard() plugboard {
	var newWires [26]uint
	for i := uint(0); i < 26; i++ {
		newWires[i] = i
	}
	return plugboard{newWires}
}

func (p *plugboard) addWire(start rune, end rune) {
	startVal := uint(start) - uint('a')
	endVal := uint(end) - uint('a')
	p.wires[startVal] = endVal
	p.wires[endVal] = startVal
}

func (p *plugboard) translateChar(char rune) rune {
	charVal := uint(char) - uint('a')
	outVal := rune(p.wires[charVal] + uint('a'))
	return outVal
}

func getPlugboard(setting string) plugboard {
	setting = strings.ToLower(setting)
	pboard := newPlugboard()
	if len(setting) == 0 {
		return pboard
	}
	for _, s := range strings.Split(setting, ",") {
		wire := strings.Split(s, "-")
		start := rune(wire[0][0])
		end := rune(wire[1][0])
		pboard.addWire(start, end)
	}
	return pboard
}
