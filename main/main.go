package main

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"yaml-compare/files"
	"yaml-compare/node"
)

func main() {
	resolveAnchors := flag.BoolP("resolve-anchors", "R", false, "Resolve Yaml Anchors, e.g. '&id001'")

	// TODO: implement
	printLineTypes := flag.BoolP("print-line-types", "L", false, "Print the Line Types, e.g. 'ListItem'")

	// TODO: implement
	bewareAnchors := flag.BoolP("beware-anchors", "A", false, "Resolve Yaml Anchors, e.g. '&id001'")

	// TODO: implement
	bewarePointer := flag.BoolP("beware-pointer", "P", false, "Resolve Yaml Anchors, e.g. '&id001'")
	// TODO: implement
	fullQualifierName := flag.BoolP("full-qualifier-name", "f", false, "Resolve Yaml Anchors, e.g. '&id001'")
	flag.Parse()

	arguments := flag.Args()
	if len(arguments) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", errors.New("please set two Files"))
		os.Exit(1)
		return
	}
	config := node.Config{
		ResolveAnchors:    *resolveAnchors,
		PrintLineTypes:    *printLineTypes,
		BewareAnchors:     *bewareAnchors,
		BewarePointer:     *bewarePointer,
		FullQualifierName: *fullQualifierName,
	}

	var roots []*node.Node
	for _, file := range arguments {
		lines, _ := files.ReadFileWithReadLine(file)
		n := node.ToNode(lines, config)
		roots = append(roots, n)
	}
	difference := roots[0].Compare(roots[1])
	difference.PrintBy(1)
}
