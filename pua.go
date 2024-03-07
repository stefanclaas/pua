package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
)

var letterToPUA = make(map[rune]rune)
var puaToLetter = make(map[rune]rune)

func init() {
	for i, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=" {
		pua := rune(0xF0000 + i)
		letterToPUA[letter] = pua
		puaToLetter[pua] = letter
	}
}

func encode(input string) string {
	var result []rune
	for _, letter := range input {
		if pua, ok := letterToPUA[letter]; ok {
			result = append(result, pua)
		} else {
			result = append(result, letter)
		}
	}
	return string(result)
}

func decode(input string) string {
	var result []rune
	for _, pua := range input {
		if letter, ok := puaToLetter[pua]; ok {
			result = append(result, letter)
		} else {
			result = append(result, pua)
		}
	}
	return string(result)
}

func main() {
	decodePtr := flag.Bool("d", false, "Set this flag to decode")
	fontPath := flag.String("f", "Path/to/font.ttf", "Path to the .ttf font file")
	flag.Parse()

	// Read the .ttf file
	fontBytes, err := ioutil.ReadFile(*fontPath)
	if err != nil {
		fmt.Println("Error reading the font file:", err)
		return
	}

	// Load the font
	_, err = freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("Error parsing the font:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if *decodePtr {
			fmt.Println(decode(input))
		} else {
			fmt.Println(encode(input))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

