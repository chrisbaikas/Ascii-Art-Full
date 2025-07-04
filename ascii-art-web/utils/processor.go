package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	totalChars  = 95  // Number of supported ASCII characters
	blockLines  = 8   // Number of lines per character block
	spaceAscii  = 32  // ASCII value for space
	tildeAscii  = 126 // ASCII value for tilde
	headerLines = 1   // Header lines in banner file
)

// LoadBanner reads the banner file and returns a map of ASCII characters to their block representations
type BannerType map[rune][]string

func LoadBanner(fileName string) (BannerType, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not open banner file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %w", err)
	}

	blocks := make([][]string, 0, totalChars)
	i := headerLines
	for blockCount := 0; blockCount < totalChars; blockCount++ {
		if i+blockLines > len(lines) {
			return nil, fmt.Errorf("not enough lines for character block %d", blockCount)
		}
		block := lines[i : i+blockLines]
		blocks = append(blocks, block)
		i += blockLines
		// skip blank separator line if present
		if i < len(lines) && strings.TrimSpace(lines[i]) == "" {
			i++
		}
	}

	if len(blocks) != totalChars {
		return nil, fmt.Errorf("expected %d blocks, got %d", totalChars, len(blocks))
	}

	banner := make(BannerType)
	for j, block := range blocks {
		banner[rune(spaceAscii+j)] = block
	}
	return banner, nil
}

// AsciiArt renders the input string into ASCII art with alignment and color support
func AsciiArt(input string, banner BannerType, align string, colorTargets []ColorTarget, width int) (string, error) {
	input = strings.ReplaceAll(input, "\r", "")
	var outputBuilder strings.Builder
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			outputBuilder.WriteString("\n")
			continue
		}

		rows, err := buildAsciiRowsWithColor(line, banner, colorTargets)
		if err != nil {
			return "", err
		}

		rows, err = alignRows(rows, align, width)
		if err != nil {
			return "", err
		}

		for _, row := range rows {
			outputBuilder.WriteString(row + "\n")
		}
	}
	return outputBuilder.String(), nil
}

// buildAsciiRows converts a single line of text to ASCII art rows
func buildAsciiRows(line string, banner BannerType) ([]string, error) {
	rows := make([]string, blockLines)
	for _, ch := range line {
		if ch < spaceAscii || ch > tildeAscii {
			return nil, fmt.Errorf("unsupported character: %q", ch)
		}
		ascii, ok := banner[ch]
		if !ok || len(ascii) != blockLines {
			return nil, fmt.Errorf("character %q not found in banner", ch)
		}
		for row := 0; row < blockLines; row++ {
			rows[row] += ascii[row]
		}
	}
	return rows, nil
}
