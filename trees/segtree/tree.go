package segtree

type Tree struct {
	left       int
	right      int
	leftChild  *Tree
	rightChild *Tree

	value int
	delta int
}

func buildTree(values []int, left, right int) *Tree {
	if left == right {
		return &Tree{
			left:  left,
			right: right,
			value: values[left-1],
		}
	}

	mid := (left + right) / 2

	t := &Tree{
		left:       left,
		right:      right,
		leftChild:  buildTree(values, left, mid),
		rightChild: buildTree(values, mid+1, right),
	}

	t.update()

	return t
}

func Build(values []int) *Tree {
	return buildTree(values, 1, len(values))
}

func (t *Tree) update() {
	t.value = t.leftChild.Sum() + t.rightChild.Sum()
}

func (t *Tree) propagate() {
	t.leftChild.delta += t.delta
	t.rightChild.delta += t.delta

	t.delta = 0

	t.update()
}

func (t *Tree) Sum() int {
	if t == nil {
		return 0
	}

	return t.value + (t.left-t.right+1)*t.delta
}

func (t *Tree) Increment(a, b, v int) {
	if a <= t.left && b >= t.right {
		t.delta += v
	} else if b < t.left || a > t.right {
		return
	} else {
		t.propagate()
		t.leftChild.Increment(a, b, v)
		t.rightChild.Increment(a, b, v)
		t.update()
	}
}

func (t *Tree) Query(a, b int) int {
	if a <= t.left && b >= t.right {
		return t.Sum()
	} else if b < t.left || a > t.right {
		return 0
	} else {
		t.propagate()
		return t.leftChild.Query(a, b) + t.rightChild.Query(a, b)
	}
}

func (t *Tree) At(i int) int {
	return t.Query(i, i)
}
