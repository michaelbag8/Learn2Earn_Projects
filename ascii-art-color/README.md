# ASCII Art Generator

## Overview

The ASCII Art Generator is a Go application that converts normal text into ASCII art using banner templates. The program supports multiple CLI features such as color.

This project demonstrates:

* File handling in Go
* String manipulation
* Error handling
* Maps and slices
* Unit testing
* Command-line flag parsing
* ANSI color formatting


---

# Features

* Convert plain text into ASCII art
* Support for banner templates
* Multi-line input support
* Output redirection into files
* ASCII art color formatting
* Error handling for invalid input and missing files
* Unit tests for core functions

---

# Project Structure

```text
ascii-art/
├── main.go
├── banner.go
├── validate.go
├── split.go
├── render.go
├── color.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── validate_test.go
├── split_test.go
├── render_test.go
└── README.md
```

---

# Requirements

* Go 1.20 or later

Check your Go version:

```bash
go version
```

---


# Running the Project

## Normal Mode

```bash
go run . "Hello"
```

Example output:

```text
 _   _          _   _
| | | |   ___  | | | |
| |_| |  / _ \ | | | |
|  _  | |  __/ | | | |
|_| |_|  \___| |_| |_|
```

---

# Using Different Banner Files

You can use different banner styles:

```bash
go run . "Hello" shadow.txt
```

or

```bash
go run . "Hello" thinkertoy.txt
```

Available banner files:

* standard.txt
* shadow.txt
* thinkertoy.txt

---

# Command-Line Flags

| Flag                  | Description                                          |
| --------------------- | ---------------------------------------------------- |
| `--color=<color>`     | Changes ASCII art color                              |

---

# Usage Examples

## Color Mode

```bash
go run . --color=red "Hello"
```

---



# Testing

Run all tests:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

---

# Core Functions

## Validate

Checks whether all characters in the input string are printable ASCII characters.

```go
func Validate(str string) (rune, error)
```

---

## SplitInput

Splits input text using escaped newline characters.

```go
func SplitInput(str string) []string
```

---

## renderLines

Renders ASCII art lines using the selected banner map.

```go
func renderLines(segment string, blockMaps map[rune][]string) []string
```

---



## Feature 2 — `--color`

Concepts used:

* ANSI escape codes
* Color sandwich formatting:

```go
\033[31m...\033[0m
```

* `colorMap`
* `strings.ContainsRune`
* `shouldColor` logic

Example:

```bash
go run . --color=blue "Hello"
```

---


# Error Handling

The project handles:

* Invalid ASCII characters
* Missing banner files
* Incorrect input
* Invalid alignment types
* Invalid color values
* File read errors

Example:

```text
Error reading file: open notfound.txt: no such file or directory
```

---

# Learning Objectives

This project helps you practice:

* Go syntax and structure
* Loops and conditionals
* Working with maps and slices
* File operations
* CLI flag parsing
* ANSI terminal formatting
* Writing reusable functions
* Unit testing in Go

---

# Future Improvements

Possible enhancements:

* Add more banner styles
* Add Docker support
* Add live preview in browser
* Improve CLI argument validation

---

# Author

Michael BAG

---
