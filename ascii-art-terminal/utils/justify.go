package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// getTerminalWidth retrieves the current terminal width using stty command
func getTerminalWidth() int {
	cmd := exec.Command("stty", "size") // Command to get terminal size
	cmd.Stdin = os.Stdin                // Connect to standard input
	out, _ := cmd.Output()              // Execute command and get output
	// Parse output (format: "rows columns") and convert width to integer
	width, _ := strconv.Atoi(strings.Fields(string(out))[1])
	return width
}

// stretching only the spaces between words to fill exactly 'width' columns.
func justifyAscii(line string, banner map[rune][]string, width int, colorTargets []ColorTarget) ([]string, error) {
	words := strings.Fields(line)
	n := len(words)
	if n == 0 {
		return make([]string, blockLines), nil
	}

	// Render each word with color support
	wordRows := make([][]string, n)
	for i, w := range words {
		rows, err := buildAsciiRowsWithColor(w, banner, colorTargets)
		if err != nil {
			return nil, fmt.Errorf("error building ASCII for word %q: %w", w, err)
		}
		wordRows[i] = rows
	}

	// Measure visible length of each word (excluding ANSI codes)
	totalWordLen := 0
	for _, wr := range wordRows {
		totalWordLen += visibleLength(wr[0]) // all rows have same length
	}

	// Determine spacing
	slots := n - 1
	extra := width - totalWordLen
	if slots <= 0 || extra <= 0 {
		// fallback to left-align
		return buildAsciiRowsWithColor(line, banner, colorTargets)
	}

	base := extra / slots
	rem := extra % slots
	gaps := make([]int, slots)
	for i := range gaps {
		gaps[i] = base
		if i < rem {
			gaps[i]++
		}
	}

	// Stitch words and gaps identically across rows
	result := make([]string, blockLines)
	for row := 0; row < blockLines; row++ {
		for wi, wr := range wordRows {
			result[row] += wr[row]
			if wi < slots {
				result[row] += strings.Repeat(" ", gaps[wi])
			}
		}
	}

	return result, nil
}

// alignRows adjusts the position of ASCII art based on alignment type
func alignRows(rows []string, align string, width int) ([]string, error) {
	if align == "left" {
		return rows, nil
	}

	// Use visible length (excluding ANSI escape codes)
	lineLen := visibleLength(rows[0])
	if width < lineLen {
		return nil, fmt.Errorf("terminal width too small for alignment")
	}

	// Calculate padding for center or right alignment
	pad := 0
	switch align {
	case "center":
		pad = (width - lineLen) / 2
	case "right":
		pad = width - lineLen
	default:
		return nil, fmt.Errorf("unknown alignment: %s", align)
	}

	// Prepend spaces to each row
	for i := range rows {
		rows[i] = strings.Repeat(" ", pad) + rows[i]
	}
	return rows, nil
}
