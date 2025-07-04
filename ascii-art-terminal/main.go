package main

import (
	"fmt"
	"os"
	"strings"

	"platform.zone01.gr/git/askordal/ascii-art-reverse/utils"
)

func main() {
	// Handle --reverse flag first
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "--reverse=") {
		fileName := strings.TrimPrefix(os.Args[1], "--reverse=")

		// Load default banner for reverse (standard.txt)
		banner, err := utils.LoadBanner("standard.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error loading banner:", err)
			os.Exit(1)
		}

		// Attempt to decode the file
		text, err := utils.ReverseAscii(fileName, banner)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reversing ASCII:", err)
			os.Exit(1)
		}

		fmt.Println(text)
		return
	}

	outputFile, alignType, inputText, bannerFile, colorTargets, err := utils.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	banner, err := utils.LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, utils.UsageMsg)
		os.Exit(1)
	}

	asciiArt, err := utils.AsciiArt(inputText, banner, alignType, colorTargets)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if outputFile == "" {
		fmt.Print(asciiArt)
	} else {
		// εδώ strip + write
		rows := strings.Split(asciiArt, "\n")
		if err := utils.WriteToFile(rows, outputFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error writing to file:", err)
			os.Exit(1)
		}
	}
}
