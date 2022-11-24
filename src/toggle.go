package main

import (
	"fmt"
	Read "read/src/lib"
)

func main() {
	path := "./testdata/read/1.txt"
	file, err := Read.ReadFile(path)
	if err != nil {
		fmt.Printf("The error of read is: %s\n", err)
	}
	var lineCont []byte
	fmt.Printf(string(rune(len(file))) + "\n")
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			fmt.Printf("{%s}, i={%d} \n", string(lineCont), i) // This should print the line till it reaches the \n char.
			lineCont = nil
		} else {
			lineCont = append(lineCont, file[i])
		}
	}
}
