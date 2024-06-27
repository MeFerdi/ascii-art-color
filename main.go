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
	bannerFlag := flag.String("banner", "standard", "Banner style to use (without .txt extension)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("\nEX: go run . --color=<color> <substring to be colored> \"something\"")
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

	// Construct the banner file name by appending the ".txt" extension
	bannerFile := *bannerFlag + ".txt"

	// Call the PrintAscii function from the color package with the provided options
	color.PrintAsciiArt(inputString, bannerFile, *colorFlag, substring)
}
