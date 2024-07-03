# ASCII-ART-COLOR

## Description
The ASCII Art Color is a command-line tool that allows users to colorize specific substrings within a given input string. Users can specify the desired color and the substring to be colored using a specific flag format.
Features

- Color Customization: Users can choose the color to be applied to the specified substring using the --color=<color> flag.
- Substring Targeting: Users can choose the substring to be colored within the input string.
- Whole String Coloring: If no substring is specified, the entire input string will be colored.
- Flexible Color Notation: The tool supports various color notation systems, such as RGB, HSL, and ANSI escape codes.
- Usage Message: If the flag format is incorrect, the tool will display a usage message.

## Usage
To use the ASCII Art Color tool, follow these steps:

- Ensure you have Go installed on your system.
- Clone the repository
```bash
git clone https://learn.zone01kisumu.ke/git/bbantu/ascii-art-color.git
```
```bash
cd ascii-art-color
```
Run the program appropriately with the flag
```bash
    go run . --color=<color> <substring to be colored> "input string"
```
Alternatively, the program can be run with a banner flag in this manner:
```bash
go run . --color=<color> --banner=<selected banner> <substring to be colored> "input string"
```
### Example:
```bash
go run . --color=red kit "a king kitten have kit"
```

In this example, the substrings "kit" in "kitten" and "kit" at the end of the string will be colored red.

## Implementation Details
The ASCII Art Color tool is implemented using the following key components:

- Command-line Parsing: The tool uses the standard Go flag package to parse the command-line arguments, including the color flag and the input string.
- Color Conversion: The tool supports various color notation systems, such as RGB, HSL, and ANSI escape codes. It uses color conversion functions to translate the specified color into the appropriate format for display.
- String Manipulation: The tool uses string manipulation techniques to identify the specified substring within the input string and apply the color formatting to it.
- Terminal Display: The tool outputs the colorized ASCII art to the terminal, using the appropriate ANSI escape codes to set the text color.

## Future Improvements

- Interactive Mode: Implement an interactive mode that allows users to experiment with different color and substring combinations.
- File Input: Add support for reading the input string from a file, rather than the command line.


## Contributing
Contributions to the ASCII Art Color project are welcome. If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.