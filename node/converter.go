package node

type Config struct {
	ResolveAnchors    bool
	PrintLineTypes    bool
	BewareAnchors     bool
	BewarePointer     bool
	printFull         bool
	FullQualifierName bool
	PrintComplete     bool
}

func ToNode(lines []string, config Config) *Node {
	rootP := NewFile()
	blockP, _ := New("---")
	rootP.AddChildren(blockP)
	lastLineNode := blockP
	for _, line := range lines {
		nP, _ := New(line)

		if nP != nil {
			parent := lastLineNode.GetIndentParent(nP)
			parent.AddChildren(nP)
			lastLineNode = nP
		}
	}
	rootP.Clean()
	if config.ResolveAnchors {
		rootP.resolveAnchors()
	}
	return rootP
}
