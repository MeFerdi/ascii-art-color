
# Ascii-art-color

This Go program allows you to apply color to ASCII art and highlight a specific substring within the art.


### Features

- Supports various color options for the ASCII art
- Allows you to specify a banner style for the ASCII art
- Enables you to highlight a specific sunstring within the ASCII art.

## Run Locally

Clone the project

```bash
  git clone https://github.com/bbantu/ascii-art-color.git
```

Go to the project directory

```bash
  cd ascii-art-color
```

Run the project

```bash
  go run . --color=<color> <substring to be colored> "something"
```
## Note
 When running with a banner, we implemented the program to accept banner in the same format as the color.
```bash
  go run . --color=<color> --banner=<provided banner> <substring to be colored> "something"
```

-   --color: Specifies the color to apply to the ASCII art. The available colors are: default, red, green, blue, yellow, magenta, and cyan.
-    --banner: Specifies the banner style to use for the ASCII art. The available banner styles are: standard, shadow, thinkertoy, and big.
-    --substring: (Optional) Specifies the substring to be highlighted within the ASCII art.
-    <text>: The input text to be converted to ASCII art.

