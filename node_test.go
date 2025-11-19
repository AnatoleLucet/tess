package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	t.Run("creates a new node", func(t *testing.T) {
		node := NewNode()
		assert.NotNil(t, node)
		assert.Equal(t, NodeTypeDefault, node.GetNodeType())
		node.Free()
	})

	t.Run("creates a new node with styles", func(t *testing.T) {
		style := &Style{
			Width:  Point(100),
			Height: Point(200),
		}
		node := NewNode(style)
		assert.NotNil(t, node)
		assert.Equal(t, Point(100), node.GetWidth())
		assert.Equal(t, Point(200), node.GetHeight())
		node.Free()
	})
}

func TestNodeClone(t *testing.T) {
	t.Run("clones a node", func(t *testing.T) {
		original := NewNode()
		original.SetWidth(Point(100))

		clone := original.Clone()
		assert.NotNil(t, clone)
		assert.Equal(t, Point(100), clone.GetWidth())

		original.Free()
		clone.Free()
	})
}

func TestNodeFree(t *testing.T) {
	t.Run("frees a node", func(t *testing.T) {
		node := NewNode()
		node.Free()
		// Calling Free again should not panic
		node.Free()
	})
}

func TestNodeFreeRecursive(t *testing.T) {
	t.Run("frees a node and its children", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		parent.AddChild(child1)
		parent.AddChild(child2)

		parent.FreeRecursive()
		// Calling FreeRecursive again should not panic
		parent.FreeRecursive()
	})
}

func TestNodeReset(t *testing.T) {
	t.Run("resets a node", func(t *testing.T) {
		node := NewNode()
		node.SetWidth(Point(100))
		node.SetHeight(Point(200))

		node.Reset()

		// After reset, width and height should be auto (Yoga's default)
		assert.Equal(t, Auto(), node.GetWidth())
		assert.Equal(t, Auto(), node.GetHeight())

		node.Free()
	})
}

func TestNodeType(t *testing.T) {
	t.Run("gets and sets node type", func(t *testing.T) {
		node := NewNode()
		assert.Equal(t, NodeTypeDefault, node.GetNodeType())

		node.SetNodeType(NodeTypeText)
		assert.Equal(t, NodeTypeText, node.GetNodeType())

		node.Free()
	})
}

func TestNodeChildren(t *testing.T) {
	t.Run("gets child count", func(t *testing.T) {
		parent := NewNode()
		assert.Equal(t, 0, parent.GetChildCount())

		child := NewNode()
		parent.AddChild(child)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("adds child", func(t *testing.T) {
		parent := NewNode()
		child := NewNode()

		parent.AddChild(child)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("gets child by index", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		child1.SetWidth(Point(100))
		child2.SetWidth(Point(200))

		parent.AddChild(child1)
		parent.AddChild(child2)

		retrievedChild1 := parent.GetChild(0)
		retrievedChild2 := parent.GetChild(1)

		assert.NotNil(t, retrievedChild1)
		assert.NotNil(t, retrievedChild2)
		assert.Equal(t, Point(100), retrievedChild1.GetWidth())
		assert.Equal(t, Point(200), retrievedChild2.GetWidth())

		parent.FreeRecursive()
	})

	t.Run("returns nil for invalid child index", func(t *testing.T) {
		parent := NewNode()
		child := parent.GetChild(10)
		assert.Nil(t, child)
		parent.Free()
	})

	t.Run("sets children", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		parent.SetChildren([]*Node{child1, child2})
		assert.Equal(t, 2, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("sets empty children array", func(t *testing.T) {
		parent := NewNode()
		child := NewNode()
		parent.AddChild(child)

		assert.Equal(t, 1, parent.GetChildCount())

		parent.SetChildren([]*Node{})
		assert.Equal(t, 0, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("inserts child at index", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()
		child3 := NewNode()

		child1.SetWidth(Point(100))
		child2.SetWidth(Point(200))
		child3.SetWidth(Point(300))

		parent.AddChild(child1)
		parent.AddChild(child3)
		parent.InsertChild(child2, 1)

		assert.Equal(t, 3, parent.GetChildCount())
		assert.Equal(t, Point(100), parent.GetChild(0).GetWidth())
		assert.Equal(t, Point(200), parent.GetChild(1).GetWidth())
		assert.Equal(t, Point(300), parent.GetChild(2).GetWidth())

		parent.FreeRecursive()
	})

	t.Run("swaps child at index", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		child1.SetWidth(Point(100))
		child2.SetWidth(Point(200))

		parent.AddChild(child1)
		parent.SwapChild(child2, 0)

		assert.Equal(t, 1, parent.GetChildCount())
		assert.Equal(t, Point(200), parent.GetChild(0).GetWidth())

		parent.FreeRecursive()
	})

	t.Run("removes child", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		parent.AddChild(child1)
		parent.AddChild(child2)
		assert.Equal(t, 2, parent.GetChildCount())

		parent.RemoveChild(child1)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
		child1.Free()
	})

	t.Run("removes all children", func(t *testing.T) {
		parent := NewNode()
		child1 := NewNode()
		child2 := NewNode()

		parent.AddChild(child1)
		parent.AddChild(child2)
		assert.Equal(t, 2, parent.GetChildCount())

		parent.RemoveAllChildren()
		assert.Equal(t, 0, parent.GetChildCount())

		parent.Free()
		child1.Free()
		child2.Free()
	})
}

func TestNodeParent(t *testing.T) {
	t.Run("gets parent node", func(t *testing.T) {
		parent := NewNode()
		child := NewNode()

		assert.Nil(t, child.GetParent())

		parent.AddChild(child)
		retrievedParent := child.GetParent()
		assert.NotNil(t, retrievedParent)

		parent.FreeRecursive()
	})
}

func TestNodeConfig(t *testing.T) {
	t.Run("has default config set", func(t *testing.T) {
		node := NewNode()
		config := node.GetConfig()
		assert.NotNil(t, config)
		node.Free()
	})
}
