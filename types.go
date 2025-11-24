package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

type DisplayType int

const (
	Flex DisplayType = iota
	Contents
	None
)

func (t DisplayType) String() string {
	switch t {
	case Flex:
		return "flex"
	case Contents:
		return "contents"
	case None:
		return "none"
	}

	return "unknown"
}

type FlexDirection int

const (
	Column FlexDirection = iota
	ColumnReverse
	Row
	RowReverse
)

func (a FlexDirection) String() string {
	switch a {
	case Column:
		return "column"
	case ColumnReverse:
		return "column-reverse"
	case Row:
		return "row"
	case RowReverse:
		return "row-reverse"
	}

	return "unknown"
}

type FlexAlign int

const (
	AlignAuto FlexAlign = iota
	AlignStretch
	AlignBaseline
	AlignStart
	AlignEnd
	AlignCenter
	AlignSpaceBetween
	AlignSpaceAround
	AlignSpaceEvenly
)

func (a FlexAlign) String() string {
	switch a {
	case AlignAuto:
		return "auto"
	case AlignStart:
		return "flex-start"
	case AlignCenter:
		return "center"
	case AlignEnd:
		return "flex-end"
	case AlignStretch:
		return "stretch"
	case AlignBaseline:
		return "baseline"
	case AlignSpaceBetween:
		return "space-between"
	case AlignSpaceAround:
		return "space-around"
	case AlignSpaceEvenly:
		return "space-evenly"
	}

	return "unknown"
}

type FlexJustify int

const (
	JustifyStart FlexJustify = iota
	JustifyEnd
	JustifyCenter
	JustifySpaceBetween
	JustifySpaceAround
	JustifySpaceEvenly
)

func (a FlexJustify) String() string {
	switch a {
	case JustifyStart:
		return "flex-start"
	case JustifyCenter:
		return "center"
	case JustifyEnd:
		return "flex-end"
	case JustifySpaceBetween:
		return "space-between"
	case JustifySpaceAround:
		return "space-around"
	case JustifySpaceEvenly:
		return "space-evenly"
	}

	return "unknown"
}

type FlexWrap int

const (
	NoWrap FlexWrap = iota
	Wrap
	WrapReverse
)

func (t FlexWrap) String() string {
	switch t {
	case NoWrap:
		return "no-wrap"
	case Wrap:
		return "wrap"
	case WrapReverse:
		return "wrap-reverse"
	}

	return "unknown"
}

type PositionType int

const (
	Static PositionType = iota
	Relative
	Absolute
)

func (t PositionType) String() string {
	switch t {
	case Static:
		return "static"
	case Relative:
		return "relative"
	case Absolute:
		return "absolute"
	}

	return "unknown"
}

type DirectionType int

const (
	Inherit DirectionType = iota
	LTR
	RTL
)

func (t DirectionType) String() string {
	switch t {
	case Inherit:
		return "inherit"
	case LTR:
		return "ltr"
	case RTL:
		return "rtl"
	}

	return "unknown"
}

type OverflowType int

const (
	Visible OverflowType = iota
	Hidden
	Scroll
)

func (t OverflowType) String() string {
	switch t {
	case Visible:
		return "visible"
	case Hidden:
		return "hidden"
	case Scroll:
		return "scroll"
	}

	return "unknown"
}

type BoxSizingType int

const (
	ContentBox BoxSizingType = iota
	BorderBox
)

func (t BoxSizingType) String() string {
	switch t {
	case ContentBox:
		return "content-box"
	case BorderBox:
		return "border-box"
	}

	return "unknown"
}
