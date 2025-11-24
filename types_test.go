package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayType(t *testing.T) {
	t.Run("DisplayType String method", func(t *testing.T) {
		assert.Equal(t, "flex", Flex.String())
		assert.Equal(t, "contents", Contents.String())
		assert.Equal(t, "none", None.String())
		assert.Equal(t, "unknown", DisplayType(999).String())
	})
}

func TestFlexDirection(t *testing.T) {
	t.Run("FlexDirection String method", func(t *testing.T) {
		assert.Equal(t, "column", Column.String())
		assert.Equal(t, "column-reverse", ColumnReverse.String())
		assert.Equal(t, "row", Row.String())
		assert.Equal(t, "row-reverse", RowReverse.String())
		assert.Equal(t, "unknown", FlexDirection(999).String())
	})
}

func TestFlexAlign(t *testing.T) {
	t.Run("FlexAlign String method", func(t *testing.T) {
		assert.Equal(t, "auto", AlignAuto.String())
		assert.Equal(t, "stretch", AlignStretch.String())
		assert.Equal(t, "baseline", AlignBaseline.String())
		assert.Equal(t, "flex-start", AlignStart.String())
		assert.Equal(t, "flex-end", AlignEnd.String())
		assert.Equal(t, "center", AlignCenter.String())
		assert.Equal(t, "space-between", AlignSpaceBetween.String())
		assert.Equal(t, "space-around", AlignSpaceAround.String())
		assert.Equal(t, "space-evenly", AlignSpaceEvenly.String())
		assert.Equal(t, "unknown", FlexAlign(999).String())
	})
}

func TestFlexJustify(t *testing.T) {
	t.Run("FlexJustify String method", func(t *testing.T) {
		assert.Equal(t, "flex-start", JustifyStart.String())
		assert.Equal(t, "flex-end", JustifyEnd.String())
		assert.Equal(t, "center", JustifyCenter.String())
		assert.Equal(t, "space-between", JustifySpaceBetween.String())
		assert.Equal(t, "space-around", JustifySpaceAround.String())
		assert.Equal(t, "space-evenly", JustifySpaceEvenly.String())
		assert.Equal(t, "unknown", FlexJustify(999).String())
	})
}

func TestFlexWrap(t *testing.T) {
	t.Run("FlexWrap String method", func(t *testing.T) {
		assert.Equal(t, "no-wrap", NoWrap.String())
		assert.Equal(t, "wrap", Wrap.String())
		assert.Equal(t, "wrap-reverse", WrapReverse.String())
		assert.Equal(t, "unknown", FlexWrap(999).String())
	})
}

func TestPositionType(t *testing.T) {
	t.Run("PositionType String method", func(t *testing.T) {
		assert.Equal(t, "relative", Relative.String())
		assert.Equal(t, "absolute", Absolute.String())
		assert.Equal(t, "unknown", PositionType(999).String())
	})
}

func TestOverflowType(t *testing.T) {
	t.Run("OverflowType String method", func(t *testing.T) {
		assert.Equal(t, "visible", Visible.String())
		assert.Equal(t, "hidden", Hidden.String())
		assert.Equal(t, "scroll", Scroll.String())
		assert.Equal(t, "unknown", OverflowType(999).String())
	})
}

func TestDirectionType(t *testing.T) {
	t.Run("DirectionType String method", func(t *testing.T) {
		assert.Equal(t, "inherit", Inherit.String())
		assert.Equal(t, "ltr", LTR.String())
		assert.Equal(t, "rtl", RTL.String())
		assert.Equal(t, "unknown", DirectionType(999).String())
	})
}

func TestBoxSizingType(t *testing.T) {
	t.Run("BoxSizingType String method", func(t *testing.T) {
		assert.Equal(t, "content-box", ContentBox.String())
		assert.Equal(t, "border-box", BorderBox.String())
		assert.Equal(t, "unknown", BoxSizingType(999).String())
	})
}
