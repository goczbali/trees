package treap

import "math/rand"

type Tree struct {
	root *node
}

func (t *Tree) ToSlice() (s []int) {
	toSlice(t.root, &s)
	return
}

func (t *Tree) Contains(key int) bool {
	return contains(t.root, key)
}

func (t *Tree) Insert(key int) {
	if t.Contains(key) {
		return
	}

	priority := rand.Int()
	t.root = insert(t.root, key, priority)
}

func (t *Tree) Split(k int) (l, r Tree) {
	l.root, r.root = split(t.root, k)

	return
}

func Merge(l, r Tree) (t Tree) {
	t.root = merge(l.root, r.root)
	return
}
