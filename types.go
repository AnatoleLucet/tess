package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

// Value represents a dimension with a unit
type Value struct {
	Value float32
	Unit  Unit
}

// Predefined Value constants
var (
	ValueAuto      = Value{Unit: UnitAuto}
	ValueUndefined = Value{Unit: UnitUndefined}
	ValueZero      = Value{Value: 0, Unit: UnitPoint}
)

// Align constants
const (
	AlignAuto         = int(C.YGAlignAuto)
	AlignFlexStart    = int(C.YGAlignFlexStart)
	AlignCenter       = int(C.YGAlignCenter)
	AlignFlexEnd      = int(C.YGAlignFlexEnd)
	AlignStretch      = int(C.YGAlignStretch)
	AlignBaseline     = int(C.YGAlignBaseline)
	AlignSpaceBetween = int(C.YGAlignSpaceBetween)
	AlignSpaceAround  = int(C.YGAlignSpaceAround)
	AlignSpaceEvenly  = int(C.YGAlignSpaceEvenly)
)

// BoxSizing constants
const (
	BoxSizingBorderBox  = int(C.YGBoxSizingBorderBox)
	BoxSizingContentBox = int(C.YGBoxSizingContentBox)
)

// Dimension constants
const (
	DimensionWidth  = int(C.YGDimensionWidth)
	DimensionHeight = int(C.YGDimensionHeight)
)

// Direction constants
const (
	DirectionInherit = int(C.YGDirectionInherit)
	DirectionLTR     = int(C.YGDirectionLTR)
	DirectionRTL     = int(C.YGDirectionRTL)
)

// Display constants
const (
	DisplayFlex     = int(C.YGDisplayFlex)
	DisplayNone     = int(C.YGDisplayNone)
	DisplayContents = int(C.YGDisplayContents)
)

// Edge constants
const (
	EdgeLeft       = int(C.YGEdgeLeft)
	EdgeTop        = int(C.YGEdgeTop)
	EdgeRight      = int(C.YGEdgeRight)
	EdgeBottom     = int(C.YGEdgeBottom)
	EdgeStart      = int(C.YGEdgeStart)
	EdgeEnd        = int(C.YGEdgeEnd)
	EdgeHorizontal = int(C.YGEdgeHorizontal)
	EdgeVertical   = int(C.YGEdgeVertical)
	EdgeAll        = int(C.YGEdgeAll)
)

// FlexDirection constants
const (
	FlexDirectionColumn        = int(C.YGFlexDirectionColumn)
	FlexDirectionColumnReverse = int(C.YGFlexDirectionColumnReverse)
	FlexDirectionRow           = int(C.YGFlexDirectionRow)
	FlexDirectionRowReverse    = int(C.YGFlexDirectionRowReverse)
)

// Gutter constants
const (
	GutterColumn = int(C.YGGutterColumn)
	GutterRow    = int(C.YGGutterRow)
	GutterAll    = int(C.YGGutterAll)
)

// Justify constants
const (
	JustifyFlexStart    = int(C.YGJustifyFlexStart)
	JustifyCenter       = int(C.YGJustifyCenter)
	JustifyFlexEnd      = int(C.YGJustifyFlexEnd)
	JustifySpaceBetween = int(C.YGJustifySpaceBetween)
	JustifySpaceAround  = int(C.YGJustifySpaceAround)
	JustifySpaceEvenly  = int(C.YGJustifySpaceEvenly)
)

// MeasureMode constants
const (
	MeasureModeUndefined = int(C.YGMeasureModeUndefined)
	MeasureModeExactly   = int(C.YGMeasureModeExactly)
	MeasureModeAtMost    = int(C.YGMeasureModeAtMost)
)

// NodeType constants
const (
	NodeTypeDefault = int(C.YGNodeTypeDefault)
	NodeTypeText    = int(C.YGNodeTypeText)
)

// Overflow constants
const (
	OverflowVisible = int(C.YGOverflowVisible)
	OverflowHidden  = int(C.YGOverflowHidden)
	OverflowScroll  = int(C.YGOverflowScroll)
)

// PositionType constants
const (
	PositionTypeStatic   = int(C.YGPositionTypeStatic)
	PositionTypeRelative = int(C.YGPositionTypeRelative)
	PositionTypeAbsolute = int(C.YGPositionTypeAbsolute)
)

// Unit type for Value
type Unit int

// Unit constants
const (
	UnitUndefined  Unit = Unit(C.YGUnitUndefined)
	UnitPoint      Unit = Unit(C.YGUnitPoint)
	UnitPercent    Unit = Unit(C.YGUnitPercent)
	UnitAuto       Unit = Unit(C.YGUnitAuto)
	UnitMaxContent Unit = Unit(C.YGUnitMaxContent)
	UnitFitContent Unit = Unit(C.YGUnitFitContent)
	UnitStretch    Unit = Unit(C.YGUnitStretch)
)

// Wrap constants
const (
	WrapNoWrap      = int(C.YGWrapNoWrap)
	WrapWrap        = int(C.YGWrapWrap)
	WrapWrapReverse = int(C.YGWrapWrapReverse)
)
