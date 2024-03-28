package main

import (
	"fmt"
	"os"
	"strings"
)

var codes = map[string]string{
	"1": "⠼⠁",
	"2": "⠼⠃",
	"3": "⠼⠉",
	"4": "⠼⠙",
	"5": "⠼⠑",
	"6": "⠼⠋",
	"7": "⠼⠛",
	"8": "⠼⠓",
	"9": "⠼⠊",
	"0": "⠼⠚",
	"A": "⠠⠁",
	"B": "⠠⠃",
	"C": "⠠⠉",
	"D": "⠠⠙",
	"E": "⠠⠑",
	"F": "⠠⠋",
	"G": "⠠⠛",
	"H": "⠠⠓",
	"I": "⠠⠊",
	"J": "⠠⠚",
	"K": "⠠⠅",
	"L": "⠠⠇",
	"M": "⠠⠍",
	"N": "⠠⠝",
	"O": "⠠⠕",
	"P": "⠠⠏",
	"Q": "⠠⠟",
	"R": "⠠⠗",
	"S": "⠠⠎",
	"T": "⠠⠞",
	"U": "⠠⠥",
	"V": "⠠⠧",
	"W": "⠠⠺",
	"X": "⠠⠭",
	"Y": "⠠⠽",
	"Z": "⠠⠵",
	"a": "⠁",
	"b": "⠃",
	"c": "⠉",
	"d": "⠙",
	"e": "⠑",
	"f": "⠋",
	"g": "⠛",
	"h": "⠓",
	"i": "⠊",
	"j": "⠚",
	"k": "⠅",
	"l": "⠇",
	"m": "⠍",
	"n": "⠝",
	"o": "⠕",
	"p": "⠏",
	"q": "⠟",
	"r": "⠗",
	"s": "⠎",
	"t": "⠞",
	"u": "⠥",
	"v": "⠧",
	"w": "⠺",
	"x": "⠭",
	"y": "⠽",
	"z": "⠵",
	".": "⠲",
	",": "⠂",
	"?": "⠦",
	";": "⠆",
	"!": "⠖",
	"<": "⠦",
	">": "⠴",
	" ": "⠶",
	"-": "⠤",
	"'": "⠄",
}

func toBraille(text string) string {
	var brailleText strings.Builder
	for _, r := range text {
		if val, ok := codes[string(r)]; ok {
			brailleText.WriteString(val)
		}
	}
	return brailleText.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: text2braille <text>")
		os.Exit(1)
	}

	text := os.Args[1]
	for _, arg := range os.Args[2:] { // Concatenate all arguments to handle spaces in input text
		text += " " + arg
	}

	braille := toBraille(text)
	fmt.Println(braille)
}
