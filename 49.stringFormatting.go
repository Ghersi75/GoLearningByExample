package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y 	int
}

var pf = fmt.Printf

func StringFormatting() {
	p := point{1, 2}
	pf("struct1 (%%v): %v\n", p)
	pf("struct2 (%%+v): %+v\n", p)
	pf("struct3 (%%#v): %#v\n", p)

	pf("type (%%T): %T\n", p)

	pf("bool (%%t): %t\n", true)

	pf("int (%%d): %d\n", 123)

	pf("binary (%%b): %b\n", 14)

	pf("char equivalent of int (%%c): %c\n", 33)

	pf("hexadecimal (%%x): %x\n", 456)

	pf("float1 (%%f): %f\n", 78.9)
	pf("float2 scientific e (%%e): %e\n", 780000000000.9)
	pf("float3 scientific E (%%E): %E\n", 780000000000.9)

	pf("str1 regular string (%%s): %s\n", "\"string\"")
	pf("str2 print exactly what's there (%%q): %q\n", "\"string\"")
	pf("str3 hex string (%%x): %x\n", "hex this")

	pf("pointer (%%p): %p\n", &p)

	// Show at least 6 spaces for the table
	// By default it's right justified and padded on the left with spaces
	pf("width1 (|%%6d|%%6d|): |%6d|%6d|\n", 122323232, 345)
	// Define float with how many decimals and it will round automatically
	pf("width2 (|%%6.2f|%%6.2f|): |%6.2f|%6.2f|\n", 123.342454, 12.445)
	// - to left justify
	pf("width3 (|%%-10.2f|%%-10.2f|): |%-10.2f|%-10.2f|\n", 123.342454, 12.445)
	// works for strings too to format a table
	pf("width4 (|%%6s|%%6s|): |%6s|%6s|\n", "foo", "bar")
	pf("width5 (|%%-6s|%%-6s|): |%-6s|%-6s|\n", "foo", "bar")

	// Sprintf formats and returns a string, no printing to cmd line
	str := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(str)

	// Formats and prints to an io var other than os.Stdout
	// Can be used for any io.Writer. For example files or stderr
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}