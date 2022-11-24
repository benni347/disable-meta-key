package main

import (
	"fmt"
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
			fmt.Printf("{%s}, i={%d} \n", string(lineCont), i) // This should print the line till it reaches the \n char.
			lineCont = nil
		} else {
			lineCont = append(lineCont, file[i])
		}
	}
}

/*func readInput() {
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
}*/
