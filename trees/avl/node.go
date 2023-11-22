package avl

import "fmt"

type node struct {
	value  int
	height int
	bf     int

	left  *node
	right *node
}

func contains(node *node, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	}

	if value < node.value {
		return contains(node.left, value)
	} else {
		return contains(node.right, value)
	}
}

func printNode(n *node) {
	if n == nil {
		return
	}

	printNode(n.left)
	fmt.Println(n.value)
	printNode(n.right)
}

func toSlice(n *node, s *[]int) {
	if n == nil {
		return
	}

	toSlice(n.left, s)
	*s = append(*s, n.value)
	toSlice(n.right, s)
}

func update(n *node) {
	lh := -1
	if n.left != nil {
		lh = n.left.height
	}

	rh := -1
	if n.right != nil {
		rh = n.right.height
	}

	h := lh
	if rh > h {
		h = rh
	}

	n.height = h + 1
	n.bf = rh - lh
}

func leftRotate(n *node) *node {
	r := n.right

	n.right = r.left
	r.left = n

	update(n)
	update(r)

	return r
}

func rightRotate(n *node) *node {
	l := n.left

	n.left = l.right
	l.right = n

	update(n)
	update(l)

	return l
}

func balance(n *node) *node {
	if n.bf == -2 {
		if n.left.bf > 0 {
			n.left = leftRotate(n.left)
		}

		n = rightRotate(n)
	} else if n.bf == 2 {
		if n.right.bf < 0 {
			n.right = rightRotate(n.right)
		}

		n = leftRotate(n)
	}

	return n
}

func findMax(n *node) int {
	if n.right == nil {
		return n.value
	}

	return findMax(n.right)
}

func findMin(n *node) int {
	if n.left == nil {
		return n.value
	}

	return findMin(n.left)
}

func insert(n *node, value int) *node {
	if n == nil {
		return &node{value: value}
	}

	if value < n.value {
		n.left = insert(n.left, value)
	} else {
		n.right = insert(n.right, value)
	}

	update(n)
	return balance(n)
}

func deleteNode(n *node, value int) *node {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = deleteNode(n.left, value)
	} else if value > n.value {
		n.right = deleteNode(n.right, value)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			if n.left.height > n.right.height {
				v := findMax(n.left)
				n.value = v

				n.left = deleteNode(n.left, v)
			} else {
				v := findMin(n.right)
				n.value = v

				n.right = deleteNode(n.right, v)
			}
		}
	}

	update(n)
	return balance(n)
}
