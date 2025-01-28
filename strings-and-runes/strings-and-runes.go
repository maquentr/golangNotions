package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("len = ", len(s))
	
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	//To count how many runes are in a string, we can use the utf8 package. 
	//Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. 
	//Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	//A range loop handles strings specially and decodes each rune along with its offset in the string.
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
	if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

