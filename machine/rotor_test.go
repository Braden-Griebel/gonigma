package machine

import (
	"testing"
)

func TestWrapSub(t *testing.T) {
	data := []struct {
		lhs    uint
		rhs    uint
		wrap   uint
		result uint
	}{
		{lhs: 5, rhs: 3, wrap: 26, result: 2},
		{lhs: 5, rhs: 7, wrap: 26, result: 24},
		{lhs: 2, rhs: 10, wrap: 30, result: 22},
		{lhs: 3, rhs: 3, wrap: 10, result: 0},
	}
	for _, instance := range data {
		got := wrapSub(instance.lhs, instance.rhs, instance.wrap)
		want := instance.result
		if got != want {
			t.Logf("Expect: %d, got %d\n", instance.result, got)
			t.Fail()
		}
	}
}

func TestTranslateForward(t *testing.T) {
	data := []struct {
		r        rotor
		in       rune
		expected rune
	}{
		{r: newI(), in: 'u', expected: 'a'},
		{r: newI(), in: 'a', expected: 'e'},
	}
	for _, instance := range data {
		got := instance.r.translateForward(instance.in)
		want := instance.expected
		if got != want {
			t.Logf("Expect: %d, got %d\n", want, got)
			t.Fail()
		}
	}
}

func TestTranslateReverse(t *testing.T) {
	data := []struct {
		r        rotor
		in       rune
		expected rune
	}{
		{r: newI(), in: 'e', expected: 'a'},
		{r: newI(), in: 'o', expected: 'm'},
	}
	for _, instance := range data {
		got := instance.r.translateReverse(instance.in)
		want := instance.expected
		if got != want {
			t.Logf("Expect: %d, got %d\n", want, got)
			t.Fail()
		}
	}
}

func TestStep(t *testing.T) {
	testRotor := newI()
	if testRotor.translateForward('a') != 'e' {
		t.Fail()
	}
	if testRotor.translateForward('b') != 'k' {
		t.Fail()
	}
	if testRotor.translateReverse('t') != 'l' {
		t.Fail()
	}

	_ = testRotor.step()
	if testRotor.offset != 1 {
		t.Fail()
	}
	if testRotor.translateForward('a') != 'j' {
		t.Fail()
	}
}

func TestNotch(t *testing.T) {
	testRotor := newI()
	testRotor.offset = 16
	if !testRotor.step() {
		t.Fail()
	}
}
