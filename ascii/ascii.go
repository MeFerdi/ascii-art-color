package color

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// bannerMap is a map that stores the ASCII art for different banner files
var bannerMap map[string]string

// init initializes the bannerMap and loads the ASCII art from the banner files
func init() {
	// Initialize the bannerMap
	bannerMap = make(map[string]string)

	// Load the ASCII art from the banner files
	loadBanner("standard.txt")
	loadBanner("shadow.txt")
	loadBanner("thinkertoy.txt")
}

// loadBanner reads the contents of a banner file and stores it in the bannerMap
func loadBanner(filename string) {
	// Construct the file path
	filePath := filepath.Join("./banner", filename)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Initialize an empty slice to store the lines
	var lines []string

	// Read the file line by line
	for scanner.Scan() {
		// Append each line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Check if the file is empty
	if len(lines) == 0 {
		return
	}

	// Check for any errors while reading the file
	if err := scanner.Err(); err != nil {
		// Handle error reading the file
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	// Join the lines with newline characters and store it in the bannerMap
	bannerMap[filename] = strings.Join(lines, "\n")
}

// GetLetterArray retrieves the ASCII art representation for a given character from the specified banner file
func GetLetterArray(char rune, bannerStyle string) []string {
	// Check if the banner file exists
	banner, ok := bannerMap[bannerStyle]
	if !ok {
		fmt.Println("File doesn't exist")
		os.Exit(1)
	}

	// Split the banner into lines
	alphabet := strings.Split(banner, "\n")

	// Calculate the starting index for the character
	start := (int(char) - 32) * 9

	// Check if the starting index is within the bounds of the alphabet
	if start < 0 || start >= len(alphabet) {
		return []string{}
	}

	// Slice the alphabet to get the ASCII art for the character
	arr := alphabet[start : start+9]

	// Return the ASCII art
	return arr
}

// PrintAscii prints the ASCII art representation of a given string
func PrintAscii(str, bannerStyle, color, substring string) {
	// Get the ANSI escape code for the specified color
	colorCode := GetColor(color)

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if substring == "" {
			fmt.Printf("%s%s%s\n", colorCode, line, ColorReset)
		} else {
			coloredLine := colorSubstring(line, substring, colorCode)
			fmt.Println(coloredLine)
		}
	}

	letters := [][]string{}
	for _, line := range lines {
		for _, letter := range line {
			if letter < 32 || letter > 126 {
				fmt.Printf("Non-ASCII character '%c' encountered\n", letter)
				return
			}
			arr := GetLetterArray(rune(letter), bannerStyle)
			letters = append(letters, arr)
			if letter == '\n' {
				fmt.Println()
			}
		}
	}

	// Print the ASCII art vertically
	for i := 1; i < 9; i++ {
		for _, letter := range letters {
			if len(letter) < i {
				fmt.Println("Error: File content modified")
				return
			}
			fmt.Print(letter[i])
		}
		fmt.Println()
	}
}

func colorSubstring(line, substring, color string) string {
	return strings.Replace(line, substring, fmt.Sprintf("%s%s%s", color, substring, ColorReset), -1)
}

// GetColor returns the ANSI escape code for the specified color
func GetColor(color string) string {
	switch color {
	case "red":
		return "\033[1;31m"
	case "green":
		return "\033[1;32m"
	case "yellow":
		return "\033[1;33m"
	case "blue":
		return "\033[1;34m"
	case "magenta":
		return "\033[1;35m"
	case "cyan":
		return "\033[1;36m"
	case "white":
		return "\033[1;37m"
	default:
		return "\033[0m" // Reset color
	}
}
