package main

import (
	"reflect"
	"testing"
)

// ─────────────────────────────────────────────
// hexToDec Tests
// ─────────────────────────────────────────────

func TestHexToDec_LowercaseHex(t *testing.T) {
	input := "1f"
	expected := "31"
	result := hexToDec(input)
	if result != expected {
		t.Errorf("hexToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestHexToDec_UppercaseHex(t *testing.T) {
	input := "FF"
	expected := "255"
	result := hexToDec(input)
	if result != expected {
		t.Errorf("hexToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestHexToDec_ZeroValue(t *testing.T) {
	input := "0"
	expected := "0"
	result := hexToDec(input)
	if result != expected {
		t.Errorf("hexToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestHexToDec_InvalidInput(t *testing.T) {
	// "xyz" is not valid hex — should return unchanged
	input := "xyz"
	expected := "xyz"
	result := hexToDec(input)
	if result != expected {
		t.Errorf("hexToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestHexToDec_LargeHex(t *testing.T) {
	input := "7FFFFFFF"
	expected := "2147483647"
	result := hexToDec(input)
	if result != expected {
		t.Errorf("hexToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

// ─────────────────────────────────────────────
// binToDec Tests
// ─────────────────────────────────────────────

func TestBinToDec_BasicBinary(t *testing.T) {
	input := "1011"
	expected := "11"
	result := binToDec(input)
	if result != expected {
		t.Errorf("binToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestBinToDec_AllOnes(t *testing.T) {
	input := "11111111"
	expected := "255"
	result := binToDec(input)
	if result != expected {
		t.Errorf("binToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestBinToDec_Zero(t *testing.T) {
	input := "0"
	expected := "0"
	result := binToDec(input)
	if result != expected {
		t.Errorf("binToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestBinToDec_InvalidInput(t *testing.T) {
	// "123" is not valid binary (contains 2 and 3)
	input := "123"
	expected := "123"
	result := binToDec(input)
	if result != expected {
		t.Errorf("binToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestBinToDec_SingleBit(t *testing.T) {
	input := "1"
	expected := "1"
	result := binToDec(input)
	if result != expected {
		t.Errorf("binToDec(%q)\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

// ─────────────────────────────────────────────
// applyCase Tests
// ─────────────────────────────────────────────

func TestApplyCase_UpperSingleWord(t *testing.T) {
	input := []string{"hello", "world", "go"}
	mode := "up"
	command := "(up)"
	expected := []string{"hello", "world", "GO"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase up single\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_UpperMultipleWords(t *testing.T) {
	input := []string{"hello", "world", "go"}
	mode := "up"
	command := "(up,2)"
	expected := []string{"hello", "WORLD", "GO"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase up 2 words\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_LowerSingleWord(t *testing.T) {
	input := []string{"HELLO", "WORLD"}
	mode := "low"
	command := "(low)"
	expected := []string{"HELLO", "world"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase low single\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_CapitalizeSingleWord(t *testing.T) {
	input := []string{"hELLO", "wORLD"}
	mode := "cap"
	command := "(cap)"
	expected := []string{"hELLO", "World"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase cap single\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_CountLargerThanSlice(t *testing.T) {
	// count=10 but only 2 words — should handle gracefully with i<0 continue
	input := []string{"hello", "world"}
	mode := "up"
	command := "(up,10)"
	expected := []string{"HELLO", "WORLD"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase count > len\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_EmptySlice(t *testing.T) {
	input := []string{}
	mode := "up"
	command := "(up)"
	expected := []string{}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase empty slice\nExpected: %v\nGot:      %v", expected, result)
	}
}

func TestApplyCase_AllWords(t *testing.T) {
	input := []string{"a", "b", "c"}
	mode := "up"
	command := "(up,3)"
	expected := []string{"A", "B", "C"}
	result := applyCase(input, mode, command)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("applyCase all words\nExpected: %v\nGot:      %v", expected, result)
	}
}
