// tree_test
package tree

import (
	"testing"
)

func TestInterator(t *testing.T) {
	tr := NewTree()
	root := tr.AddRoot(1)

	node2 := root.AddChild(2)
	node3 := root.AddChild(3)

	node2.AddChild(4)
	node2.AddChild(5)
	node6 := node2.AddChild(6)

	node7 := node3.AddChild(7)
	node3.AddChild(8)

	node6.AddChild(9)

	node7.AddChild(10)

	tr.Interator(func(v interface{}) {
		t.Log(v)
	})
}

func TestInteratorDeepRecursive(t *testing.T) {
	tr := NewTree()
	root := tr.AddRoot(1)

	node2 := root.AddChild(2)
	node3 := root.AddChild(3)

	node2.AddChild(4)
	node2.AddChild(5)
	node6 := node2.AddChild(6)

	node7 := node3.AddChild(7)
	node3.AddChild(8)

	node6.AddChild(9)

	node7.AddChild(10)

	tr.interatorDeepRecursive(func(v interface{}) {
		t.Log(v)
	})
}