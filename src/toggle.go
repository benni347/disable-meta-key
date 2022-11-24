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

	fmt.Printf(string(rune(len(file))) + "\n")
	fmt.Printf("Hello, {%s}\n", string(file[4])) // FIXME: Currently it returns the nth character and not the nth line
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
