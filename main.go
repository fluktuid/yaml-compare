package main

import (
	"./files"
	"./node"
	"errors"
	"fmt"
	. "github.com/logrusorgru/aurora"
	flag "github.com/spf13/pflag"
	"os"
)

func main() {
	config, arguments := getFlags()

	if len(arguments) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", errors.New("please specify one or two Files"))
		flag.Usage()
		os.Exit(1)
		return
	} else if len(arguments) > 2 {
		fmt.Println("warining: I got more than two files. I will compare them step by step.")
	}

	var roots []*node.Node
	fls, _ := files.Read(arguments...)
	for _, file := range fls {
		n := node.ToNode(file.Lines, config)
		if config.PrintFiles || len(fls) == 1 {
			fmt.Printf("File: %v\n", file.Name)
			fmt.Println(file.Lines)
			n.PrintBy(0, config)
			fmt.Println("==========")
			fmt.Println()
		}
		roots = append(roots, n)
	}
	if len(fls) > 1 {
		for i := 1; i < len(roots); i++ {
			if i > 1 {
				fmt.Println()
			}
			if config.ColorLess {
				fmt.Printf("%v => %v\n", fls[i-1].Name, fls[i].Name)
			} else {
				fmt.Println(Underline(fls[i-1].Name + " => " + fls[i].Name))
			}
			difference := roots[i].Compare(roots[i-1], config)
			difference.PrintBy(1, config)
		}
	}
}

// Returns all user-given flags
// The return values are a config file and a string array containing all arguments
func getFlags() (node.Config, []string) {
	resolveAnchors := flag.BoolP("resolve-anchors", "R", true, "Resolve Yaml Anchors, e.g. '&id001'")
	printLineTypes := flag.BoolP("print-line-types", "L", false, "Print the Line Types, e.g. 'ListItem'")
	printFiles := flag.BoolP("print", "p", false, "Print files after anchor resolving")
	printComplete := flag.BoolP("print-complete", "c", false, "Print the complete diff file after comparing")
	colorLess := flag.BoolP("white", "w", false, "Print without ANSI color")
	bewareAnchors := flag.BoolP("beware-anchors", "A", false, "Beware anchor names while resolving (not implemented)")
	bewarePointer := flag.BoolP("beware-pointer", "P", false, "Beware pointer names while resolving (not implemented)")
	fullQualifierName := flag.BoolP("full-qualifier-name", "f", false, "use full-qualifier names, e.g. 'step[0].instrument' (alpha)")
	version := flag.BoolP("version", "V", false, "Print Version and exit")
	flag.Parse()
	if *version {
		fmt.Println(Underline("Yaml Compare"))
		fmt.Println("Version: 0.1.0-SNAPSHOT")
		fmt.Println("Author: Lukas f. Paluch")
		os.Exit(0)
	}

	config := node.Config{
		ResolveAnchors:    *resolveAnchors,
		PrintLineTypes:    *printLineTypes,
		BewareAnchors:     *bewareAnchors,
		BewarePointer:     *bewarePointer,
		FullQualifierName: *fullQualifierName,
		PrintComplete:     *printComplete,
		PrintFiles:        *printFiles,
		ColorLess:         *colorLess,
	}
	return config, flag.Args()
}
