package utils

import (
	"fmt"
	"os"
	"strings"
)

// ReverseAscii reads an ASCII-art file and recovers the original text
// by matching horizontally against known ASCII glyphs from a banner map.
func ReverseAscii(fileName string, banner map[rune][]string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("could not read reverse file: %w", err)
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	if len(lines)%blockLines != 0 {
		lines = padToMultipleOfBlockLines(lines)
	}

	var result []string
	for i := 0; i+blockLines <= len(lines); i += blockLines {
		block := lines[i : i+blockLines]
		text := decodeAsciiBlock(block, banner)
		result = append(result, strings.TrimRight(text, " "))
	}

	return strings.Join(result, "\n"), nil
}

// decodeAsciiBlock takes 8 lines and matches them left-to-right against banner chars.
func decodeAsciiBlock(block []string, banner map[rune][]string) string {
	var decoded strings.Builder
	lines := make([]string, blockLines)
	copy(lines, block)

	for {
		if allLinesEmpty(lines) {
			break
		}

		matchChar, matchWidth := findMatchingChar(lines, banner)
		if matchChar == 0 && matchWidth == 0 {
			decoded.WriteRune(' ')
			for i := 0; i < blockLines; i++ {
				if len(lines[i]) > 0 {
					lines[i] = lines[i][1:]
				}
			}
		} else {
			decoded.WriteRune(matchChar)
			for i := 0; i < blockLines; i++ {
				lines[i] = lines[i][matchWidth:]
			}
		}
	}

	return decoded.String()
}

// findMatchingChar tries to match the left part of lines with any banner character.
func findMatchingChar(lines []string, banner map[rune][]string) (rune, int) {
	for ch, glyph := range banner {
		if len(glyph) != blockLines {
			continue
		}

		width := len(glyph[0])
		if !hasMinWidth(lines, width) {
			continue
		}

		matched := true
		for i := 0; i < blockLines; i++ {
			if lines[i][:width] != glyph[i] {
				matched = false
				break
			}
		}
		if matched {
			return ch, width
		}
	}
	return 0, 0
}

func hasMinWidth(lines []string, w int) bool {
	for _, line := range lines {
		if len(line) < w {
			return false
		}
	}
	return true
}

func allLinesEmpty(lines []string) bool {
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			return false
		}
	}
	return true
}

// padToMultipleOfBlockLines pads with empty lines so the total is a multiple of blockLines.
func padToMultipleOfBlockLines(lines []string) []string {
	remainder := len(lines) % blockLines
	if remainder == 0 {
		return lines
	}
	padding := blockLines - remainder
	for i := 0; i < padding; i++ {
		lines = append(lines, "")
	}
	return lines
}
