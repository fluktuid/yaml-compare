package files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadFileWithReadLine(fn string) ([]string, error) {
	var rows []string

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return rows, err
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

	return rows, nil
}
