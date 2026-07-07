# ASCII Art Generator

## Overview

The ASCII Art Generator is a Go application that converts normal text into ASCII art using banner templates. The program supports multiple CLI features such as color formatting, alignment, file output, reverse ASCII decoding, and a web interface mode.

This project demonstrates:

* File handling in Go
* String manipulation
* Error handling
* Maps and slices
* Unit testing
* Command-line flag parsing
* ANSI color formatting
* Terminal width handling
* HTTP server development in Go

---

# Features

* Convert plain text into ASCII art
* Support for banner templates
* Multi-line input support
* Output redirection into files
* ASCII art color formatting
* Text alignment support
* Reverse ASCII art decoding
* Web server interface
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
├── output.go
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
go run . "Hello" shadow
```

or

```bash
go run . "Hello" thinkertoy
```

Available banner files:

* standard.txt
* shadow.txt
* thinkertoy.txt

---

# Command-Line Flags

| Flag                  | Description                                          |
| --------------------- | ---------------------------------------------------- |
| `--output=<file>`     | Saves output into a file                             |


---

# Usage Examples


## Output Mode

Save ASCII art into a file:

```bash
go run . --output=test.txt "Hello"
```

Read the saved output:

```bash
cat test.txt
```

---

## Reverse Mode

Generate ASCII art into a file:

```bash
go run . --output=test.txt "Hello\nWorld"
```

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
* HTTP server development
* Writing reusable functions
* Unit testing in Go

---

# Future Improvements

Possible enhancements:

* Add more banner styles
* Add downloadable web output
* Improve reverse detection accuracy
* Add Docker support
* Add live preview in browser
* Improve CLI argument validation

---

# Author

Michael BAG

---
