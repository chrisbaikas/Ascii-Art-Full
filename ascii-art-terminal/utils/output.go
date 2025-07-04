package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// StripANSI removes ANSI escape sequences.
func StripANSI(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); {
		if s[i] == '\x1b' && i+1 < len(s) && s[i+1] == '[' {
			i += 2
			for i < len(s) && s[i] != 'm' {
				i++
			}
			i++ // skip 'm'
		} else {
			b.WriteByte(s[i])
			i++
		}
	}
	return b.String()
}

// WriteToFile strips ANSI codes and writes lines to a file (adds .txt if needed).
func WriteToFile(rows []string, path string) error {
	if filepath.Ext(path) == "" {
		path += ".txt"
	}
	hasANSI := false
	for _, r := range rows {
		if strings.IndexByte(r, '\x1b') >= 0 {
			hasANSI = true
			break
		}
	}
	if hasANSI {
		fmt.Fprintf(os.Stderr, "\x1b[31mwarning: stripping ANSI codes to %s\x1b[0m\n", path)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, r := range rows {
		if _, err := f.WriteString(StripANSI(r) + "\n"); err != nil {
			return err
		}
	}
	return nil
}
