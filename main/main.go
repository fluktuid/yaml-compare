package main

import (
	"flag"
	"fmt"
	"yaml-compare/files"
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
		lines, _ := files.ReadFileWithReadLine(file)
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
