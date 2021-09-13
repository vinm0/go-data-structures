package main

type tree struct {
	root   *treeNode
	height int
	size   int
}

type treeNode struct {
	value  int
	left   *treeNode
	right  *treeNode
	parent *treeNode
}

func Tree() *tree {
	return &tree{}
}

func (t *tree) Size() int {
	return t.size
}

func (t *tree) Height() int {
	return t.height
}

func (t *tree) Root() *treeNode {
	return t.root
}

func (n *treeNode) Parent() *treeNode {
	return n.parent
}

func (n *treeNode) ChildL() *treeNode {
	return n.left
}

func (n *treeNode) ChildR() *treeNode {
	return n.right
}

func (n *treeNode) HasLeft() bool {
	return n.left != nil
}

func (n *treeNode) HasRight() bool {
	return n.right != nil
}

func (n *treeNode) IsLeaf() bool {
	return !(n.HasLeft() || n.HasRight())
}

func (t *tree) Add(val int) (ok bool) {
	n := &treeNode{value: val}

	if t.root == nil {
		t.root = n
		return true
	}

	ok = appendNode(t.root, n)

	return ok
}

func (t *tree) Remove(val int) (ok bool) {

	return removeValue(t.root, val)
}

// TODO: refactor as a generic function to
// ? Maybe add valType string field to tree, switch methods by type
func appendNode(curr *treeNode, n *treeNode) (ok bool) {

	if n.value < curr.value {
		if !curr.HasLeft() {
			curr.left = n
			n.parent = curr
			return true
		}

		return appendNode(curr.left, n)
	}

	if !curr.HasRight() {
		curr.right = n
		return true
	}

	return appendNode(curr.right, n)
}

func removeValue(curr *treeNode, val int) (ok bool) {
	// TODO
	return false
}
