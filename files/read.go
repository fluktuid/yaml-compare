package files

import (
	"../helper"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Read(fn ...string) ([]*File, error) {
	var paths []string
	var files []*File
	for _, v := range fn {
		p, err := readAllFilesInDirectory(v)
		if err != nil {
			return nil, err
		}
		paths = append(paths, p...)
	}
	for _, v := range paths {
		file, err := readFileWithReadLine(v)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func readFileWithReadLine(fn string) (*File, error) {
	var rows []string

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()

		// Process the line here.
		//		fmt.Println(line)
		rows = append(rows, line)
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return &File{Name: fn, Lines: rows}, nil
}

func readAllFilesInDirectory(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if helper.YmlFile(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files, err
}
