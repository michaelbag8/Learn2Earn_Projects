package main

import (
	"testing"
)

func TestFixArticles_BasicVowel(t *testing.T) {
	input := "I saw a elephant"
	expected := "I saw an elephant"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test BasicVowel FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_UppercaseA(t *testing.T) {
	input := "A elephant walked by"
	expected := "An elephant walked by"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test UppercaseA FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_MultipleArticles(t *testing.T) {
	input := "I saw a elephant and a orange"
	expected := "I saw an elephant and an orange"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test MultipleArticles FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_NoArticleNeeded(t *testing.T) {
	// "a cat" — 'c' is not a vowel, should NOT change
	input := "I have a cat"
	expected := "I have a cat"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test NoArticleNeeded FAILED\nInput:    %q\nExpected: %q\nGot:      %q",
			input, expected, result)
	}
}

func TestFixArticles_HWord(t *testing.T) {
	// "h" is in the vowels string, so "a hour" → "an hour"
	input := "Wait a hour"
	expected := "Wait an hour"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test HWord FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_EmptyString(t *testing.T) {
	input := ""
	expected := ""
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test EmptyString FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_SingleWord(t *testing.T) {
	// Only one word — loop never runs (i < len-1 = 0 is false from start)
	input := "elephant"
	expected := "elephant"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test SingleWord FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_ArticleAtEnd(t *testing.T) {
	// "a" is the last word — loop stops before it, so it won't be touched
	input := "I lost a"
	expected := "I lost a"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test ArticleAtEnd FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}

func TestFixArticles_MixedCase(t *testing.T) {
	input := "She is A honest person"
	expected := "She is An honest person"
	result := fixArticles(input)
	if result != expected {
		t.Errorf("Test MixedCase FAILED\nInput:    %q\nExpected: %q\nGot:      %q", input, expected, result)
	}
}
