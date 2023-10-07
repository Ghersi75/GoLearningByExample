package main

import (
	"fmt"
	"unicode/utf8"
)

// Cool and useful article with examples of how runes work and are stored
// https://www.educative.io/answers/what-is-the-rune-type-in-golang
func StringsAndRunes() {
	const s = "สวัสดี"

	// 18 len here because of how utf-8 encoding works
	fmt.Println("String:", s, "Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// single quotes for char else it will throw error bc double quotes is string, just like C
	// Basically runes are chars/ints used to store all possible utf-8 values and chars
	if r == 't' {
		fmt.Println("Found t")
	} else if r == 'ส' {
		fmt.Println("Found ส")
	}
}