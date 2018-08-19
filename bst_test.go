package bst

import (
	// "fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("not equal, a=%s,b=%s", a, b)
	}
}

func TestNode_Insert(t *testing.T) {
	bst := NewBST()
	for i := 1; i < 10; i++ {
		bst.Insert(i)
	}
	for i := 1; i < 10; i++ {
		isPresent := bst.Find(i)
		assertEqual(t, isPresent, true)
	}
}

func TestNode_Delete(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Delete(1)
	bst.Traverse()
}

func TestNode_Delete_BST_C1(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)
	bst.Delete(30)
	bst.Traverse()
}

func TestNode_Delete_BST_C2(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(70)
	bst.Insert(60)
	bst.Insert(80)
	bst.Delete(50)
	bst.Traverse()
}
