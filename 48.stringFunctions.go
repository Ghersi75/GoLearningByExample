package main

import (
	"fmt"
	// alias string as s
	s "strings"
)

// alias println to a shorter name as well
// would have been great to know in the past
var p = fmt.Println

func StringFunctions() {
	p("Contains:	", s.Contains("test", "es"))
	p("Count: 		", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Index: 		", s.Index("test", "te"))
	p("Join:			", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat:		", s.Repeat("What a Save! ", 5))
	p("Replace:		", s.Replace("foo", "o", "0", -1))
	p("Replace:		", s.Replace("foo", "o", "0", 1))
	p("Split:			", s.Split("a-b-c-d-e", "-"))
	p("ToLower:		", s.ToLower("LOWER"))
	p("ToUpper:		", s.ToUpper("upper"))
}
