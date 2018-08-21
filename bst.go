package bst

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ERR_NOT_FOUND = errors.New("elem not found")
)

func NewBST() *BST {
	return &BST{}
}

type BST struct {
	mu   sync.RWMutex
	Root *Node
}

func (b *BST) Insert(val int) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	item := newNode(val)
	return b.insert(item)
}

func (b *BST) insert(item *Node) error {

	var y *Node
	x := b.Root

	for x != nil {
		y = x
		if item.Value < x.Value {
			// insert value into the left node
			x = x.Left
		} else if item.Value > x.Value {
			// insert value into the left node
			x = x.Right
		} else {
			// value exists
			return nil
		}
	}
	item.Parent = y
	if y == nil {
		b.Root = item
		return nil
	} else if item.Value < y.Value {
		y.Left = item
	} else {
		y.Right = item
	}
	return nil
}

func (b *BST) Search(val int) *Node {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.Root == nil {
		return nil
	}
	return b.Root.search(val)
}

func (b *BST) Delete(val int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	removeNode(b.Root, val)
}

// Inorder traversal
func (b *BST) Traverse() {
	b.mu.RLock()
	defer b.mu.RUnlock()
	fn := func(n *Node) {
		fmt.Println(n.Value)
	}
	b.Root.traverse(fn)
}

func (b *BST) Minimum() int {
	if b.Root == nil {
		return 0
	}
	return b.Root.minimum().Value
}

func (b *BST) Maximum() int {
	if b.Root == nil {
		return 0
	}
	return b.Root.maximum().Value
}

func (b *BST) Predecessor(val int) (int, error) {
	item := b.Search(val)
	if item == nil {
		return 0, ERR_NOT_FOUND
	}
	n := item.predecessor()
	if n == nil {
		return 0, ERR_NOT_FOUND
	}
	return n.Value, nil
}

func (b *BST) Successor(val int) (int, error) {
	item := b.Search(val)
	if item == nil {
		return 0, ERR_NOT_FOUND
	}
	n := item.successor()
	if n == nil {
		return 0, ERR_NOT_FOUND
	}
	return n.Value, nil
}

func newNode(val int) *Node {
	return &Node{Value: val}
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func (n *Node) search(val int) *Node {
	switch {
	case val == n.Value:
		return n
	case val < n.Value:
		if n.Left == nil {
			return nil
		}
		return n.Left.search(val)
	case val > n.Value:
		if n.Right == nil {
			return nil
		}
		return n.Right.search(val)
	}
	return nil

}

func (n *Node) successor() *Node {
	return findSuccessor(n)
}

func (n *Node) predecessor() *Node {
	return findPredecessor(n)
}

func (n *Node) minimum() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *Node) maximum() *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *Node) traverse(fn func(*Node)) {
	if n == nil {
		return
	}
	n.Left.traverse(fn)
	fn(n)
	n.Right.traverse(fn)
}

func swapNode(a, b *Node) {
	*a, *b = *b, *a
}

func findSuccessor(x *Node) *Node {
	if x.Right != nil {
		return x.Right.minimum()
	}
	y := x.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

func findPredecessor(x *Node) *Node {
	if x.Left != nil {
		return x.Left.maximum()
	}
	y := x.Parent
	for y != nil && x == y.Left {
		x = y
		y = y.Parent
	}
	return y
}

func removeNode(a *Node, val int) *Node {
	if a == nil {
		return nil
	}
	if val < a.Value {
		a.Left = removeNode(a.Left, val)
		return a
	}
	if val > a.Value {
		a.Right = removeNode(a.Right, val)
		return a
	}
	// Remove leaf node
	if a.Left == nil && a.Right == nil {
		a = nil
		return nil
	}
	// Remove half leaf node
	if a.Left == nil {
		swapNode(a, a.Right)
		return a
	}
	// Remove half leaf node
	if a.Right == nil {
		swapNode(a, a.Left)
		return a
	}
	// Remove half parent node
	tempNode := findSuccessor(a.Left)
	a.Value = tempNode.Value
	a.Left = removeNode(a.Left, tempNode.Value)
	return a
}
