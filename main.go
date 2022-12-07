package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Define the flags
var outputFile = flag.String("output", "output.txt", "the file to write output to")

func main() {
	// Parse the flags
	flag.Parse()
	// Access the non-flag arguments
	args := flag.Args()
	fonts := ""

	const usage = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
	if len(args) <= 0 || len(args) > 2 {
		fmt.Println(usage)
		return
	}

	// Chose default font when not chosen in cli
	if len(args) == 1 {
		fonts = "fonts/standard.txt"
	} else {
		switch args[1] {

		case "standard", "standard.txt":
			fonts = "fonts/standard.txt"

		case "shadow", "shadow.txt":
			fonts = "fonts/shadow.txt"

		case "tinkertoy", "tinkertoy.txt":
			fonts = "fonts/tinkertoy.txt"
		}
	}

	// Check the new line and if there is nothing at the beginning of the input, map user input
	if args[0] == "\\n" {
		fmt.Printf("\n")
	} else if args[0] != "" {
		asciiMap, err := mapFonts(fonts)
		if err != nil {
			fmt.Println(err)
			return
		}
		mappedInput, err := mapUserInput(args[0], asciiMap)
		if err != nil {
			fmt.Println(err)
			return
		}
		writeOutput(mappedInput)
	}
}

// Read the font and map it to ascii
func mapFonts(font string) (map[int][]string, error) {
	readFile, err := os.ReadFile(font)
	if err != nil {
		return nil, fmt.Errorf("could not read the content in the file: %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}
	return ascii, nil
}

// Handle user input and map it to ascii
func mapUserInput(input string, ascii map[int][]string) (string, error) {
	lines := strings.Split(input, "\\n")
	output := ""

	for _, line := range lines {
		characters := []rune(line)

		if line != "" {
			for i := 0; i < 8; i++ {
				for _, ch := range characters {
					if ch >= 32 && ch <= 126 {
						output = output + ascii[int(ch)][i]
					} else {
						return "", fmt.Errorf("input includes non ascii character(s), please use ascii character(s)")
					}
				}
				output = output + "\n"
			}
		} else {
			output = output + "\n"
		}
	}
	return output, nil
}

// Write an output to a new file with the specified file name
func writeOutput(s string) {
	// Create a new file with the specified file name
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}
}
