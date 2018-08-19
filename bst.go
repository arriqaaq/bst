package bst

import (
	"fmt"
	"sync"
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
	if b.Root == nil {
		b.Root = newNode(val)
		return nil
	}
	return b.Root.insert(val)
}

func (b *BST) Find(val int) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.Root == nil {
		return false
	}
	return b.Root.find(val)
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
	b.Root.traverse()
}

func newNode(val int) *Node {
	return &Node{Value: val}
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) insert(val int) error {
	switch {
	case val == n.Value:
		return nil
	case val < n.Value:
		if n.Left == nil {
			n.Left = newNode(val)
			return nil
		}
		return n.Left.insert(val)
	case val > n.Value:
		if n.Right == nil {
			n.Right = newNode(val)
			return nil
		}
		return n.Right.insert(val)
	}
	return nil
}

func (n *Node) find(val int) bool {
	switch {
	case val == n.Value:
		return true
	case val < n.Value:
		if n.Left == nil {
			return false
		}
		return n.Left.find(val)
	case val > n.Value:
		if n.Right == nil {
			return false
		}
		return n.Right.find(val)
	}
	return false

}

func (n *Node) traverse() {
	if n == nil {
		return
	}
	n.Left.traverse()
	fmt.Println(n.Value)
	n.Right.traverse()
}

func swapNode(a, b *Node) {
	*a, *b = *b, *a
}

func findSuccessor(a *Node) *Node {
	if a.Right == nil {
		return a
	}
	return findSuccessor(a.Right)
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
