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

func (b *BST) Insert(key int, value string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	item := newNode(key, value)
	return b.insert(item)
}

func (b *BST) insert(item *Node) error {

	var y *Node
	x := b.Root

	for x != nil {
		y = x
		if item.Key < x.Key {
			// insert key into the left node
			x = x.Left
		} else if item.Key > x.Key {
			// insert key into the left node
			x = x.Right
		} else {
			// key exists
			return nil
		}
	}
	item.Parent = y
	if y == nil {
		b.Root = item
		return nil
	} else if item.Key < y.Key {
		y.Left = item
	} else {
		y.Right = item
	}
	return nil
}

func (b *BST) Search(key int) *Node {
	b.mu.RLock()
	defer b.mu.RUnlock()
	x := b.Root

	if x == nil {
		return nil
	}

	for x != nil {
		switch {
		case key == x.Key:
			return x
		case key < x.Key:
			x = x.Left
		case key > x.Key:
			x = x.Right
		}
	}

	return nil
}

func (b *BST) Delete(key int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	removeNode(b.Root, key)
}

// Inorder traversal
func (b *BST) Traverse() {
	b.mu.RLock()
	defer b.mu.RUnlock()
	fn := func(n *Node) {
		fmt.Println(n.Key)
	}
	b.Root.traverse(fn)
}

func (b *BST) Minimum() int {
	if b.Root == nil {
		return 0
	}
	return b.Root.minimum().Key
}

func (b *BST) Maximum() int {
	if b.Root == nil {
		return 0
	}
	return b.Root.maximum().Key
}

func (b *BST) Predecessor(key int) (int, error) {
	item := b.Search(key)
	if item == nil {
		return 0, ERR_NOT_FOUND
	}
	n := item.predecessor()
	if n == nil {
		return 0, ERR_NOT_FOUND
	}
	return n.Key, nil
}

func (b *BST) Successor(key int) (int, error) {
	item := b.Search(key)
	if item == nil {
		return 0, ERR_NOT_FOUND
	}
	n := item.successor()
	if n == nil {
		return 0, ERR_NOT_FOUND
	}
	return n.Key, nil
}

func newNode(key int, val string) *Node {
	return &Node{Key: key, Value: val}
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
	Value  string
}

func (n *Node) successor() *Node {
	if n.Right != nil {
		return n.Right.minimum()
	}
	y := n.Parent
	for y != nil && n == y.Right {
		n = y
		y = y.Parent
	}
	return y
}

func (n *Node) predecessor() *Node {
	if n.Left != nil {
		return n.Left.maximum()
	}
	y := n.Parent
	for y != nil && n == y.Left {
		n = y
		y = y.Parent
	}
	return y
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

func removeNode(a *Node, key int) *Node {
	if a == nil {
		return nil
	}
	if key < a.Key {
		a.Left = removeNode(a.Left, key)
		return a
	}
	if key > a.Key {
		a.Right = removeNode(a.Right, key)
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
	tempNode := a.Left.successor()
	a.Key = tempNode.Key
	a.Left = removeNode(a.Left, tempNode.Key)
	return a
}
