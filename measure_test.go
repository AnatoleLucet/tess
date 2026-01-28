package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMeasureFunc(t *testing.T) {
	t.Run("sets measure function", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		assert.False(t, node.HasMeasureFunc())

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			return Size{Width: 100, Height: 50}
		})

		assert.True(t, node.HasMeasureFunc())

		node.Free()
	})

	t.Run("measure function is called during layout", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		called := false
		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			called = true
			return Size{Width: 100, Height: 50}
		})

		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)
		assert.True(t, called)

		layout := node.GetLayout()
		assert.Equal(t, float32(100), layout.Width())
		assert.Equal(t, float32(50), layout.Height())

		node.Free()
	})

	t.Run("measure function receives correct parameters", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		parent.SetWidth(Point(200))
		parent.SetHeight(Point(100))

		child, err := NewNode()
		assert.NoError(t, err)

		var receivedWidth float32
		var receivedWidthMode MeasureMode

		child.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			receivedWidth = width
			receivedWidthMode = widthMode
			return Size{Width: 50, Height: 25}
		})

		parent.AppendChild(child)
		err = parent.ComputeLayout(Container{})
		assert.NoError(t, err)

		// The child should receive AtMost width constraint from the parent
		assert.Equal(t, float32(200), receivedWidth)
		assert.Equal(t, MeasureModeAtMost, receivedWidthMode)

		parent.FreeRecursive()
	})

	t.Run("measure function with undefined constraints", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		var receivedWidthMode, receivedHeightMode MeasureMode

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			receivedWidthMode = widthMode
			receivedHeightMode = heightMode
			return Size{Width: 100, Height: 50}
		})

		// No container constraints
		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)

		// Should receive undefined mode when no constraints are set
		assert.Equal(t, MeasureModeUndefined, receivedWidthMode)
		assert.Equal(t, MeasureModeUndefined, receivedHeightMode)

		node.Free()
	})

	t.Run("can read node values inside measure function", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetPadding(Edges{All: Point(10)})
		var padding *StyleEdges

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			padding = node.GetPadding()
			return Size{Width: 100, Height: 50}
		})

		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)

		assert.Equal(t, Point(10), padding.GetTop())
		assert.Equal(t, Point(10), padding.GetBottom())
		assert.Equal(t, Point(10), padding.GetLeft())
		assert.Equal(t, Point(10), padding.GetRight())
		node.Free()
	})
}

func TestUnsetMeasureFunc(t *testing.T) {
	t.Run("unsets measure function", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			return Size{Width: 100, Height: 50}
		})
		assert.True(t, node.HasMeasureFunc())

		node.UnsetMeasureFunc()
		assert.False(t, node.HasMeasureFunc())

		node.Free()
	})

	t.Run("replacing measure function cleans up old handle", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		callCount := 0

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			callCount = 1
			return Size{Width: 100, Height: 50}
		})

		// Replace with a new measure function
		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			callCount = 2
			return Size{Width: 200, Height: 100}
		})

		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)

		// Should call the new function
		assert.Equal(t, 2, callCount)

		layout := node.GetLayout()
		assert.Equal(t, float32(200), layout.Width())
		assert.Equal(t, float32(100), layout.Height())

		node.Free()
	})
}

func TestMarkDirtyWithMeasureFunc(t *testing.T) {
	t.Run("MarkDirty triggers re-measurement", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		measureCount := 0
		currentSize := Size{Width: 100, Height: 50}

		node.SetMeasureFunc(func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
			measureCount++
			return currentSize
		})

		// First layout
		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)
		assert.Equal(t, 1, measureCount)

		layout := node.GetLayout()
		assert.Equal(t, float32(100), layout.Width())
		assert.Equal(t, float32(50), layout.Height())

		// Change the size and mark dirty
		currentSize = Size{Width: 200, Height: 100}
		node.MarkDirty()

		// Second layout should re-measure
		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)
		assert.Equal(t, 2, measureCount)

		layout = node.GetLayout()
		assert.Equal(t, float32(200), layout.Width())
		assert.Equal(t, float32(100), layout.Height())

		node.Free()
	})
}

func TestMeasureModeString(t *testing.T) {
	assert.Equal(t, "undefined", MeasureModeUndefined.String())
	assert.Equal(t, "exactly", MeasureModeExactly.String())
	assert.Equal(t, "at-most", MeasureModeAtMost.String())
	assert.Equal(t, "unknown", MeasureMode(99).String())
}
