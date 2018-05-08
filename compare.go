package fsdiff

func Compare(original, changed *Node) *Diff {
	return compareNode(original, changed)
}

func compareNode(original, changed *Node) *Diff {
	var rootDiff Diff
	rootDiff.Path = original.Path
	rootDiff.Children = []*Diff{}

	if original.Hash == changed.Hash && original.IsDirectory == changed.IsDirectory {
		rootDiff.DiffType = DiffTypeNone
		return &rootDiff
	}

	rootDiff.DiffType = DiffTypeChanged

	if original.IsDirectory && changed.IsDirectory {
		// search for changes to original children
		for _, origChild := range original.Children {
			changeChild, ok := searchChildrenForPath(changed, origChild.Path)

			// child not found in changed tree
			if !ok {
				rootDiff.Children = append(rootDiff.Children, &Diff{
					Path:     origChild.Path,
					DiffType: DiffTypeRemoved,
					Children: []*Diff{},
				})
				continue
			}

			// compare the two nodes
			rootDiff.Children = append(rootDiff.Children, compareNode(origChild, changeChild))
		}

		// search for additions to children
		for _, changeChild := range changed.Children {
			_, ok := searchChildrenForPath(original, changeChild.Path)

			// child not found in changed tree
			if !ok {
				rootDiff.Children = append(rootDiff.Children, addAll(changeChild))
				continue
			}
		}
	}

	return &rootDiff
}

func addAll(rootNode *Node) *Diff {
	child := &Diff{
		Path:     rootNode.Path,
		DiffType: DiffTypeAdded,
		Children: []*Diff{},
	}

	for _, ch := range rootNode.Children {
		child.Children = append(child.Children, addAll(ch))
	}

	return child
}

func searchChildrenForPath(parent *Node, path string) (*Node, bool) {
	for _, ch := range parent.Children {
		if ch.Path == path {
			return ch, true
		}
	}
	return nil, false
}
