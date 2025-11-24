package tess

/*
#include <yoga/Yoga.h>
*/
import "C"
import (
	"fmt"
)

func toYGDisplay(display DisplayType) (C.YGDisplay, error) {
	switch display {
	case Flex:
		return C.YGDisplay(C.YGDisplayFlex), nil
	case Contents:
		return C.YGDisplay(C.YGDisplayContents), nil
	case None:
		return C.YGDisplay(C.YGDisplayNone), nil
	}

	return 0, fmt.Errorf("Unsupported display type '%s'", display.String())
}

func toYGAlign(align FlexAlign) (C.YGAlign, error) {
	switch align {
	case AlignAuto:
		return C.YGAlign(C.YGAlignAuto), nil
	case AlignStretch:
		return C.YGAlign(C.YGAlignStretch), nil
	case AlignBaseline:
		return C.YGAlign(C.YGAlignBaseline), nil
	case AlignStart:
		return C.YGAlign(C.YGAlignFlexStart), nil
	case AlignEnd:
		return C.YGAlign(C.YGAlignFlexEnd), nil
	case AlignCenter:
		return C.YGAlign(C.YGAlignCenter), nil
	case AlignSpaceBetween:
		return C.YGAlign(C.YGAlignSpaceBetween), nil
	case AlignSpaceAround:
		return C.YGAlign(C.YGAlignSpaceAround), nil
	case AlignSpaceEvenly:
		return C.YGAlign(C.YGAlignSpaceEvenly), nil
	}

	return 0, fmt.Errorf("Unsupported alignment for align-items/self '%s'", align.String())
}

func toYGJustify(align FlexJustify) (C.YGJustify, error) {
	switch align {
	case JustifyStart:
		return C.YGJustify(C.YGJustifyFlexStart), nil
	case JustifyEnd:
		return C.YGJustify(C.YGJustifyFlexEnd), nil
	case JustifyCenter:
		return C.YGJustify(C.YGJustifyCenter), nil
	case JustifySpaceBetween:
		return C.YGJustify(C.YGJustifySpaceBetween), nil
	case JustifySpaceAround:
		return C.YGJustify(C.YGJustifySpaceAround), nil
	case JustifySpaceEvenly:
		return C.YGJustify(C.YGJustifySpaceEvenly), nil
	}

	return 0, fmt.Errorf("Unsupported alignment justify-content/self '%s'", align.String())
}

func toYGFlexDirection(direction FlexDirection) (C.YGFlexDirection, error) {
	switch direction {
	case Column:
		return C.YGFlexDirection(C.YGFlexDirectionColumn), nil
	case ColumnReverse:
		return C.YGFlexDirection(C.YGFlexDirectionColumnReverse), nil
	case Row:
		return C.YGFlexDirection(C.YGFlexDirectionRow), nil
	case RowReverse:
		return C.YGFlexDirection(C.YGFlexDirectionRowReverse), nil
	}

	return 0, fmt.Errorf("Unsupported flex direction '%s'", direction.String())
}

func toYGWrap(wrap FlexWrap) (C.YGWrap, error) {
	switch wrap {
	case NoWrap:
		return C.YGWrap(C.YGWrapNoWrap), nil
	case Wrap:
		return C.YGWrap(C.YGWrapWrap), nil
	case WrapReverse:
		return C.YGWrap(C.YGWrapWrapReverse), nil
	}

	return 0, fmt.Errorf("Unsupported wrap type '%s'", wrap.String())
}

func toYGPositionType(position PositionType) (C.YGPositionType, error) {
	switch position {
	case Static:
		return C.YGPositionType(C.YGPositionTypeStatic), nil
	case Relative:
		return C.YGPositionType(C.YGPositionTypeRelative), nil
	case Absolute:
		return C.YGPositionType(C.YGPositionTypeAbsolute), nil
	}

	return 0, fmt.Errorf("Unsupported position type '%s'", position.String())
}

func toYGDirection(direction DirectionType) (C.YGDirection, error) {
	switch direction {
	case Inherit:
		return C.YGDirection(C.YGDirectionInherit), nil
	case LTR:
		return C.YGDirection(C.YGDirectionLTR), nil
	case RTL:
		return C.YGDirection(C.YGDirectionRTL), nil
	}

	return 0, fmt.Errorf("Unsupported direction type '%s'", direction.String())
}

func toYGOverflow(overflow OverflowType) (C.YGOverflow, error) {
	switch overflow {
	case Visible:
		return C.YGOverflow(C.YGOverflowVisible), nil
	case Hidden:
		return C.YGOverflow(C.YGOverflowHidden), nil
	case Scroll:
		return C.YGOverflow(C.YGOverflowScroll), nil
	}

	return 0, fmt.Errorf("Unsupported overflow type '%s'", overflow.String())
}

func toYGUnit(unit Unit) (C.YGUnit, error) {
	switch unit {
	case UnitUndefined:
		return C.YGUnit(C.YGUnitUndefined), nil
	case UnitPoint:
		return C.YGUnit(C.YGUnitPoint), nil
	case UnitPercent:
		return C.YGUnit(C.YGUnitPercent), nil
	case UnitAuto:
		return C.YGUnit(C.YGUnitAuto), nil
	case UnitMaxContent:
		return C.YGUnit(C.YGUnitMaxContent), nil
	case UnitFitContent:
		return C.YGUnit(C.YGUnitFitContent), nil
	case UnitStretch:
		return C.YGUnit(C.YGUnitStretch), nil
	}

	return 0, fmt.Errorf("Unknown unit '%s'", unit.String())
}

func toYGValue(value Value) (C.YGValue, error) {
	unit, err := toYGUnit(value.unit)
	if err != nil {
		return C.YGValue{}, fmt.Errorf("Cannot convert value '%s': %w", value.String(), err)
	}

	return C.YGValue{
		value: C.float(value.value),
		unit:  unit,
	}, nil
}

func fromYGValue(ygValue C.YGValue) Value {
	switch ygValue.unit {
	case C.YGUnitUndefined:
		return Undefined()
	case C.YGUnitPoint:
		return Point(float32(ygValue.value))
	case C.YGUnitPercent:
		return Percent(float32(ygValue.value))
	case C.YGUnitAuto:
		return Auto()
	case C.YGUnitMaxContent:
		return MaxContent()
	case C.YGUnitFitContent:
		return FitContent()
	case C.YGUnitStretch:
		return Stretch()
	}
	return Undefined()
}

func fromYGDisplay(ygDisplay C.YGDisplay) DisplayType {
	switch ygDisplay {
	case C.YGDisplayFlex:
		return Flex
	case C.YGDisplayContents:
		return Contents
	case C.YGDisplayNone:
		return None
	}
	return Flex
}

func fromYGAlign(ygAlign C.YGAlign) FlexAlign {
	switch ygAlign {
	case C.YGAlignAuto:
		return AlignAuto
	case C.YGAlignFlexStart:
		return AlignStart
	case C.YGAlignFlexEnd:
		return AlignEnd
	case C.YGAlignCenter:
		return AlignCenter
	case C.YGAlignStretch:
		return AlignStretch
	case C.YGAlignBaseline:
		return AlignBaseline
	case C.YGAlignSpaceBetween:
		return AlignSpaceBetween
	case C.YGAlignSpaceAround:
		return AlignSpaceAround
	case C.YGAlignSpaceEvenly:
		return AlignSpaceEvenly
	}
	return AlignAuto
}

func fromYGJustify(ygJustify C.YGJustify) FlexJustify {
	switch ygJustify {
	case C.YGJustifyFlexStart:
		return JustifyStart
	case C.YGJustifyFlexEnd:
		return JustifyEnd
	case C.YGJustifyCenter:
		return JustifyCenter
	case C.YGJustifySpaceBetween:
		return JustifySpaceBetween
	case C.YGJustifySpaceAround:
		return JustifySpaceAround
	case C.YGJustifySpaceEvenly:
		return JustifySpaceEvenly
	}
	return JustifyStart
}

func fromYGFlexDirection(ygDirection C.YGFlexDirection) FlexDirection {
	switch ygDirection {
	case C.YGFlexDirectionColumn:
		return Column
	case C.YGFlexDirectionColumnReverse:
		return ColumnReverse
	case C.YGFlexDirectionRow:
		return Row
	case C.YGFlexDirectionRowReverse:
		return RowReverse
	}
	return Row
}

func fromYGWrap(ygWrap C.YGWrap) FlexWrap {
	switch ygWrap {
	case C.YGWrapNoWrap:
		return NoWrap
	case C.YGWrapWrap:
		return Wrap
	case C.YGWrapWrapReverse:
		return WrapReverse
	}
	return NoWrap
}

func fromYGPositionType(ygPosition C.YGPositionType) PositionType {
	switch ygPosition {
	case C.YGPositionTypeStatic:
		return Static
	case C.YGPositionTypeRelative:
		return Relative
	case C.YGPositionTypeAbsolute:
		return Absolute
	}
	return Relative
}

func fromYGDirection(ygDirection C.YGDirection) DirectionType {
	switch ygDirection {
	case C.YGDirectionInherit:
		return Inherit
	case C.YGDirectionLTR:
		return LTR
	case C.YGDirectionRTL:
		return RTL
	}
	return Inherit
}

func fromYGOverflow(ygOverflow C.YGOverflow) OverflowType {
	switch ygOverflow {
	case C.YGOverflowVisible:
		return Visible
	case C.YGOverflowHidden:
		return Hidden
	case C.YGOverflowScroll:
		return Scroll
	}
	return Visible
}

func toYGNodeType(nodeType NodeType) (C.YGNodeType, error) {
	switch nodeType {
	case NodeTypeDefault:
		return C.YGNodeType(C.YGNodeTypeDefault), nil
	case NodeTypeText:
		return C.YGNodeType(C.YGNodeTypeText), nil
	}

	return 0, fmt.Errorf("Unknown node type '%s'", nodeType.String())
}

func fromYGNodeType(ygNodeType C.YGNodeType) NodeType {
	switch ygNodeType {
	case C.YGNodeTypeDefault:
		return NodeTypeDefault
	case C.YGNodeTypeText:
		return NodeTypeText
	}
	return NodeTypeDefault
}

func toYGBoxSizing(boxSizing BoxSizingType) (C.YGBoxSizing, error) {
	switch boxSizing {
	case ContentBox:
		return C.YGBoxSizing(C.YGBoxSizingContentBox), nil
	case BorderBox:
		return C.YGBoxSizing(C.YGBoxSizingBorderBox), nil
	}

	return 0, fmt.Errorf("Unknown box sizing '%s'", boxSizing.String())
}

func fromYGBoxSizing(ygBoxSizing C.YGBoxSizing) BoxSizingType {
	switch ygBoxSizing {
	case C.YGBoxSizingContentBox:
		return ContentBox
	case C.YGBoxSizingBorderBox:
		return BorderBox
	}
	return ContentBox
}
