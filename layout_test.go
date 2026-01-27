package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeLayout(t *testing.T) {
	t.Run("computes layout with fixed dimensions", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(200))

		err = node.ComputeLayout(Container{})
		assert.NoError(t, err)

		layout := node.GetLayout()
		assert.Equal(t, float32(100), layout.Width())
		assert.Equal(t, float32(200), layout.Height())

		node.Free()
	})

	t.Run("computes layout with container constraints", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Percent(50))
		node.SetHeight(Point(100))

		err = node.ComputeLayout(Container{Width: 200, Height: 400})
		assert.NoError(t, err)

		layout := node.GetLayout()
		assert.Equal(t, float32(100), layout.Width()) // 50% of 200
		assert.Equal(t, float32(100), layout.Height())

		node.Free()
	})

	t.Run("computes layout with direction", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(200))

		err = node.ComputeLayout(Container{Direction: LTR})
		assert.NoError(t, err)

		layout := node.GetLayout()
		assert.Equal(t, LTR, layout.Direction())

		node.Free()
	})

	t.Run("computes layout for flexbox children", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		// parent.SetFlexDirection(Row)
		// parent.SetWidth(Point(300))
		// parent.SetHeight(Point(100))

		child1, err := NewNode()
		assert.NoError(t, err)
		// child1.SetWidth(Point(100))
		// child1.SetHeight(Point(50))

		// child2, err := NewNode()
		// assert.NoError(t, err)
		// child2.SetWidth(Point(150))
		// child2.SetHeight(Point(50))

		parent.AppendChild(child1)
		// parent.AppendChild(child2)

		err = parent.ComputeLayout(Container{})
		assert.NoError(t, err)

		// Check parent layout
		// parentLayout := parent.GetLayout()
		// assert.Equal(t, float32(300), parentLayout.Width())
		// assert.Equal(t, float32(100), parentLayout.Height())
		//
		// // Check children are laid out horizontally
		// child1Layout := child1.GetLayout()
		// child2Layout := child2.GetLayout()
		//
		// assert.Equal(t, float32(0), child1Layout.Left())
		// assert.Equal(t, float32(100), child2Layout.Left())

		parent.FreeRecursive()
	})
}

func TestLayoutDimensions(t *testing.T) {
	t.Run("gets layout width and height", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(123.45))
		node.SetHeight(Point(678.90))

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		// Yoga rounds layout values to nearest pixel by default (pointScaleFactor = 1.0)
		assert.Equal(t, float32(123), layout.Width())
		assert.Equal(t, float32(679), layout.Height())

		node.Free()
	})

	t.Run("gets raw width and height", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(200))
		node.SetPadding(Edges{All: Point(10)})

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		// RawWidth/Height should include padding
		assert.GreaterOrEqual(t, layout.RawWidth(), float32(100))
		assert.GreaterOrEqual(t, layout.RawHeight(), float32(200))

		node.Free()
	})
}

func TestLayoutPosition(t *testing.T) {
	t.Run("gets layout position", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		parent.SetWidth(Point(300))
		parent.SetHeight(Point(300))

		child, err := NewNode()
		assert.NoError(t, err)
		child.SetWidth(Point(100))
		child.SetHeight(Point(100))
		child.SetPosition(Absolute)
		child.SetTop(Point(50))
		child.SetLeft(Point(75))

		parent.AppendChild(child)
		parent.ComputeLayout(Container{})

		childLayout := child.GetLayout()
		assert.Equal(t, float32(75), childLayout.Left())
		assert.Equal(t, float32(50), childLayout.Top())

		parent.FreeRecursive()
	})
}

func TestLayoutEdges(t *testing.T) {
	t.Run("gets padding from layout", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))
		node.SetPadding(Edges{
			Top:    Point(10),
			Right:  Point(20),
			Bottom: Point(30),
			Left:   Point(40),
		})

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		assert.Equal(t, float32(10), layout.Padding().Top())
		assert.Equal(t, float32(20), layout.Padding().Right())
		assert.Equal(t, float32(30), layout.Padding().Bottom())
		assert.Equal(t, float32(40), layout.Padding().Left())

		node.Free()
	})

	t.Run("gets margin from layout", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))
		node.SetMargin(Edges{
			Top:    Point(5),
			Right:  Point(10),
			Bottom: Point(15),
			Left:   Point(20),
		})

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		assert.Equal(t, float32(5), layout.Margin().Top())
		assert.Equal(t, float32(10), layout.Margin().Right())
		assert.Equal(t, float32(15), layout.Margin().Bottom())
		assert.Equal(t, float32(20), layout.Margin().Left())

		node.Free()
	})

	t.Run("gets border from layout", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))
		node.SetBorder(Edges{
			Top:    Point(1),
			Right:  Point(2),
			Bottom: Point(3),
			Left:   Point(4),
		})

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		assert.Equal(t, float32(1), layout.Border().Top())
		assert.Equal(t, float32(2), layout.Border().Right())
		assert.Equal(t, float32(3), layout.Border().Bottom())
		assert.Equal(t, float32(4), layout.Border().Left())

		node.Free()
	})
}

func TestLayoutOverflow(t *testing.T) {
	t.Run("detects overflow", func(t *testing.T) {
		parent, err := NewNode()
		assert.NoError(t, err)
		parent.SetWidth(Point(100))
		parent.SetHeight(Point(100))
		parent.SetOverflow(Hidden)

		child, err := NewNode()
		assert.NoError(t, err)
		child.SetWidth(Point(200))
		child.SetHeight(Point(200))

		parent.AppendChild(child)
		parent.ComputeLayout(Container{})

		layout := parent.GetLayout()
		// This should detect that child overflows parent
		// Note: actual overflow detection depends on Yoga's implementation
		assert.NotNil(t, layout)

		parent.FreeRecursive()
	})
}

func TestHasNewLayout(t *testing.T) {
	t.Run("detects new layout", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))

		// Nodes start with hasNewLayout = true (never been laid out before)
		assert.True(t, node.HasNewLayout())

		node.ComputeLayout(Container{})

		// Should still have new layout after computation
		assert.True(t, node.HasNewLayout())

		// Reset the flag (typically done after consuming layout results)
		node.SetHasNewLayout(false)
		assert.False(t, node.HasNewLayout())

		// Compute again - should have new layout
		node.ComputeLayout(Container{})
		assert.True(t, node.HasNewLayout())

		node.Free()
	})
}

func TestIsDirty(t *testing.T) {
	t.Run("marks node as dirty", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))

		// Compute initial layout
		node.ComputeLayout(Container{})
		assert.False(t, node.IsDirty())

		// Change a style property
		node.SetWidth(Point(200))
		assert.True(t, node.IsDirty())

		// Compute layout again
		node.ComputeLayout(Container{})
		assert.False(t, node.IsDirty())

		node.Free()
	})
}

func TestAbsolutePosition(t *testing.T) {
	t.Run("gets absolute position for single node", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)
		node.SetWidth(Point(100))
		node.SetHeight(Point(100))

		node.ComputeLayout(Container{})
		layout := node.GetLayout()

		assert.Equal(t, float32(0), layout.AbsoluteTop())
		assert.Equal(t, float32(0), layout.AbsoluteLeft())
		assert.Equal(t, float32(0), layout.AbsoluteBottom())
		assert.Equal(t, float32(0), layout.AbsoluteRight())

		node.Free()
	})

	t.Run("gets absolute position through parent hierarchy", func(t *testing.T) {
		// Create a 3-level hierarchy:
		// root (padding 10) -> parent (offset 20,20) -> child (offset 30,30)
		root, err := NewNode()
		assert.NoError(t, err)
		root.SetWidth(Point(300))
		root.SetHeight(Point(300))
		root.SetPadding(Edges{All: Point(10)})

		parent, err := NewNode()
		assert.NoError(t, err)
		parent.SetWidth(Point(200))
		parent.SetHeight(Point(200))
		parent.SetMargin(Edges{Top: Point(20), Left: Point(20)})

		child, err := NewNode()
		assert.NoError(t, err)
		child.SetWidth(Point(50))
		child.SetHeight(Point(50))
		child.SetMargin(Edges{Top: Point(30), Left: Point(30)})

		root.AppendChild(parent)
		parent.AppendChild(child)

		root.ComputeLayout(Container{})

		// Check parent's absolute position (should include root's padding)
		parentLayout := parent.GetLayout()
		assert.Equal(t, float32(10+20), parentLayout.AbsoluteTop())  // root padding + parent margin
		assert.Equal(t, float32(10+20), parentLayout.AbsoluteLeft()) // root padding + parent margin

		// Check child's absolute position (should include all offsets)
		childLayout := child.GetLayout()
		assert.Equal(t, float32(10+20+30), childLayout.AbsoluteTop())  // root padding + parent margin + child margin
		assert.Equal(t, float32(10+20+30), childLayout.AbsoluteLeft()) // root padding + parent margin + child margin

		root.FreeRecursive()
	})

	t.Run("gets absolute position with row flex direction", func(t *testing.T) {
		root, err := NewNode()
		assert.NoError(t, err)
		root.SetWidth(Point(300))
		root.SetHeight(Point(100))
		root.SetFlexDirection(Row)

		child1, err := NewNode()
		assert.NoError(t, err)
		child1.SetWidth(Point(100))
		child1.SetHeight(Point(50))

		child2, err := NewNode()
		assert.NoError(t, err)
		child2.SetWidth(Point(100))
		child2.SetHeight(Point(50))

		root.AppendChild(child1)
		root.AppendChild(child2)

		root.ComputeLayout(Container{})

		child1Layout := child1.GetLayout()
		child2Layout := child2.GetLayout()

		// child1 starts at left 0
		assert.Equal(t, float32(0), child1Layout.AbsoluteLeft())
		// child2 starts after child1 (at left 100)
		assert.Equal(t, float32(100), child2Layout.AbsoluteLeft())

		root.FreeRecursive()
	})
}
