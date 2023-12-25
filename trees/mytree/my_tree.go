package mytree

type MyTree struct {
	root *Node
	size int
}

type Node struct {
	value interface{}
	left  *Node
}

func (t *MyTree) Insert(value interface{}) {
	newNode := &Node{value: value}

	if t.root == nil {
		t.root = newNode
	} else {
		t.insertRecursive(t.root, newNode)
	}

	t.size++
}

func (t *MyTree) insertRecursive(currentNode, newNode *Node) {

	if currentNode.left == nil {
		currentNode.left = newNode
	} else {
		t.insertRecursive(currentNode.left, newNode)
	}
}

// Iterator
type TreeIterator struct {
	current *Node
	stack   []*Node
}

func (t *MyTree) Iterator() *TreeIterator {
	return &TreeIterator{current: t.root, stack: []*Node{}}
}

func (it *TreeIterator) Next() bool {
	if it.current != nil && it.stack != nil {

		if len(it.stack) == 0 {
			it.stack = append(it.stack[:0], it.current)

			return true
		}

		it.current = it.current.left
		if it.current != nil {
			it.stack = append(it.stack[:0], it.current)
		}

		return true
	}

	return false
}

func (it *TreeIterator) Value() interface{} {
	if it.current == nil {
		return nil
	}
	return it.current.value
}
