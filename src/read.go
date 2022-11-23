package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	path := "test.txt"
	content, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
	}
	fmt.Printf("The content of '%s' : \n%s\n", path, content)
}

func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pullPath := filepath.Join(parentPath, path)
	file, err := os.Open(pullPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	return read(file)
}

func read(fdR io.Reader) ([]byte, error) {
	br := bufio.NewReader(fdR)
	var buf bytes.Buffer

	for {
		ba, isPrefix, err := br.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}

	}
	return buf.Bytes(), nil
}
