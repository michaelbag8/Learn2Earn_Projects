package main

import (
	"os"
	"path/filepath"
	"testing"
)

// ─────────────────────────────────────────────
// Helper: create a temp file with content
// ─────────────────────────────────────────────

func writeTempFile(t *testing.T, content string) string {
	t.Helper() // marks this as a helper so error lines point to the caller
	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("could not create temp input file: %v", err)
	}
	return path
}

func readTempFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}
	return string(data)
}

// ─────────────────────────────────────────────
// run() Tests
// ─────────────────────────────────────────────

func TestRun_BasicPassthrough(t *testing.T) {
	// No commands — text should pass through unchanged (except trailing newline)
	inputPath := writeTempFile(t, "hello world")
	outputPath := filepath.Join(t.TempDir(), "output.txt")

	if err := run(inputPath, outputPath); err != nil {
		t.Fatalf("run() returned unexpected error: %v", err)
	}

	got := readTempFile(t, outputPath)
	expected := "hello world\n"
	if got != expected {
		t.Errorf("run() basic\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestRun_InvalidInputFile(t *testing.T) {
	// Non-existent file — should return an error
	err := run("nonexistent_file.txt", "output.txt")
	if err == nil {
		t.Error("Expected error for missing input file, got nil")
	}
}

func TestRun_WritesNewlineAtEnd(t *testing.T) {
	inputPath := writeTempFile(t, "hello")
	outputPath := filepath.Join(t.TempDir(), "output.txt")

	if err := run(inputPath, outputPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readTempFile(t, outputPath)
	if got[len(got)-1] != '\n' {
		t.Errorf("Expected output to end with newline, got: %q", got)
	}
}

// ─────────────────────────────────────────────
// processText() Tests
// ─────────────────────────────────────────────

func TestProcessText_NoCommands(t *testing.T) {
	input := "hello world"
	expected := "hello world"
	got := processText(input)
	if got != expected {
		t.Errorf("processText no commands\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_HexCommand(t *testing.T) {
	input := "1f (hex)"
	expected := "31"
	got := processText(input)
	if got != expected {
		t.Errorf("processText hex\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_BinCommand(t *testing.T) {
	input := "1011 (bin)"
	expected := "11"
	got := processText(input)
	if got != expected {
		t.Errorf("processText bin\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_UpperCommand(t *testing.T) {
	input := "hello world (up)"
	expected := "hello WORLD"
	got := processText(input)
	if got != expected {
		t.Errorf("processText up\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_UpperCommandWithCount(t *testing.T) {
	input := "hello world go (up,2)"
	expected := "hello WORLD GO"
	got := processText(input)
	if got != expected {
		t.Errorf("processText up,2\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_LowCommand(t *testing.T) {
	input := "HELLO WORLD (low)"
	expected := "HELLO world"
	got := processText(input)
	if got != expected {
		t.Errorf("processText low\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_CapCommand(t *testing.T) {
	input := "hELLO wORLD (cap)"
	expected := "hELLO World"
	got := processText(input)
	if got != expected {
		t.Errorf("processText cap\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_SplitCommand(t *testing.T) {
	// "(up," and "2)" are split across two tokens
	input := "hello world go (up, 2)"
	expected := "hello WORLD GO"
	got := processText(input)
	if got != expected {
		t.Errorf("processText split command\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_HexAtStart(t *testing.T) {
	// (hex) with empty result — should not panic
	input := "(hex) hello"
	expected := "hello"
	got := processText(input)
	if got != expected {
		t.Errorf("processText hex at start\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_ArticleFixApplied(t *testing.T) {
	input := "I saw a elephant"
	expected := "I saw an elephant"
	got := processText(input)
	if got != expected {
		t.Errorf("processText article fix\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_MultipleCommands(t *testing.T) {
	input := "hello 1f (hex) world (up,2) I saw a elephant"
	expected := "hello 31 WORLD I saw an elephant"
	got := processText(input)
	if got != expected {
		t.Errorf("processText multiple commands\nExpected: %q\nGot:      %q", expected, got)
	}
}

func TestProcessText_EmptyInput(t *testing.T) {
	input := ""
	expected := ""
	got := processText(input)
	if got != expected {
		t.Errorf("processText empty\nExpected: %q\nGot:      %q", expected, got)
	}
}
