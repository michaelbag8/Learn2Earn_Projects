package main

import "testing"

// ─────────────────────────────────────────────
// multiPunctuation Tests
// ─────────────────────────────────────────────

func TestMultiPunctuation_DoubleExclamation(t *testing.T) {
	input := "wow ! ! amazing"
	expected := "wow !! amazing"
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation double !\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestMultiPunctuation_DoubleQuestion(t *testing.T) {
	input := "really ? ? wow"
	expected := "really ?? wow"
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation double ?\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestMultiPunctuation_ExclamationQuestion(t *testing.T) {
	input := "what ! ? no"
	expected := "what !? no"
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation !?\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestMultiPunctuation_QuestionExclamation(t *testing.T) {
	input := "what ? ! no"
	expected := "what ?! no"
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation ?!\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestMultiPunctuation_NoMatch(t *testing.T) {
	input := "hello world"
	expected := "hello world"
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation no match\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestMultiPunctuation_EmptyString(t *testing.T) {
	input := ""
	expected := ""
	got := multiPunctuation(input)
	if got != expected {
		t.Errorf("multiPunctuation empty\nExpected: %q\nGot:      %q", expected, got)
	}
}

// ─────────────────────────────────────────────
// fixPunctuation Tests
// ─────────────────────────────────────────────

func TestFixPunctuation_SpaceBeforeComma(t *testing.T) {
	input := "hello , world"
	expected := "hello, world"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before comma\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_SpaceBeforePeriod(t *testing.T) {
	input := "hello . world"
	expected := "hello. world"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before period\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_SpaceBeforeExclamation(t *testing.T) {
	input := "hello !"
	expected := "hello!"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before !\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_SpaceBeforeQuestion(t *testing.T) {
	input := "hello ?"
	expected := "hello?"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before ?\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_SpaceBeforeColon(t *testing.T) {
	input := "note :"
	expected := "note:"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before colon\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_SpaceBeforeSemicolon(t *testing.T) {
	input := "one ; two"
	expected := "one; two"
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation space before semicolon\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixPunctuation_NoChange(t *testing.T) {
	input := "hello, world."
	expected := "hello, world."
	got := fixPunctuation(input)
	if got != expected {
		t.Errorf("fixPunctuation no change needed\nExpected: %q\nGot:      %q", expected, got)
	}
}

// ─────────────────────────────────────────────
// fixQuotes Tests
// ─────────────────────────────────────────────

func TestFixQuotes_SpaceAfterQuote(t *testing.T) {
	input := "don' t"
	expected := "don't"
	got := fixQuotes(input)
	if got != expected {
		t.Errorf("fixQuotes space after quote\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixQuotes_QuoteBeforePeriod(t *testing.T) {
	input := "he said '."
	expected := "he said'."
	got := fixQuotes(input)
	if got != expected {
		t.Errorf("fixQuotes quote before period\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixQuotes_PeriodBeforeQuote(t *testing.T) {
	input := "end. '"
	expected := "end.'"
	got := fixQuotes(input)
	if got != expected {
		t.Errorf("fixQuotes period before quote\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixQuotes_NoChange(t *testing.T) {
	input := "it's fine"
	expected := "it's fine"
	got := fixQuotes(input)
	if got != expected {
		t.Errorf("fixQuotes no change needed\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestFixQuotes_EmptyString(t *testing.T) {
	input := ""
	expected := ""
	got := fixQuotes(input)
	if got != expected {
		t.Errorf("fixQuotes empty\nExpected: %q\nGot:      %q", expected, got)
	}
}
