package avl

type Tree struct {
	root *node
}

func (t *Tree) Print() {
	printNode(t.root)
}

func (t *Tree) ToSlice() []int {
	var s []int
	toSlice(t.root, &s)
	return s
}

func (t *Tree) Height() int {
	return t.root.height
}

func (t *Tree) Contains(value int) bool {
	return contains(t.root, value)
}

func (t *Tree) Insert(value int) {
	if t.Contains(value) {
		return
	}

	t.root = insert(t.root, value)
}

func (t *Tree) Delete(value int) {
	t.root = deleteNode(t.root, value)
}
