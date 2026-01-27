package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	t.Run("creates a new node", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		assert.NotNil(t, node)
		assert.Equal(t, NodeTypeDefault, node.GetNodeType())
		node.Free()
	})

	t.Run("creates a new node with styles", func(t *testing.T) {
		style := &Style{
			Width:  Point(100),
			Height: Point(200),
		}
		node, err := NewNode(style)
		assert.NoError(t, err)
		assert.NotNil(t, node)
		assert.Equal(t, Point(100), node.GetWidth())
		assert.Equal(t, Point(200), node.GetHeight())
		node.Free()
	})
}

func TestNodeClone(t *testing.T) {
	t.Run("clones a node", func(t *testing.T) {
		original, err := NewNode()
		assert.NoError(t, err)
		original.SetWidth(Point(100))

		clone := original.Clone()
		assert.NotNil(t, clone)
		assert.Equal(t, Point(100), clone.GetWidth())

		original.Free()
		clone.Free()
	})

	t.Run("clones a node with children", func(t *testing.T) {
		original, err := NewNode()
		assert.NoError(t, err)

		originalChild, err := NewNode()
		assert.NoError(t, err)
		original.AppendChild(originalChild)

		clone := original.CloneRecursive()
		assert.NotNil(t, clone)

		child := clone.GetChild(0)
		assert.NotNil(t, child)

		child.SetWidth(Point(200))
		assert.Equal(t, Point(200), child.GetWidth(), "cloned child's width should be independent")
		assert.Equal(t, Auto(), originalChild.GetWidth(), "original child's width should remain unchanged")

		original.Free()
		clone.Free()
	})

	t.Run("clones a node with nested children", func(t *testing.T) {
		original, err := NewNode()
		assert.NoError(t, err)

		parentChild, err := NewNode()
		assert.NoError(t, err)
		original.AppendChild(parentChild)

		nestedChild, err := NewNode()
		assert.NoError(t, err)
		parentChild.AppendChild(nestedChild)

		clone := original.CloneRecursive()
		assert.NotNil(t, clone)

		clonedParentChild := clone.GetChild(0)
		assert.NotNil(t, clonedParentChild)

		clonedNestedChild := clonedParentChild.GetChild(0)
		assert.NotNil(t, clonedNestedChild)

		clonedNestedChild.SetHeight(Point(150))
		assert.Equal(t, Point(150), clonedNestedChild.GetHeight(), "cloned nested child's height should be independent")
		assert.Equal(t, Auto(), nestedChild.GetHeight(), "original nested child's height should remain unchanged")

		original.Free()
		clone.Free()
	})

	t.Run("clones node with measure func", func(t *testing.T) {
		measureCalled := 0
		measureFunc := func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			measureCalled++
			return Size{Width: 50, Height: 20}
		}

		parent, err := NewNode()
		assert.NoError(t, err)
		parent.SetWidth(Point(200))
		parent.SetHeight(Point(200))

		original, err := NewNode()
		assert.NoError(t, err)
		original.SetMeasureFunc(measureFunc)
		parent.AppendChild(original)

		clonedParent := parent.CloneRecursive()
		clone := clonedParent.GetChild(0)
		assert.NotNil(t, clone)
		assert.True(t, clone.HasMeasureFunc(), "clone should have measure function")

		clonedParent.ComputeLayout(Container{Width: 300, Height: 300})
		assert.Equal(t, 1, measureCalled, "measure should be called once for clone")

		parent.ComputeLayout(Container{Width: 300, Height: 300})
		assert.Equal(t, 2, measureCalled, "measure should be called once for original")

		original.UnsetMeasureFunc()
		assert.False(t, original.HasMeasureFunc(), "original should not have measure function")
		assert.True(t, clone.HasMeasureFunc(), "clone should still have measure function")

		clone.MarkDirty()
		clonedParent.ComputeLayout(Container{Width: 300, Height: 300})
		assert.Equal(t, 3, measureCalled, "measure should be called again for clone")

		parent.FreeRecursive()
		clonedParent.FreeRecursive()
	})
}

func TestNodeFree(t *testing.T) {
	t.Run("frees a node", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.Free()
		// Calling Free again should not panic
		node.Free()
	})
}

func TestNodeFreeRecursive(t *testing.T) {
	t.Run("frees a node and its children", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		assert.NoError(t, err)

		parent.AppendChild(child1)
		parent.AppendChild(child2)

		parent.FreeRecursive()
		// Calling FreeRecursive again should not panic
		parent.FreeRecursive()
	})
}

func TestNodeReset(t *testing.T) {
	t.Run("resets a node", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
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
		node, err := NewNode()
		assert.NoError(t, err)
		assert.Equal(t, NodeTypeDefault, node.GetNodeType())

		node.SetNodeType(NodeTypeText)
		assert.Equal(t, NodeTypeText, node.GetNodeType())

		node.Free()
	})
}

func TestNodeChildren(t *testing.T) {
	t.Run("gets child count", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		assert.Equal(t, 0, parent.GetChildCount())

		child, err := NewNode()
		assert.NoError(t, err)
		parent.AppendChild(child)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("adds child", func(t *testing.T) {
		parent, err := NewNode()
		child, err := NewNode()
		assert.NoError(t, err)

		parent.AppendChild(child)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("gets child by index", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		assert.NoError(t, err)

		child1.SetWidth(Point(100))
		child2.SetWidth(Point(200))

		parent.AppendChild(child1)
		parent.AppendChild(child2)

		retrievedChild1 := parent.GetChild(0)
		retrievedChild2 := parent.GetChild(1)

		assert.NotNil(t, retrievedChild1)
		assert.NotNil(t, retrievedChild2)
		assert.Equal(t, Point(100), retrievedChild1.GetWidth())
		assert.Equal(t, Point(200), retrievedChild2.GetWidth())

		parent.FreeRecursive()
	})

	t.Run("returns nil for invalid child index", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		child := parent.GetChild(10)
		assert.Nil(t, child)
		parent.Free()
	})

	t.Run("sets children", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		assert.NoError(t, err)

		parent.SetChildren([]*Node{child1, child2})
		assert.Equal(t, 2, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("sets empty children array", func(t *testing.T) {
		parent, err := NewNode()
		child, err := NewNode()
		assert.NoError(t, err)
		parent.AppendChild(child)

		assert.Equal(t, 1, parent.GetChildCount())

		parent.SetChildren([]*Node{})
		assert.Equal(t, 0, parent.GetChildCount())

		parent.FreeRecursive()
	})

	t.Run("inserts child at index", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		child3, err := NewNode()
		assert.NoError(t, err)

		child1.SetWidth(Point(100))
		child2.SetWidth(Point(200))
		child3.SetWidth(Point(300))

		parent.AppendChild(child1)
		parent.AppendChild(child3)
		parent.InsertChild(child2, 1)

		assert.Equal(t, 3, parent.GetChildCount())
		assert.Equal(t, Point(100), parent.GetChild(0).GetWidth())
		assert.Equal(t, Point(200), parent.GetChild(1).GetWidth())
		assert.Equal(t, Point(300), parent.GetChild(2).GetWidth())

		parent.FreeRecursive()
	})

	t.Run("removes child", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		assert.NoError(t, err)

		parent.AppendChild(child1)
		parent.AppendChild(child2)
		assert.Equal(t, 2, parent.GetChildCount())

		parent.RemoveChild(child1)
		assert.Equal(t, 1, parent.GetChildCount())

		parent.FreeRecursive()
		child1.Free()
	})

	t.Run("removes all children", func(t *testing.T) {
		parent, err := NewNode()
		child1, err := NewNode()
		child2, err := NewNode()
		assert.NoError(t, err)

		parent.AppendChild(child1)
		parent.AppendChild(child2)
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
		parent, err := NewNode()
		child, err := NewNode()
		assert.NoError(t, err)

		assert.Nil(t, child.GetParent())

		parent.AppendChild(child)
		retrievedParent := child.GetParent()
		assert.NotNil(t, retrievedParent)

		parent.FreeRecursive()
	})
}

func TestNodeConfig(t *testing.T) {
	t.Run("has default config set", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		config := node.GetConfig()
		assert.NotNil(t, config)
		node.Free()
	})
}
