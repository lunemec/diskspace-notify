package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	bools := []bool{true, true, true}
	if !all(bools) {
		t.Error("func all returned false, want true")
	}

	bools = []bool{true, false, true}
	if all(bools) {
		t.Error("func all returned true, want false")
	}
}

func TestCheckIntField(t *testing.T) {
	want := false
	returned := checkIntField(0, 0, "", false)
	if want != returned {
		t.Errorf("func checkIntField returned %v, want %v", returned, want)
	}

	want = true
	returned = checkIntField(1, 0, "", false)
	if want != returned {
		t.Errorf("func checkIntField returned %v, want %v", returned, want)
	}
}

func TestCheckUint8Field(t *testing.T) {
	want := false
	returned := checkUint8Field(uint8(0), uint8(0), "", false)
	if want != returned {
		t.Errorf("func checkUint8Field returned %v, want %v", returned, want)
	}

	want = true
	returned = checkUint8Field(uint8(1), uint8(0), "", false)
	if want != returned {
		t.Errorf("func checkUint8Field returned %v, want %v", returned, want)
	}
}

func TestCheckStringField(t *testing.T) {
	want := false
	returned := checkStringField("", "", "", false)
	if want != returned {
		t.Errorf("func checkStringField returned %v, want %v", returned, want)
	}

	want = true
	returned = checkStringField("something", "", "", false)
	if want != returned {
		t.Errorf("func checkStringField returned %v, want %v", returned, want)
	}
}

func TestCheckStringArrayField(t *testing.T) {
	want := false
	returned := checkStringArrayField([]string{}, 0, "", false)
	if want != returned {
		t.Errorf("func checkStringArrayField returned %v, want %v", returned, want)
	}

	want = true
	returned = checkStringArrayField([]string{"something"}, 0, "", false)
	if want != returned {
		t.Errorf("func checkStringArrayField returned %v, want %v", returned, want)
	}
}
