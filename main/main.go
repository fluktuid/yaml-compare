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

	var roots []*node.Node
	for _, file := range arguments {
		lines, _ := files.ReadFileWithReadLine(file)
		//		fmt.Println(lines)
		//		fmt.Println("===")
		n := toNode(lines)
		n.Print()
		roots = append(roots, n)
	}
	difference := roots[0].Compare(roots[1])
	difference.Print()
}

func toNode(lines []string) *node.Node {
	rootP := node.NewFile()
	blockP, _ := node.New("---")
	rootP.AddChildren(blockP)
	lastLineNode := blockP
	for _, line := range lines {
		nP, _ := node.New(line)

		if nP != nil {
			parent := lastLineNode.GetIndentParent(nP)
			parent.AddChildren(nP)
			lastLineNode = nP
		}
	}
	rootP.Clean()
	// TODO: add possibility to turn resolve off via flag
	rootP.ResolveAnchors()
	return rootP
}
