package bst

import (
	"fmt"
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
		node := bst.Search(i)
		assertEqual(t, i, node.Value)
	}
}

func TestNode_Delete(t *testing.T) {
	bst := NewBST()
	bst.Insert(1)
	bst.Delete(1)
	bst.Traverse()
}

func Test_Minimum(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(40)
	fmt.Println("before: ", bst.Root.Left)
	y := bst.Minimum()
	fmt.Println("after: ", bst.Root.Left)
	assertEqual(t, y, 20)
}

func Test_Maximum(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(40)
	y := bst.Maximum()
	assertEqual(t, y, 50)
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
