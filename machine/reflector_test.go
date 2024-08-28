package machine

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	testReflector := newA()
	if testReflector.translateChar('e') != 'a' {
		t.Errorf("Expected 'a', got %q", testReflector.translateChar('e'))
	}
}

func testLoop(r *reflector, t *testing.T) {
	for i := uint(0); i < 26; i++ {
		char := rune(i + uint('a'))
		if r.translateChar(r.translateChar(char)) != char {
			t.Errorf("Error when testing reflectors")
		}
	}
}

func TestConfigurations(t *testing.T) {
	testReflector := newA()
	testLoop(&testReflector, t)
	testReflector = newB()
	testLoop(&testReflector, t)
	testReflector = newC()
	testLoop(&testReflector, t)
}
