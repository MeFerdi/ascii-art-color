package main

import (
	"flag"
	"fmt"

	color "color/ascii" // Import the color package from the local directory
)

func main() {
	// Define command-line flags
	colorFlag := flag.String("color", "", "Color to apply to the ASCII art")
	// substringFlag := flag.String("substring", "", "Substring to be colored")
	flag.Parse()

	// Get the remaining arguments
	args := flag.Args()

	// Handle the different command formats
	switch len(args) {
	case 2:
		// Format: go run main.go <string> <banner>
		inputString := args[0]
		bannerStyle := args[1] + ".txt"
		color.PrintAsciiArt(inputString, bannerStyle, "", "")
	case 1:
		// Format: go run main.go <string>
		inputString := args[0]
		bannerStyle := "standard.txt"
		color.PrintAsciiArt(inputString, bannerStyle, "", "")
	case 3:
		// Format: go run . --color=<color> <substring to be colored> "something"
		// Extract the color and substring values from the flags
		substringValue := args[1]
		inputString := args[2]
		bannerStyle := "standard.txt"
		colorValue := *colorFlag
		color.PrintAsciiArt(colorValue, substringValue, inputString, bannerStyle)
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("\nEx: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}
}
