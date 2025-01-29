package main

import "os"

/*A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle. 
Here’s an example of panicking if we get an unexpected error when creating a new file.*/

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	/*When first panic in main fires, the program exits without reaching the rest of the code. 
	If you’d like to see the program try to create a temp file, comment the first panic out.*/
}