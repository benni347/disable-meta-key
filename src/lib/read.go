package lib

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

/*
	func main() {
		path := "test.txt"
		content, err := readFile(path)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
		}
		fmt.Printf("The content of '%s' : \n%s\n", path, content)
	}
*/
// This function takes the path parameter which is a string of the path were the file is located.
func ReadFile(path string) ([]byte, error) {

	var regexCwd = regexp.MustCompile(`(?m)\./.*\n`) // regexCwd checks if it is in the current working directory.
	var regexSys = regexp.MustCompile(`(?m)/home.*\n|/dev.*\n|/sys/.*\n`)
	// FIXME: At the moment this always returns the error.
	if regexCwd.Find([]byte(path)) != nil {
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
	} else if regexSys.Find([]byte(path)) != nil {
		file, err := os.Open(path)
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
	} else {
		var err = errors.New("the given path isn't valid to read from")
		return nil, err
	}
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
