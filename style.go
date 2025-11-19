package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

// CopyStyle copies the style from srcNode to the current node
func (n *Node) CopyStyle(srcNode *Node) {
	C.YGNodeCopyStyle(n.node, srcNode.node)
}

// Direction

func (n *Node) SetDirection(direction int) {
	C.YGNodeStyleSetDirection(n.node, C.YGDirection(direction))
}

func (n *Node) GetDirection() int {
	return int(C.YGNodeStyleGetDirection(n.node))
}

// FlexDirection

func (n *Node) SetFlexDirection(direction int) {
	C.YGNodeStyleSetFlexDirection(n.node, C.YGFlexDirection(direction))
}

func (n *Node) GetFlexDirection() int {
	return int(C.YGNodeStyleGetFlexDirection(n.node))
}

// JustifyContent

func (n *Node) SetJustifyContent(justify int) {
	C.YGNodeStyleSetJustifyContent(n.node, C.YGJustify(justify))
}

func (n *Node) GetJustifyContent() int {
	return int(C.YGNodeStyleGetJustifyContent(n.node))
}

// AlignContent

func (n *Node) SetAlignContent(align int) {
	C.YGNodeStyleSetAlignContent(n.node, C.YGAlign(align))
}

func (n *Node) GetAlignContent() int {
	return int(C.YGNodeStyleGetAlignContent(n.node))
}

// AlignItems

func (n *Node) SetAlignItems(align int) {
	C.YGNodeStyleSetAlignItems(n.node, C.YGAlign(align))
}

func (n *Node) GetAlignItems() int {
	return int(C.YGNodeStyleGetAlignItems(n.node))
}

// AlignSelf

func (n *Node) SetAlignSelf(align int) {
	C.YGNodeStyleSetAlignSelf(n.node, C.YGAlign(align))
}

func (n *Node) GetAlignSelf() int {
	return int(C.YGNodeStyleGetAlignSelf(n.node))
}

// PositionType

func (n *Node) SetPositionType(positionType int) {
	C.YGNodeStyleSetPositionType(n.node, C.YGPositionType(positionType))
}

func (n *Node) GetPositionType() int {
	return int(C.YGNodeStyleGetPositionType(n.node))
}

// FlexWrap

func (n *Node) SetFlexWrap(wrap int) {
	C.YGNodeStyleSetFlexWrap(n.node, C.YGWrap(wrap))
}

func (n *Node) GetFlexWrap() int {
	return int(C.YGNodeStyleGetFlexWrap(n.node))
}

// Overflow

func (n *Node) SetOverflow(overflow int) {
	C.YGNodeStyleSetOverflow(n.node, C.YGOverflow(overflow))
}

func (n *Node) GetOverflow() int {
	return int(C.YGNodeStyleGetOverflow(n.node))
}

// Display

func (n *Node) SetDisplay(display int) {
	C.YGNodeStyleSetDisplay(n.node, C.YGDisplay(display))
}

func (n *Node) GetDisplay() int {
	return int(C.YGNodeStyleGetDisplay(n.node))
}

// Flex

func (n *Node) SetFlex(flex float32) {
	C.YGNodeStyleSetFlex(n.node, C.float(flex))
}

func (n *Node) GetFlex() float32 {
	return float32(C.YGNodeStyleGetFlex(n.node))
}

// FlexGrow

func (n *Node) SetFlexGrow(flexGrow float32) {
	C.YGNodeStyleSetFlexGrow(n.node, C.float(flexGrow))
}

func (n *Node) GetFlexGrow() float32 {
	return float32(C.YGNodeStyleGetFlexGrow(n.node))
}

// FlexShrink

func (n *Node) SetFlexShrink(flexShrink float32) {
	C.YGNodeStyleSetFlexShrink(n.node, C.float(flexShrink))
}

func (n *Node) GetFlexShrink() float32 {
	return float32(C.YGNodeStyleGetFlexShrink(n.node))
}

// FlexBasis

func (n *Node) SetFlexBasis(flexBasis float32) {
	C.YGNodeStyleSetFlexBasis(n.node, C.float(flexBasis))
}

func (n *Node) SetFlexBasisPercent(flexBasis float32) {
	C.YGNodeStyleSetFlexBasisPercent(n.node, C.float(flexBasis))
}

func (n *Node) SetFlexBasisAuto() {
	C.YGNodeStyleSetFlexBasisAuto(n.node)
}

func (n *Node) SetFlexBasisMaxContent() {
	C.YGNodeStyleSetFlexBasisMaxContent(n.node)
}

func (n *Node) SetFlexBasisFitContent() {
	C.YGNodeStyleSetFlexBasisFitContent(n.node)
}

func (n *Node) SetFlexBasisStretch() {
	C.YGNodeStyleSetFlexBasisStretch(n.node)
}

func (n *Node) GetFlexBasis() Value {
	v := C.YGNodeStyleGetFlexBasis(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// Position

func (n *Node) SetPosition(edge int, position float32) {
	C.YGNodeStyleSetPosition(n.node, C.YGEdge(edge), C.float(position))
}

func (n *Node) SetPositionPercent(edge int, position float32) {
	C.YGNodeStyleSetPositionPercent(n.node, C.YGEdge(edge), C.float(position))
}

func (n *Node) SetPositionAuto(edge int) {
	C.YGNodeStyleSetPositionAuto(n.node, C.YGEdge(edge))
}

func (n *Node) GetPosition(edge int) Value {
	v := C.YGNodeStyleGetPosition(n.node, C.YGEdge(edge))
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// Margin

func (n *Node) SetMargin(edge int, margin float32) {
	C.YGNodeStyleSetMargin(n.node, C.YGEdge(edge), C.float(margin))
}

func (n *Node) SetMarginPercent(edge int, margin float32) {
	C.YGNodeStyleSetMarginPercent(n.node, C.YGEdge(edge), C.float(margin))
}

func (n *Node) SetMarginAuto(edge int) {
	C.YGNodeStyleSetMarginAuto(n.node, C.YGEdge(edge))
}

func (n *Node) GetMargin(edge int) Value {
	v := C.YGNodeStyleGetMargin(n.node, C.YGEdge(edge))
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// Padding

func (n *Node) SetPadding(edge int, padding float32) {
	C.YGNodeStyleSetPadding(n.node, C.YGEdge(edge), C.float(padding))
}

func (n *Node) SetPaddingPercent(edge int, padding float32) {
	C.YGNodeStyleSetPaddingPercent(n.node, C.YGEdge(edge), C.float(padding))
}

func (n *Node) GetPadding(edge int) Value {
	v := C.YGNodeStyleGetPadding(n.node, C.YGEdge(edge))
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// Border

func (n *Node) SetBorder(edge int, border float32) {
	C.YGNodeStyleSetBorder(n.node, C.YGEdge(edge), C.float(border))
}

func (n *Node) GetBorder(edge int) float32 {
	return float32(C.YGNodeStyleGetBorder(n.node, C.YGEdge(edge)))
}

// Gap

func (n *Node) SetGap(gutter int, gap float32) {
	C.YGNodeStyleSetGap(n.node, C.YGGutter(gutter), C.float(gap))
}

func (n *Node) SetGapPercent(gutter int, gap float32) {
	C.YGNodeStyleSetGapPercent(n.node, C.YGGutter(gutter), C.float(gap))
}

func (n *Node) GetGap(gutter int) Value {
	v := C.YGNodeStyleGetGap(n.node, C.YGGutter(gutter))
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// BoxSizing

func (n *Node) SetBoxSizing(boxSizing int) {
	C.YGNodeStyleSetBoxSizing(n.node, C.YGBoxSizing(boxSizing))
}

func (n *Node) GetBoxSizing() int {
	return int(C.YGNodeStyleGetBoxSizing(n.node))
}

// Width

func (n *Node) SetWidth(width float32) {
	C.YGNodeStyleSetWidth(n.node, C.float(width))
}

func (n *Node) SetWidthPercent(width float32) {
	C.YGNodeStyleSetWidthPercent(n.node, C.float(width))
}

func (n *Node) SetWidthAuto() {
	C.YGNodeStyleSetWidthAuto(n.node)
}

func (n *Node) SetWidthMaxContent() {
	C.YGNodeStyleSetWidthMaxContent(n.node)
}

func (n *Node) SetWidthFitContent() {
	C.YGNodeStyleSetWidthFitContent(n.node)
}

func (n *Node) SetWidthStretch() {
	C.YGNodeStyleSetWidthStretch(n.node)
}

func (n *Node) GetWidth() Value {
	v := C.YGNodeStyleGetWidth(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// Height

func (n *Node) SetHeight(height float32) {
	C.YGNodeStyleSetHeight(n.node, C.float(height))
}

func (n *Node) SetHeightPercent(height float32) {
	C.YGNodeStyleSetHeightPercent(n.node, C.float(height))
}

func (n *Node) SetHeightAuto() {
	C.YGNodeStyleSetHeightAuto(n.node)
}

func (n *Node) SetHeightMaxContent() {
	C.YGNodeStyleSetHeightMaxContent(n.node)
}

func (n *Node) SetHeightFitContent() {
	C.YGNodeStyleSetHeightFitContent(n.node)
}

func (n *Node) SetHeightStretch() {
	C.YGNodeStyleSetHeightStretch(n.node)
}

func (n *Node) GetHeight() Value {
	v := C.YGNodeStyleGetHeight(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// MinWidth

func (n *Node) SetMinWidth(minWidth float32) {
	C.YGNodeStyleSetMinWidth(n.node, C.float(minWidth))
}

func (n *Node) SetMinWidthPercent(minWidth float32) {
	C.YGNodeStyleSetMinWidthPercent(n.node, C.float(minWidth))
}

func (n *Node) SetMinWidthMaxContent() {
	C.YGNodeStyleSetMinWidthMaxContent(n.node)
}

func (n *Node) SetMinWidthFitContent() {
	C.YGNodeStyleSetMinWidthFitContent(n.node)
}

func (n *Node) SetMinWidthStretch() {
	C.YGNodeStyleSetMinWidthStretch(n.node)
}

func (n *Node) GetMinWidth() Value {
	v := C.YGNodeStyleGetMinWidth(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// MinHeight

func (n *Node) SetMinHeight(minHeight float32) {
	C.YGNodeStyleSetMinHeight(n.node, C.float(minHeight))
}

func (n *Node) SetMinHeightPercent(minHeight float32) {
	C.YGNodeStyleSetMinHeightPercent(n.node, C.float(minHeight))
}

func (n *Node) SetMinHeightMaxContent() {
	C.YGNodeStyleSetMinHeightMaxContent(n.node)
}

func (n *Node) SetMinHeightFitContent() {
	C.YGNodeStyleSetMinHeightFitContent(n.node)
}

func (n *Node) SetMinHeightStretch() {
	C.YGNodeStyleSetMinHeightStretch(n.node)
}

func (n *Node) GetMinHeight() Value {
	v := C.YGNodeStyleGetMinHeight(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// MaxWidth

func (n *Node) SetMaxWidth(maxWidth float32) {
	C.YGNodeStyleSetMaxWidth(n.node, C.float(maxWidth))
}

func (n *Node) SetMaxWidthPercent(maxWidth float32) {
	C.YGNodeStyleSetMaxWidthPercent(n.node, C.float(maxWidth))
}

func (n *Node) SetMaxWidthMaxContent() {
	C.YGNodeStyleSetMaxWidthMaxContent(n.node)
}

func (n *Node) SetMaxWidthFitContent() {
	C.YGNodeStyleSetMaxWidthFitContent(n.node)
}

func (n *Node) SetMaxWidthStretch() {
	C.YGNodeStyleSetMaxWidthStretch(n.node)
}

func (n *Node) GetMaxWidth() Value {
	v := C.YGNodeStyleGetMaxWidth(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// MaxHeight

func (n *Node) SetMaxHeight(maxHeight float32) {
	C.YGNodeStyleSetMaxHeight(n.node, C.float(maxHeight))
}

func (n *Node) SetMaxHeightPercent(maxHeight float32) {
	C.YGNodeStyleSetMaxHeightPercent(n.node, C.float(maxHeight))
}

func (n *Node) SetMaxHeightMaxContent() {
	C.YGNodeStyleSetMaxHeightMaxContent(n.node)
}

func (n *Node) SetMaxHeightFitContent() {
	C.YGNodeStyleSetMaxHeightFitContent(n.node)
}

func (n *Node) SetMaxHeightStretch() {
	C.YGNodeStyleSetMaxHeightStretch(n.node)
}

func (n *Node) GetMaxHeight() Value {
	v := C.YGNodeStyleGetMaxHeight(n.node)
	return Value{
		Value: float32(v.value),
		Unit:  Unit(v.unit),
	}
}

// AspectRatio

func (n *Node) SetAspectRatio(aspectRatio float32) {
	C.YGNodeStyleSetAspectRatio(n.node, C.float(aspectRatio))
}

func (n *Node) GetAspectRatio() float32 {
	return float32(C.YGNodeStyleGetAspectRatio(n.node))
}
