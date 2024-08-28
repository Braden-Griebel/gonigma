package machine

import (
	"testing"
)

func TestPlugBoard(t *testing.T) {
	testBoard := newPlugboard()
	testBoard.addWire('e', 'w')
	testBoard.addWire('f', 'q')
	if testBoard.translateChar('e') != 'w' {
		t.Error("Expected 'w', got ", testBoard.translateChar('e'))
	}
	if testBoard.translateChar('w') != 'e' {
		t.Error("Expected 'e', got ", testBoard.translateChar('w'))
	}
	if testBoard.translateChar('f') != 'q' {
		t.Error("Expected 'q', got ", testBoard.translateChar('f'))
	}
	if testBoard.translateChar('q') != 'f' {
		t.Error("Expected 'f', got ", testBoard.translateChar('q'))
	}
	testBoard = getPlugboard("a-e,b-r")
	if testBoard.translateChar('b') != 'r' {
		t.Errorf("Expected 'r', got %q", testBoard.translateChar('b'))
	}
	if testBoard.translateChar('r') != 'b' {
		t.Errorf("Expected 'e', got %q", testBoard.translateChar('b'))
	}
	if testBoard.translateChar('e') != 'a' {
		t.Errorf("Expected 'a', got %q", testBoard.translateChar('e'))
	}
	if testBoard.translateChar('a') != 'e' {
		t.Errorf("Expected 'e', got %q", testBoard.translateChar('a'))
	}
}
