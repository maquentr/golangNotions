package main

/*A line filter is a common type of program that reads input on stdin, processes it, and then prints some derived result to stdout. 
grep and sed are common line filters.*/

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	/*Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token; 
	which is the next line in the default scanner.*/
	scanner := bufio.NewScanner(os.Stdin)

	/*Text returns the current token, here the next line, from the input.*/
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text()) //Write out the uppercased line.
		fmt.Println(ucl)
	}

	/*Check for errors during Scan. End of file is expected and not reported by Scan as an error.*/
	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }


}