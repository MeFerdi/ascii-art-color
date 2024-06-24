package main

import (
	"flag"
	"fmt"

	color "color/ascii" // Import the color package from the local directory
)

func main() {
	colorFlag := flag.String("color", "red", "Color to apply to the ASCII art")
	substringFlag := flag.String("substring", "", "Substring to be colored")
	bannerFlag := flag.String("banner", "standard.txt", "Banner style to use")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: go run . [--color=<color>] [--substring=<substring>] [--banner=<banner_style>] <input_string>")
		return
	}

	inputString := flag.Arg(0)
	if inputString == "" {
		return
	}

	// Call the PrintAscii function from the color package
	color.PrintAscii(inputString, *bannerFlag, *colorFlag, *substringFlag)
}
