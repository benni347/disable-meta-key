package main

import (
	"bufio"
	"fmt"
	"os"
	Read "read/src/lib"
)

func main() {
	path := "testdata/read/1.txt"
	file, err := Read.ReadFile(path)
	if err != nil {
		return
	}
	var lineCont []byte
	fmt.Printf(string(rune(len(file))) + "\n")
	for i := 0; i < len(file); i++ {

		if string(file[i]) == "\n" {
			fmt.Printf("{%s}, i={%d} \n", string(lineCont), i) // FIXME: Currently this prints the nth char and not the nth word.
			var lineCont []byte                                // TODO: Implement this so I can set the lineCont var to a nil value, and call it on the next line again.
		} else {
			lineCont := append(lineCont, file[i]) // FIXME: This should be callable from inside the if operation so I can print the lineCont
		}

	}
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		_, err2 := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if err2 != nil {
			return
		}
	}
}
