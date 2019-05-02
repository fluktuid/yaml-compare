package node

func (n *Node) Compare(other *Node) *Node {
	return n.getDifference(other)
}

func (n *Node) getDifference(other *Node) *Node {
	if n.same(other) {
		return nil
	}
	difference := n.deepCopy()

	thisSpare := n.children
	otherSpare := other.children

	/*
	   if (same(other))
	   return null

	   val y = clone()

	   val thisSpare = mutableListOf(* children.toTypedArray())
	   val otherSpare = mutableListOf(* other.children.toTypedArray())

	   val joined = thisSpare.map {
	   val sameFromOther =
	   otherSpare.firstOrNull { o -> o.same(it) } ?: otherSpare.firstOrNull { o -> o.value == it.value }
	   otherSpare.remove(sameFromOther)
	   Pair(it, sameFromOther)
	   }.toMutableList()
	   joined.addAll(
	   otherSpare.map {
	   val sameFromThis =
	   thisSpare.firstOrNull { o -> o.same(it) } ?: thisSpare.firstOrNull { o -> o.value == it.value }
	   thisSpare.remove(sameFromThis)
	   Pair(it, sameFromThis)
	   }
	   )

	   val differentChilds = joined.mapNotNull {
	   when {
	   it.second != null -> {
	   it.first.getDifference(it.second!!)
	   }
	   else -> it.first
	   }
	   }

	   y.addChild(differentChilds)
	   return y
	*/
}

func (n *Node) same(other *Node) bool {
	if n.Value != other.Value {
		return false
	}
	if n.parent != other.parent {
		return false
	}
	if n.Indent != other.Indent {
		return false
	}
	if len(n.children) != len(other.children) {
		return false
	}

	for i := range n.children {
		if &n.children[i] != &other.children[i] {
			return false
		}
	}
	return true
}
