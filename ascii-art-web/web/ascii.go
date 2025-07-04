// Generates ASCII art output with optional color and alignment

package web

import (
	"fmt"

	"platform.zone01.gr/git/askordal/ascii-art-web-export-file/utils"
)

// Struct to hold parsed ASCII parameters from the form
type asciiRequest struct {
	Text         string
	Banner       string
	Align        string
	GlobalColor  string
	ColorTargets []string
	TargetColors []string
}

// GenerateAsciiArt generates ASCII art with colors and alignment (now exported)
func generateAsciiArt(p *asciiRequest) (string, error) {
	bannerMap, ok := LoadedBanners[p.Banner]
	if !ok {
		return "", fmt.Errorf("internal: failed to load banner %q", p.Banner)
	}

	targets := []utils.ColorTarget{}

	// Handle targeted colors first
	for i, substr := range p.ColorTargets {
		color := ""
		if i < len(p.TargetColors) && p.TargetColors[i] != "" {
			color = p.TargetColors[i]
		}
		if substr != "" && color != "" {
			targets = append(targets, utils.ColorTarget{
				ColorCode: color,
				Substring: substr,
			})
		}
	}

	// Handle global color (prepended so it's applied first)
	if p.GlobalColor != "" {
		targets = append([]utils.ColorTarget{{
			ColorCode: p.GlobalColor,
			Substring: "",
		}}, targets...)
	}

	ascii, err := utils.AsciiArt(p.Text, bannerMap, p.Align, targets, 150)
	if err != nil {
		return "", fmt.Errorf("error generating ASCII art: %w", err)
	}
	html, err := utils.AnsiToHTML256(ascii)
	if err != nil {
		return "", fmt.Errorf("error converting ANSI to HTML: %w", err)
	}
	return html, nil
}
