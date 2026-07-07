package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func genQuadA(w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	var lines []string
	for r := 0; r < h; r++ {
		if r == 0 || r == h-1 {
			if w == 1 {
				lines = append(lines, "o")
			} else {
				lines = append(lines, "o"+strings.Repeat("-", w-2)+"o")
			}
		} else {
			if w == 1 {
				lines = append(lines, "|")
			} else {
				lines = append(lines, "|"+strings.Repeat(" ", w-2)+"|")
			}
		}
	}
	return strings.Join(lines, "\n")
}

func genQuadB(w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	var lines []string
	for r := 0; r < h; r++ {
		if r == 0 {
			if w == 1 {
				lines = append(lines, "/")
			} else {
				lines = append(lines, "/"+strings.Repeat("*", w-2)+"\\")
			}
		} else if r == h-1 {
			if w == 1 {
				lines = append(lines, "\\")
			} else {
				lines = append(lines, "\\"+strings.Repeat("*", w-2)+"/")
			}
		} else {
			if w == 1 {
				lines = append(lines, "*")
			} else {
				lines = append(lines, "*"+strings.Repeat(" ", w-2)+"*")
			}
		}
	}
	return strings.Join(lines, "\n")
}

func genQuadC(w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	var lines []string
	for r := 0; r < h; r++ {
		if r == 0 {
			if w == 1 {
				lines = append(lines, "A")
			} else {
				lines = append(lines, "A"+strings.Repeat("B", w-2)+"A")
			}
		} else if r == h-1 {
			if w == 1 {
				lines = append(lines, "C")
			} else {
				lines = append(lines, "C"+strings.Repeat("B", w-2)+"C")
			}
		} else {
			if w == 1 {
				lines = append(lines, "B")
			} else {
				lines = append(lines, "B"+strings.Repeat(" ", w-2)+"B")
			}
		}
	}
	return strings.Join(lines, "\n")
}

func genQuadD(w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	var lines []string
	for r := 0; r < h; r++ {
		if r == 0 {
			if w == 1 {
				lines = append(lines, "A")
			} else {
				lines = append(lines, "A"+strings.Repeat("B", w-2)+"C")
			}
		} else if r == h-1 {
			if w == 1 {
				lines = append(lines, "A")
			} else {
				lines = append(lines, "A"+strings.Repeat("B", w-2)+"C")
			}
		} else {
			if w == 1 {
				lines = append(lines, "B")
			} else {
				lines = append(lines, "B"+strings.Repeat(" ", w-2)+"B")
			}
		}
	}
	return strings.Join(lines, "\n")
}

func genQuadE(w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	var lines []string
	for r := 0; r < h; r++ {
		if r == 0 {
			if w == 1 {
				lines = append(lines, "A")
			} else {
				lines = append(lines, "A"+strings.Repeat("B", w-2)+"C")
			}
		} else if r == h-1 {
			if w == 1 {
				lines = append(lines, "C")
			} else {
				lines = append(lines, "C"+strings.Repeat("B", w-2)+"A")
			}
		} else {
			if w == 1 {
				lines = append(lines, "B")
			} else {
				lines = append(lines, "B"+strings.Repeat(" ", w-2)+"B")
			}
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading stdin:", err)
		os.Exit(1)
	}
	input := string(raw)
	if strings.HasSuffix(input, "\n") {
		input = strings.TrimRight(input, "\n")
	}
	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	lines := strings.Split(input, "\n")
	w := len(lines[0])
	h := len(lines)
	for _, L := range lines {
		if len(L) != w {
			fmt.Println("Not a quad function")
			return
		}
	}
	matches := []string{}
	if input == genQuadA(w, h) {
		matches = append(matches, "quadA")
	}
	if input == genQuadB(w, h) {
		matches = append(matches, "quadB")
	}
	if input == genQuadC(w, h) {
		matches = append(matches, "quadC")
	}
	if input == genQuadD(w, h) {
		matches = append(matches, "quadD")
	}
	if input == genQuadE(w, h) {
		matches = append(matches, "quadE")
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	sort.Strings(matches)
	out := []string{}
	for _, m := range matches {
		out = append(out, fmt.Sprintf("[%s] [%d] [%d]", m, w, h))
	}
	fmt.Println(strings.Join(out, " || "))
}
