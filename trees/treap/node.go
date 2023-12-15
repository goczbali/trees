package treap

type node struct {
	key      int
	priority int
	left     *node
	right    *node
}

func toSlice(n *node, s *[]int) {
	if n == nil {
		return
	}

	toSlice(n.left, s)
	*s = append(*s, n.key)
	toSlice(n.right, s)
}

func contains(n *node, k int) bool {
	if n == nil {
		return false
	} else if n.key == k {
		return true
	} else {
		if k < n.key {
			return contains(n.left, k)
		} else {
			return contains(n.right, k)
		}
	}
}

func split(n *node, k int) (*node, *node) {
	if n == nil {
		return nil, nil
	}

	if n.key <= k {
		l, r := split(n.right, k)
		n.right = l
		return n, r
	} else {
		l, r := split(n.left, k)
		n.left = r
		return l, n
	}
}

func merge(l *node, r *node) *node {
	if l == nil {
		return r
	} else if r == nil {
		return l
	}

	if l.priority < r.priority {
		r.left = merge(l, r.left)
		return r
	} else {
		l.right = merge(l.right, r)
		return l
	}
}

func union(n1 *node, n2 *node) *node {
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}

	if n1.priority < n2.priority {
		l, r := split(n2, n1.key)
		n1.left = union(n1.left, l)
		n1.right = union(n1.right, r)

		return n1
	} else {
		l, r := split(n1, n2.key)
		n2.left = union(n1.left, l)
		n2.right = union(n2.right, r)

		return n2
	}
}

func insert(n *node, key, priority int) *node {
	if n == nil || priority < n.priority {
		l, r := split(n, key)

		return &node{
			key:      key,
			priority: priority,
			left:     l,
			right:    r,
		}
	}

	if key < n.key {
		n.left = insert(n.left, key, priority)
	} else {
		n.right = insert(n.right, key, priority)
	}

	return n
}
