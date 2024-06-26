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
	bannerFlag := flag.String("banner", "standard", "Banner style to use")
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
			if strings.Contains(strings.ToLower(word), strings.ToLower(*substringFlag)) {
				substring = word
				break
			}
		}
	}

	// Construct the banner file name
	bannerFile := *bannerFlag + ".txt"

	// Call the PrintAscii function from the color package with the provided options
	color.PrintAscii(inputString, bannerFile, *colorFlag, substring)
}
