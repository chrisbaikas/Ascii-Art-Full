package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	totalChars  = 95  // number of ASCII characters supported
	blockLines  = 8   // height of each character block in lines
	spaceAscii  = 32  // starting ASCII code for banners
	tildeAscii  = 126 // ending ASCII code for banners
	headerLines = 1   // number of header lines in banner files
)

// LoadBanner opens a banner file and parses it into ASCII art blocks
func LoadBanner(fileName string) (map[rune][]string, error) {
	file, err := os.Open(fileName) // open banner file
	if err != nil {
		return nil, fmt.Errorf("could not open banner file: %w", err)
	}
	defer file.Close() // ensure file is closed

	scanner := bufio.NewScanner(file) // read lines
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // collect lines
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %w", err)
	}

	// split into blocks of blockLines, skipping blank separators
	blocks := make([][]string, 0, totalChars)
	i := headerLines
	for b := 0; b < totalChars; b++ {
		if i+blockLines > len(lines) {
			return nil, fmt.Errorf("not enough lines for character %d", b)
		}
		blocks = append(blocks, lines[i:i+blockLines]) // extract block
		i += blockLines
		if i < len(lines) && strings.TrimSpace(lines[i]) == "" {
			i++ // skip blank line
		}
	}

	// map ASCII codes to their art blocks
	banner := make(map[rune][]string)
	for idx, block := range blocks {
		banner[rune(spaceAscii+idx)] = block
	}
	return banner, nil
}

// AsciiArt renders input text to ASCII art respecting alignment and colours
func AsciiArt(input string, banner map[rune][]string, align string, colorTargets []ColorTarget) (string, error) {
	width := getTerminalWidth()
	var out strings.Builder

	if input == "" { // absolutely empty: no output
		return "", nil
	}

	// keep the trailing \n tokens so we know exactly how many blank lines the user asked for
	chunks := strings.SplitAfter(input, "\n")
	for _, chunk := range chunks {
		if chunk == "\n" { // explicit blank line → single raw newline out
			out.WriteString("\n")
			continue
		}

		// remove the trailing newline (if any) so we can process the text itself
		line := strings.TrimSuffix(chunk, "\n")
		if line == "" { // ignore empty tail produced by SplitAfter
			continue
		}

		// JUSTIFY handling
		if align == "justify" && len(strings.Fields(line)) > 1 {
			rows, err := justifyAscii(line, banner, width, colorTargets)
			if err != nil {
				return "", err
			}
			for _, r := range rows {
				out.WriteString(r + "\n")
			}
			continue
		} else if align == "justify" {
			fmt.Fprintln(os.Stderr, "\033[33mwarning: cannot justify line with one word, using left align\033[0m")
			align = "left"
		}

		// Normal rendering path
		rows, err := buildAsciiRowsWithColor(line, banner, colorTargets)
		if err != nil {
			return "", err
		}
		rows, err = alignRows(rows, align, width)
		if err != nil {
			return "", err
		}
		for _, r := range rows {
			out.WriteString(r + "\n")
		}
	}

	return out.String(), nil
}

// buildAsciiRows converts a single line to ASCII art rows
func buildAsciiRows(line string, banner map[rune][]string) ([]string, error) {
	rows := make([]string, blockLines) // prepare rows
	for _, ch := range line {
		if ch < spaceAscii || ch > tildeAscii {
			return nil, fmt.Errorf("unsupported char: %q", ch) // validate char
		}
		ascii, ok := banner[ch]
		if !ok || len(ascii) != blockLines {
			return nil, fmt.Errorf("char %q missing banner data", ch)
		}
		for i := 0; i < blockLines; i++ {
			rows[i] += ascii[i] // build each row
		}
	}
	return rows, nil
}
