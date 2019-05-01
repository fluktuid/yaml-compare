package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"yaml-compare/node"
)

func main() {
	//	wordPtr := flag.String("word", "foo", "a string")
	//	numbPtr := flag.Int("numb", 42, "an int")
	//	boolPtr := flag.Bool("fork", false, "a bool")

	//	var svar string
	//	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	//	fmt.Println("word:", *wordPtr)
	//	fmt.Println("numb:", *numbPtr)
	//	fmt.Println("fork:", *boolPtr)
	//	fmt.Println("svar:", svar)

	arguments := flag.Args()

	fmt.Println(arguments)

	for _, file := range arguments {
		lines, _ := readFileWithReadLine(file)
		//		fmt.Println(lines)
		//		fmt.Println("===")
		n := toNode(lines)
		n.ToString()
	}
}

func toNode(lines []string) *node.Node {
	root := node.Node{Value: "FILE", Indent: -1}
	lastLineNode := &root
	for _, line := range lines {
		var nP = node.NewNode(line)

		parent := lastLineNode.GetIndentParent(nP)
		parent.AddChildren(nP)

		lastLineNode = nP
	}
	return &root
}

func getIndent(line string) int {
	r, _ := regexp.Compile("^\\s*")
	return len(r.FindString(line))
}

func readFileWithReadLine(fn string) ([]string, error) {
	var rows []string
	fmt.Println("readFileWithReadLine")

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
