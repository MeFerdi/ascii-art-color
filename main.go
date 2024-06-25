package main

import (
	"flag"
	"fmt"
	"strings"

	color "color/ascii" // Import the color package from the local directory
)

func main() {
	colorFlag := flag.String("color", "default", "Color to apply to the ASCII art")
	substringFlag := flag.String("substring", "", "Substring to be colored")
	bannerFlag := flag.String("banner", "standard.txt", "Banner style to use")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . --color=<color> --banner=<banner> [--substring=<substring>] <text>")
		return
	}

	inputString := strings.Join(flag.Args(), " ")
	if inputString == "" {
		return
	}

	// Get the substring to be colored from the command-line arguments
	var substring string
	if *substringFlag != "" {
		substring = *substringFlag
	} else {
		words := strings.Split(inputString, " ")
		for _, word := range words {
			if strings.Contains(word, "kit") {
				substring = word
				break
			}
		}
	}

	// Call the PrintAscii function from the color package with the provided options
	color.PrintAscii(inputString, *bannerFlag, *colorFlag, substring)
}
