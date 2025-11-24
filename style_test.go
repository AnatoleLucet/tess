package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleDisplay(t *testing.T) {
	t.Run("sets and gets display", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetDisplay(None)
		assert.Equal(t, None, node.GetDisplay())

		node.SetDisplay(Flex)
		assert.Equal(t, Flex, node.GetDisplay())

		node.Free()
	})
}

func TestStyleFlexDirection(t *testing.T) {
	t.Run("sets and gets flex direction", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetFlexDirection(Row)
		assert.Equal(t, Row, node.GetFlexDirection())

		node.SetFlexDirection(Column)
		assert.Equal(t, Column, node.GetFlexDirection())

		node.SetFlexDirection(RowReverse)
		assert.Equal(t, RowReverse, node.GetFlexDirection())

		node.Free()
	})
}

func TestStyleAlignment(t *testing.T) {
	t.Run("sets and gets justify content", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetJustifyContent(JustifyCenter)
		assert.Equal(t, JustifyCenter, node.GetJustifyContent())

		node.SetJustifyContent(JustifySpaceBetween)
		assert.Equal(t, JustifySpaceBetween, node.GetJustifyContent())

		node.Free()
	})

	t.Run("sets and gets align items", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetAlignItems(AlignStretch)
		assert.Equal(t, AlignStretch, node.GetAlignItems())

		node.SetAlignItems(AlignCenter)
		assert.Equal(t, AlignCenter, node.GetAlignItems())

		node.Free()
	})

	t.Run("sets and gets align self", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetAlignSelf(AlignAuto)
		assert.Equal(t, AlignAuto, node.GetAlignSelf())

		node.SetAlignSelf(AlignEnd)
		assert.Equal(t, AlignEnd, node.GetAlignSelf())

		node.Free()
	})

	t.Run("sets and gets align content", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetAlignContent(AlignSpaceAround)
		assert.Equal(t, AlignSpaceAround, node.GetAlignContent())

		node.Free()
	})
}

func TestStyleFlexWrap(t *testing.T) {
	t.Run("sets and gets flex wrap", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetFlexWrap(Wrap)
		assert.Equal(t, Wrap, node.GetFlexWrap())

		node.SetFlexWrap(NoWrap)
		assert.Equal(t, NoWrap, node.GetFlexWrap())

		node.Free()
	})
}

func TestStyleDirection(t *testing.T) {
	t.Run("sets and gets direction", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetDirection(LTR)
		assert.Equal(t, LTR, node.GetDirection())

		node.SetDirection(RTL)
		assert.Equal(t, RTL, node.GetDirection())

		node.Free()
	})
}

func TestStyleDimensions(t *testing.T) {
	t.Run("sets and gets width", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetWidth(Point(100))
		assert.Equal(t, Point(100), node.GetWidth())

		node.SetWidth(Percent(50))
		assert.Equal(t, Percent(50), node.GetWidth())

		node.SetWidth(Auto())
		assert.Equal(t, Auto(), node.GetWidth())

		node.Free()
	})

	t.Run("sets and gets height", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetHeight(Point(200))
		assert.Equal(t, Point(200), node.GetHeight())

		node.SetHeight(Percent(75))
		assert.Equal(t, Percent(75), node.GetHeight())

		node.Free()
	})

	t.Run("sets and gets min dimensions", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetMinWidth(Point(50))
		assert.Equal(t, Point(50), node.GetMinWidth())

		node.SetMinHeight(Point(30))
		assert.Equal(t, Point(30), node.GetMinHeight())

		node.Free()
	})

	t.Run("sets and gets max dimensions", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetMaxWidth(Point(500))
		assert.Equal(t, Point(500), node.GetMaxWidth())

		node.SetMaxHeight(Point(300))
		assert.Equal(t, Point(300), node.GetMaxHeight())

		node.Free()
	})
}

func TestStyleAspectRatio(t *testing.T) {
	t.Run("sets and gets aspect ratio", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetAspectRatio(16.0 / 9.0)
		assert.InDelta(t, 16.0/9.0, node.GetAspectRatio(), 0.001)

		node.Free()
	})
}

func TestStyleFlex(t *testing.T) {
	t.Run("sets and gets flex grow", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetFlexGrow(1.5)
		assert.Equal(t, float32(1.5), node.GetFlexGrow())

		node.Free()
	})

	t.Run("sets and gets flex shrink", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetFlexShrink(0.5)
		assert.Equal(t, float32(0.5), node.GetFlexShrink())

		node.Free()
	})

	t.Run("sets and gets flex basis", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetFlexBasis(Point(100))
		assert.Equal(t, Point(100), node.GetFlexBasis())

		node.SetFlexBasis(Percent(50))
		assert.Equal(t, Percent(50), node.GetFlexBasis())

		node.SetFlexBasis(Auto())
		assert.Equal(t, Auto(), node.GetFlexBasis())

		node.Free()
	})
}

func TestStylePosition(t *testing.T) {
	t.Run("sets and gets position type", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetPosition(Absolute)
		assert.Equal(t, Absolute, node.GetPosition())

		node.SetPosition(Relative)
		assert.Equal(t, Relative, node.GetPosition())

		node.Free()
	})

	t.Run("sets and gets position edges", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetTop(Point(10))
		assert.Equal(t, Point(10), node.GetTop())

		node.SetRight(Point(20))
		assert.Equal(t, Point(20), node.GetRight())

		node.SetBottom(Point(30))
		assert.Equal(t, Point(30), node.GetBottom())

		node.SetLeft(Point(40))
		assert.Equal(t, Point(40), node.GetLeft())

		node.Free()
	})

	t.Run("sets position with percentage", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetTop(Percent(10))
		assert.Equal(t, Percent(10), node.GetTop())

		node.Free()
	})
}

func TestStyleOverflow(t *testing.T) {
	t.Run("sets and gets overflow", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetOverflow(Hidden)
		assert.Equal(t, Hidden, node.GetOverflow())

		node.SetOverflow(Scroll)
		assert.Equal(t, Scroll, node.GetOverflow())

		node.Free()
	})
}

func TestStyleBoxSizing(t *testing.T) {
	t.Run("sets and gets box sizing", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetBoxSizing(BorderBox)
		assert.Equal(t, BorderBox, node.GetBoxSizing())

		node.SetBoxSizing(ContentBox)
		assert.Equal(t, ContentBox, node.GetBoxSizing())

		node.Free()
	})
}

func TestStylePadding(t *testing.T) {
	t.Run("sets and gets padding", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetPadding(Edges{
			Top:    Point(10),
			Right:  Point(20),
			Bottom: Point(30),
			Left:   Point(40),
		})

		assert.Equal(t, Point(10), node.GetPadding().GetTop())
		assert.Equal(t, Point(20), node.GetPadding().GetRight())
		assert.Equal(t, Point(30), node.GetPadding().GetBottom())
		assert.Equal(t, Point(40), node.GetPadding().GetLeft())

		node.Free()
	})

	t.Run("sets padding with all shorthand", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetPadding(Edges{All: Point(15)})

		assert.Equal(t, Point(15), node.GetPadding().GetTop())
		assert.Equal(t, Point(15), node.GetPadding().GetRight())
		assert.Equal(t, Point(15), node.GetPadding().GetBottom())
		assert.Equal(t, Point(15), node.GetPadding().GetLeft())

		node.Free()
	})

	t.Run("sets padding with percentage", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetPadding(Edges{All: Percent(10)})

		assert.Equal(t, Percent(10), node.GetPadding().GetTop())

		node.Free()
	})
}

func TestStyleMargin(t *testing.T) {
	t.Run("sets and gets margin", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetMargin(Edges{
			Top:    Point(5),
			Right:  Point(10),
			Bottom: Point(15),
			Left:   Point(20),
		})

		assert.Equal(t, Point(5), node.GetMargin().GetTop())
		assert.Equal(t, Point(10), node.GetMargin().GetRight())
		assert.Equal(t, Point(15), node.GetMargin().GetBottom())
		assert.Equal(t, Point(20), node.GetMargin().GetLeft())

		node.Free()
	})

	t.Run("sets margin with auto", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetMargin(Edges{
			Left:  Auto(),
			Right: Auto(),
		})

		assert.Equal(t, Auto(), node.GetMargin().GetLeft())
		assert.Equal(t, Auto(), node.GetMargin().GetRight())

		node.Free()
	})
}

func TestStyleBorder(t *testing.T) {
	t.Run("sets and gets border", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetBorder(Edges{
			Top:    Point(1),
			Right:  Point(2),
			Bottom: Point(3),
			Left:   Point(4),
		})

		assert.Equal(t, Point(1), node.GetBorder().GetTop())
		assert.Equal(t, Point(2), node.GetBorder().GetRight())
		assert.Equal(t, Point(3), node.GetBorder().GetBottom())
		assert.Equal(t, Point(4), node.GetBorder().GetLeft())

		node.Free()
	})

	t.Run("sets border with all shorthand", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetBorder(Edges{All: Point(2)})

		assert.Equal(t, Point(2), node.GetBorder().GetTop())
		assert.Equal(t, Point(2), node.GetBorder().GetRight())
		assert.Equal(t, Point(2), node.GetBorder().GetBottom())
		assert.Equal(t, Point(2), node.GetBorder().GetLeft())

		node.Free()
	})
}

func TestStyleGap(t *testing.T) {
	t.Run("sets and gets gap", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetGap(Gap{
			Row:    Point(10),
			Column: Point(20),
		})

		assert.Equal(t, Point(10), node.GetGap().GetRow())
		assert.Equal(t, Point(20), node.GetGap().GetColumn())

		node.Free()
	})

	t.Run("sets gap with all shorthand", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		node.SetGap(Gap{All: Point(15)})

		assert.Equal(t, Point(15), node.GetGap().GetAll())

		node.Free()
	})
}

func TestStyleApply(t *testing.T) {
	t.Run("applies complete style", func(t *testing.T) {
		style := &Style{
			Display:        Flex,
			FlexDirection:  Row,
			JustifyContent: JustifyCenter,
			AlignItems:     AlignStretch,
			Width:          Point(300),
			Height:         Point(200),
			Padding:        Edges{All: Point(10)},
			Margin:         Edges{All: Point(5)},
			FlexGrow:       1.0,
		}

		node, err := NewNode(style)
		assert.NoError(t, err)

		assert.Equal(t, Flex, node.GetDisplay())
		assert.Equal(t, Row, node.GetFlexDirection())
		assert.Equal(t, JustifyCenter, node.GetJustifyContent())
		assert.Equal(t, AlignStretch, node.GetAlignItems())
		assert.Equal(t, Point(300), node.GetWidth())
		assert.Equal(t, Point(200), node.GetHeight())
		assert.Equal(t, Point(10), node.GetPadding().GetTop())
		assert.Equal(t, Point(5), node.GetMargin().GetTop())
		assert.Equal(t, float32(1.0), node.GetFlexGrow())

		node.Free()
	})

	t.Run("applies style after creation", func(t *testing.T) {
		node, err := NewNode()
		assert.NoError(t, err)

		style := &Style{
			Width:  Point(100),
			Height: Point(100),
		}

		node.Apply(style)

		assert.Equal(t, Point(100), node.GetWidth())
		assert.Equal(t, Point(100), node.GetHeight())

		node.Free()
	})

	t.Run("applies multiple styles", func(t *testing.T) {
		style1 := &Style{
			Width: Point(100),
		}
		style2 := &Style{
			Height: Point(200),
		}

		node, err := NewNode(style1, style2)
		assert.NoError(t, err)

		assert.Equal(t, Point(100), node.GetWidth())
		assert.Equal(t, Point(200), node.GetHeight())

		node.Free()
	})
}
