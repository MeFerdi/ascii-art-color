package main

import (
	"flag"
	"fmt"
	"os"

	color "color/ascii" // Ensure this path is correct for your local setup
)

func main() {
	// Define command-line flags
	args := os.Args[1:]
	colorFlag := flag.String("color", "", "Color to apply to the ASCII art")
	flag.Parse()

	// Get the remaining arguments
	// args := flag.Args()

	// Handle the different command formats
	if len(args) == 3 {
		// Format: go run main.go --color=<color> <substring to be colored> "something"
		substringValue := args[1]
		inputString := args[2]
		handleMultipleArguments(*colorFlag, substringValue, inputString, "standard.txt")
	} else if len(args) == 4 {
		// Format: go run main.go --color=<color> <substring to be colored> "something" <banner>
		substringValue := args[1]
		inputString := args[2]
		bannerStyle := args[3]
		handleMultipleArguments(*colorFlag, substringValue, inputString, bannerStyle+".txt")
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("\nEx: go run . --color=<color> <substring to be colored> \"something\"")
		return
	}
}

// func handleSingleArgument(inputString, bannerStyle, substringValue, colorValue string) {
// 	color.PrintAsciiArt(inputString, bannerStyle, substringValue, colorValue)
// }

func handleMultipleArguments(colorValue, substringValue, inputString, bannerStyle string) {
	color.PrintAsciiArt(colorValue, substringValue, inputString, bannerStyle)
}
