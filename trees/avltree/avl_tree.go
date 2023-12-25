package avltree

import (
	"github.com/emirpasic/gods/trees/avltree"
	"tree-complete/trees/mytree"
)

func ConvertAVLToMyTree(avlTree *avltree.Tree) *mytree.MyTree {
	myTree := &mytree.MyTree{}
	avlIterator := avlTree.Iterator()
	for avlIterator.Next() {
		value := avlIterator.Value()
		myTree.Insert(value)
	}
	return myTree
}
