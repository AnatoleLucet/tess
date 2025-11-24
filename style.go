package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

type Style struct {
	// Display & Layout
	Display        DisplayType
	FlexDirection  FlexDirection
	JustifyContent FlexJustify
	AlignItems     FlexAlign
	AlignSelf      FlexAlign
	AlignContent   FlexAlign
	FlexWrap       FlexWrap
	Direction      DirectionType

	// Dimensions
	Width, Height       Value
	MinWidth, MinHeight Value
	MaxWidth, MaxHeight Value
	AspectRatio         float32

	// Spacing
	Padding Edges
	Margin  Edges
	Border  Edges
	Gap     Gap

	// Flex item properties
	FlexGrow   float32
	FlexShrink float32
	FlexBasis  Value

	// Positioning
	Position                 PositionType
	Top, Right, Bottom, Left Value

	// Overflow
	Overflow OverflowType

	// Box Sizing
	BoxSizing BoxSizingType
}

func (n *Node) SetDisplay(display DisplayType) error {
	ygDisplay, err := toYGDisplay(display)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetDisplay(n.node, ygDisplay)
	return nil
}

func (n *Node) SetFlexDirection(direction FlexDirection) error {
	ygDirection, err := toYGFlexDirection(direction)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetFlexDirection(n.node, ygDirection)
	return nil
}

func (n *Node) SetJustifyContent(justify FlexJustify) error {
	ygJustify, err := toYGJustify(justify)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetJustifyContent(n.node, ygJustify)
	return nil
}

func (n *Node) SetAlignItems(align FlexAlign) error {
	ygAlign, err := toYGAlign(align)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetAlignItems(n.node, ygAlign)
	return nil
}

func (n *Node) SetAlignSelf(align FlexAlign) error {
	ygAlign, err := toYGAlign(align)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetAlignSelf(n.node, ygAlign)
	return nil
}

func (n *Node) SetAlignContent(align FlexAlign) error {
	ygAlign, err := toYGAlign(align)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetAlignContent(n.node, ygAlign)
	return nil
}

func (n *Node) SetFlexWrap(wrap FlexWrap) error {
	ygWrap, err := toYGWrap(wrap)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetFlexWrap(n.node, ygWrap)
	return nil
}

func (n *Node) SetDirection(direction DirectionType) error {
	ygDirection, err := toYGDirection(direction)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetDirection(n.node, ygDirection)
	return nil
}

func (n *Node) SetWidth(width Value) error {
	switch width.unit {
	case UnitPoint:
		C.YGNodeStyleSetWidth(n.node, C.float(width.value))
	case UnitPercent:
		C.YGNodeStyleSetWidthPercent(n.node, C.float(width.value))
	case UnitAuto:
		C.YGNodeStyleSetWidthAuto(n.node)
	case UnitMaxContent:
		C.YGNodeStyleSetWidthMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetWidthFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetWidthStretch(n.node)
	}
	return nil
}

func (n *Node) SetHeight(height Value) error {
	switch height.unit {
	case UnitPoint:
		C.YGNodeStyleSetHeight(n.node, C.float(height.value))
	case UnitPercent:
		C.YGNodeStyleSetHeightPercent(n.node, C.float(height.value))
	case UnitAuto:
		C.YGNodeStyleSetHeightAuto(n.node)
	case UnitMaxContent:
		C.YGNodeStyleSetHeightMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetHeightFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetHeightStretch(n.node)
	}
	return nil
}

func (n *Node) SetMinWidth(minWidth Value) error {
	switch minWidth.unit {
	case UnitPoint:
		C.YGNodeStyleSetMinWidth(n.node, C.float(minWidth.value))
	case UnitPercent:
		C.YGNodeStyleSetMinWidthPercent(n.node, C.float(minWidth.value))
	case UnitMaxContent:
		C.YGNodeStyleSetMinWidthMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetMinWidthFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetMinWidthStretch(n.node)
	}
	return nil
}

func (n *Node) SetMinHeight(minHeight Value) error {
	switch minHeight.unit {
	case UnitPoint:
		C.YGNodeStyleSetMinHeight(n.node, C.float(minHeight.value))
	case UnitPercent:
		C.YGNodeStyleSetMinHeightPercent(n.node, C.float(minHeight.value))
	case UnitMaxContent:
		C.YGNodeStyleSetMinHeightMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetMinHeightFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetMinHeightStretch(n.node)
	}
	return nil
}

func (n *Node) SetMaxWidth(maxWidth Value) error {
	switch maxWidth.unit {
	case UnitPoint:
		C.YGNodeStyleSetMaxWidth(n.node, C.float(maxWidth.value))
	case UnitPercent:
		C.YGNodeStyleSetMaxWidthPercent(n.node, C.float(maxWidth.value))
	case UnitMaxContent:
		C.YGNodeStyleSetMaxWidthMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetMaxWidthFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetMaxWidthStretch(n.node)
	}
	return nil
}

func (n *Node) SetMaxHeight(maxHeight Value) error {
	switch maxHeight.unit {
	case UnitPoint:
		C.YGNodeStyleSetMaxHeight(n.node, C.float(maxHeight.value))
	case UnitPercent:
		C.YGNodeStyleSetMaxHeightPercent(n.node, C.float(maxHeight.value))
	case UnitMaxContent:
		C.YGNodeStyleSetMaxHeightMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetMaxHeightFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetMaxHeightStretch(n.node)
	}
	return nil
}

func (n *Node) SetAspectRatio(aspectRatio float32) error {
	C.YGNodeStyleSetAspectRatio(n.node, C.float(aspectRatio))
	return nil
}

func (n *Node) SetFlexGrow(flexGrow float32) error {
	C.YGNodeStyleSetFlexGrow(n.node, C.float(flexGrow))
	return nil
}

func (n *Node) SetFlexShrink(flexShrink float32) error {
	C.YGNodeStyleSetFlexShrink(n.node, C.float(flexShrink))
	return nil
}

func (n *Node) SetFlexBasis(flexBasis Value) error {
	switch flexBasis.unit {
	case UnitPoint:
		C.YGNodeStyleSetFlexBasis(n.node, C.float(flexBasis.value))
	case UnitPercent:
		C.YGNodeStyleSetFlexBasisPercent(n.node, C.float(flexBasis.value))
	case UnitAuto:
		C.YGNodeStyleSetFlexBasisAuto(n.node)
	case UnitMaxContent:
		C.YGNodeStyleSetFlexBasisMaxContent(n.node)
	case UnitFitContent:
		C.YGNodeStyleSetFlexBasisFitContent(n.node)
	case UnitStretch:
		C.YGNodeStyleSetFlexBasisStretch(n.node)
	}
	return nil
}

func (n *Node) SetPosition(position PositionType) error {
	ygPosition, err := toYGPositionType(position)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetPositionType(n.node, ygPosition)
	return nil
}

func (n *Node) SetTop(top Value) error {
	switch top.unit {
	case UnitPoint:
		C.YGNodeStyleSetPosition(n.node, C.YGEdgeTop, C.float(top.value))
	case UnitPercent:
		C.YGNodeStyleSetPositionPercent(n.node, C.YGEdgeTop, C.float(top.value))
	case UnitAuto:
		C.YGNodeStyleSetPositionAuto(n.node, C.YGEdgeTop)
	}
	return nil
}

func (n *Node) SetRight(right Value) error {
	switch right.unit {
	case UnitPoint:
		C.YGNodeStyleSetPosition(n.node, C.YGEdgeRight, C.float(right.value))
	case UnitPercent:
		C.YGNodeStyleSetPositionPercent(n.node, C.YGEdgeRight, C.float(right.value))
	case UnitAuto:
		C.YGNodeStyleSetPositionAuto(n.node, C.YGEdgeRight)
	}
	return nil
}

func (n *Node) SetBottom(bottom Value) error {
	switch bottom.unit {
	case UnitPoint:
		C.YGNodeStyleSetPosition(n.node, C.YGEdgeBottom, C.float(bottom.value))
	case UnitPercent:
		C.YGNodeStyleSetPositionPercent(n.node, C.YGEdgeBottom, C.float(bottom.value))
	case UnitAuto:
		C.YGNodeStyleSetPositionAuto(n.node, C.YGEdgeBottom)
	}
	return nil
}

func (n *Node) SetLeft(left Value) error {
	switch left.unit {
	case UnitPoint:
		C.YGNodeStyleSetPosition(n.node, C.YGEdgeLeft, C.float(left.value))
	case UnitPercent:
		C.YGNodeStyleSetPositionPercent(n.node, C.YGEdgeLeft, C.float(left.value))
	case UnitAuto:
		C.YGNodeStyleSetPositionAuto(n.node, C.YGEdgeLeft)
	}
	return nil
}

func (n *Node) SetOverflow(overflow OverflowType) error {
	ygOverflow, err := toYGOverflow(overflow)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetOverflow(n.node, ygOverflow)
	return nil
}

func (n *Node) SetBoxSizing(boxSizing BoxSizingType) error {
	ygBoxSizing, err := toYGBoxSizing(boxSizing)
	if err != nil {
		return err
	}
	C.YGNodeStyleSetBoxSizing(n.node, ygBoxSizing)
	return nil
}

// SetPadding sets the padding for the node.
// Note: Padding can only be set in points or percent.
func (n *Node) SetPadding(edges Edges) error {
	setPaddingEdge := func(edge C.YGEdge, value Value) {
		switch value.unit {
		case UnitPoint:
			C.YGNodeStyleSetPadding(n.node, edge, C.float(value.value))
		case UnitPercent:
			C.YGNodeStyleSetPaddingPercent(n.node, edge, C.float(value.value))
		}
	}

	if edges.All.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeTop, edges.All)
		setPaddingEdge(C.YGEdgeBottom, edges.All)
		setPaddingEdge(C.YGEdgeLeft, edges.All)
		setPaddingEdge(C.YGEdgeRight, edges.All)
		setPaddingEdge(C.YGEdgeStart, edges.All)
		setPaddingEdge(C.YGEdgeEnd, edges.All)
	}
	if edges.Horizontal.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeLeft, edges.Horizontal)
		setPaddingEdge(C.YGEdgeRight, edges.Horizontal)
		setPaddingEdge(C.YGEdgeStart, edges.Horizontal)
		setPaddingEdge(C.YGEdgeEnd, edges.Horizontal)
	}
	if edges.Vertical.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeTop, edges.Vertical)
		setPaddingEdge(C.YGEdgeBottom, edges.Vertical)
	}
	if edges.Start.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeStart, edges.Start)
	}
	if edges.End.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeEnd, edges.End)
	}
	if edges.Top.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeTop, edges.Top)
	}
	if edges.Bottom.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeBottom, edges.Bottom)
	}
	if edges.Left.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeLeft, edges.Left)
	}
	if edges.Right.unit != UnitUndefined {
		setPaddingEdge(C.YGEdgeRight, edges.Right)
	}

	return nil
}

// SetMargin sets the margin widths for the node.
// Note: Margins can be set in points, percent, or auto.
func (n *Node) SetMargin(edges Edges) error {
	setMarginEdge := func(edge C.YGEdge, value Value) {
		switch value.unit {
		case UnitPoint:
			C.YGNodeStyleSetMargin(n.node, edge, C.float(value.value))
		case UnitPercent:
			C.YGNodeStyleSetMarginPercent(n.node, edge, C.float(value.value))
		case UnitAuto:
			C.YGNodeStyleSetMarginAuto(n.node, edge)
		}
	}

	if edges.All.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeTop, edges.All)
		setMarginEdge(C.YGEdgeBottom, edges.All)
		setMarginEdge(C.YGEdgeLeft, edges.All)
		setMarginEdge(C.YGEdgeRight, edges.All)
		setMarginEdge(C.YGEdgeStart, edges.All)
		setMarginEdge(C.YGEdgeEnd, edges.All)
	}
	if edges.Horizontal.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeLeft, edges.Horizontal)
		setMarginEdge(C.YGEdgeRight, edges.Horizontal)
		setMarginEdge(C.YGEdgeStart, edges.Horizontal)
		setMarginEdge(C.YGEdgeEnd, edges.Horizontal)
	}
	if edges.Vertical.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeTop, edges.Vertical)
		setMarginEdge(C.YGEdgeBottom, edges.Vertical)
	}
	if edges.Start.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeStart, edges.Start)
	}
	if edges.End.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeEnd, edges.End)
	}
	if edges.Top.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeTop, edges.Top)
	}
	if edges.Bottom.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeBottom, edges.Bottom)
	}
	if edges.Left.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeLeft, edges.Left)
	}
	if edges.Right.unit != UnitUndefined {
		setMarginEdge(C.YGEdgeRight, edges.Right)
	}

	return nil
}

// SetBorder sets the border widths for the node.
// Note: Borders can only be set in points.
func (n *Node) SetBorder(edges Edges) error {
	setBorderEdge := func(edge C.YGEdge, value Value) {
		if value.unit == UnitPoint {
			C.YGNodeStyleSetBorder(n.node, edge, C.float(value.value))
		}
	}

	if edges.All.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeTop, edges.All)
		setBorderEdge(C.YGEdgeBottom, edges.All)
		setBorderEdge(C.YGEdgeLeft, edges.All)
		setBorderEdge(C.YGEdgeRight, edges.All)
		setBorderEdge(C.YGEdgeStart, edges.All)
		setBorderEdge(C.YGEdgeEnd, edges.All)
	}
	if edges.Horizontal.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeLeft, edges.Horizontal)
		setBorderEdge(C.YGEdgeRight, edges.Horizontal)
		setBorderEdge(C.YGEdgeStart, edges.Horizontal)
		setBorderEdge(C.YGEdgeEnd, edges.Horizontal)
	}
	if edges.Vertical.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeTop, edges.Vertical)
		setBorderEdge(C.YGEdgeBottom, edges.Vertical)
	}
	if edges.Start.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeStart, edges.Start)
	}
	if edges.End.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeEnd, edges.End)
	}
	if edges.Top.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeTop, edges.Top)
	}
	if edges.Bottom.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeBottom, edges.Bottom)
	}
	if edges.Left.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeLeft, edges.Left)
	}
	if edges.Right.unit != UnitUndefined {
		setBorderEdge(C.YGEdgeRight, edges.Right)
	}

	return nil
}

// SetGap sets the gap sizes for the node.
// Note: Gaps can be set in points or percent.
func (n *Node) SetGap(gap Gap) error {
	setGapGutter := func(gutter C.YGGutter, value Value) {
		switch value.unit {
		case UnitPoint:
			C.YGNodeStyleSetGap(n.node, gutter, C.float(value.value))
		case UnitPercent:
			C.YGNodeStyleSetGapPercent(n.node, gutter, C.float(value.value))
		}
	}

	if gap.All.unit != UnitUndefined {
		setGapGutter(C.YGGutterAll, gap.All)
	}
	if gap.Row.unit != UnitUndefined {
		setGapGutter(C.YGGutterRow, gap.Row)
	}
	if gap.Column.unit != UnitUndefined {
		setGapGutter(C.YGGutterColumn, gap.Column)
	}

	return nil
}

func (n *Node) GetDisplay() DisplayType {
	return fromYGDisplay(C.YGNodeStyleGetDisplay(n.node))
}

func (n *Node) GetFlexDirection() FlexDirection {
	return fromYGFlexDirection(C.YGNodeStyleGetFlexDirection(n.node))
}

func (n *Node) GetJustifyContent() FlexJustify {
	return fromYGJustify(C.YGNodeStyleGetJustifyContent(n.node))
}

func (n *Node) GetAlignItems() FlexAlign {
	return fromYGAlign(C.YGNodeStyleGetAlignItems(n.node))
}

func (n *Node) GetAlignSelf() FlexAlign {
	return fromYGAlign(C.YGNodeStyleGetAlignSelf(n.node))
}

func (n *Node) GetAlignContent() FlexAlign {
	return fromYGAlign(C.YGNodeStyleGetAlignContent(n.node))
}

func (n *Node) GetFlexWrap() FlexWrap {
	return fromYGWrap(C.YGNodeStyleGetFlexWrap(n.node))
}

func (n *Node) GetDirection() DirectionType {
	return fromYGDirection(C.YGNodeStyleGetDirection(n.node))
}

func (n *Node) GetWidth() Value {
	return fromYGValue(C.YGNodeStyleGetWidth(n.node))
}

func (n *Node) GetHeight() Value {
	return fromYGValue(C.YGNodeStyleGetHeight(n.node))
}

func (n *Node) GetMinWidth() Value {
	return fromYGValue(C.YGNodeStyleGetMinWidth(n.node))
}

func (n *Node) GetMinHeight() Value {
	return fromYGValue(C.YGNodeStyleGetMinHeight(n.node))
}

func (n *Node) GetMaxWidth() Value {
	return fromYGValue(C.YGNodeStyleGetMaxWidth(n.node))
}

func (n *Node) GetMaxHeight() Value {
	return fromYGValue(C.YGNodeStyleGetMaxHeight(n.node))
}

func (n *Node) GetAspectRatio() float32 {
	return float32(C.YGNodeStyleGetAspectRatio(n.node))
}

func (n *Node) GetFlexGrow() float32 {
	return float32(C.YGNodeStyleGetFlexGrow(n.node))
}

func (n *Node) GetFlexShrink() float32 {
	return float32(C.YGNodeStyleGetFlexShrink(n.node))
}

func (n *Node) GetFlexBasis() Value {
	return fromYGValue(C.YGNodeStyleGetFlexBasis(n.node))
}

func (n *Node) GetPosition() PositionType {
	return fromYGPositionType(C.YGNodeStyleGetPositionType(n.node))
}

func (n *Node) GetTop() Value {
	return fromYGValue(C.YGNodeStyleGetPosition(n.node, C.YGEdgeTop))
}

func (n *Node) GetRight() Value {
	return fromYGValue(C.YGNodeStyleGetPosition(n.node, C.YGEdgeRight))
}

func (n *Node) GetBottom() Value {
	return fromYGValue(C.YGNodeStyleGetPosition(n.node, C.YGEdgeBottom))
}

func (n *Node) GetLeft() Value {
	return fromYGValue(C.YGNodeStyleGetPosition(n.node, C.YGEdgeLeft))
}

func (n *Node) GetOverflow() OverflowType {
	return fromYGOverflow(C.YGNodeStyleGetOverflow(n.node))
}

func (n *Node) GetBoxSizing() BoxSizingType {
	return fromYGBoxSizing(C.YGNodeStyleGetBoxSizing(n.node))
}

func (n *Node) GetPadding() *StyleEdges {
	return &StyleEdges{styleEdgePadding, n.node}
}

func (n *Node) GetMargin() *StyleEdges {
	return &StyleEdges{styleEdgeMargin, n.node}
}

func (n *Node) GetBorder() *StyleEdges {
	return &StyleEdges{styleEdgeBorder, n.node}
}

func (n *Node) GetGap() *StyleGap {
	return &StyleGap{n.node}
}

// Apply applies a Style to the node
func (n *Node) Apply(style *Style) error {
	if err := n.SetDisplay(style.Display); err != nil {
		return err
	}
	if err := n.SetFlexDirection(style.FlexDirection); err != nil {
		return err
	}
	if err := n.SetJustifyContent(style.JustifyContent); err != nil {
		return err
	}
	if err := n.SetAlignItems(style.AlignItems); err != nil {
		return err
	}
	if err := n.SetAlignSelf(style.AlignSelf); err != nil {
		return err
	}
	if err := n.SetAlignContent(style.AlignContent); err != nil {
		return err
	}
	if err := n.SetFlexWrap(style.FlexWrap); err != nil {
		return err
	}
	if err := n.SetDirection(style.Direction); err != nil {
		return err
	}

	if style.Width.unit != UnitUndefined {
		if err := n.SetWidth(style.Width); err != nil {
			return err
		}
	}
	if style.Height.unit != UnitUndefined {
		if err := n.SetHeight(style.Height); err != nil {
			return err
		}
	}
	if style.MinWidth.unit != UnitUndefined {
		if err := n.SetMinWidth(style.MinWidth); err != nil {
			return err
		}
	}
	if style.MinHeight.unit != UnitUndefined {
		if err := n.SetMinHeight(style.MinHeight); err != nil {
			return err
		}
	}
	if style.MaxWidth.unit != UnitUndefined {
		if err := n.SetMaxWidth(style.MaxWidth); err != nil {
			return err
		}
	}
	if style.MaxHeight.unit != UnitUndefined {
		if err := n.SetMaxHeight(style.MaxHeight); err != nil {
			return err
		}
	}
	if style.AspectRatio != 0 {
		if err := n.SetAspectRatio(style.AspectRatio); err != nil {
			return err
		}
	}

	if err := n.SetPadding(style.Padding); err != nil {
		return err
	}
	if err := n.SetMargin(style.Margin); err != nil {
		return err
	}
	if err := n.SetBorder(style.Border); err != nil {
		return err
	}
	if err := n.SetGap(style.Gap); err != nil {
		return err
	}

	if err := n.SetFlexGrow(style.FlexGrow); err != nil {
		return err
	}
	if err := n.SetFlexShrink(style.FlexShrink); err != nil {
		return err
	}
	if style.FlexBasis.unit != UnitUndefined {
		if err := n.SetFlexBasis(style.FlexBasis); err != nil {
			return err
		}
	}

	if err := n.SetPosition(style.Position); err != nil {
		return err
	}
	if style.Top.unit != UnitUndefined {
		if err := n.SetTop(style.Top); err != nil {
			return err
		}
	}
	if style.Right.unit != UnitUndefined {
		if err := n.SetRight(style.Right); err != nil {
			return err
		}
	}
	if style.Bottom.unit != UnitUndefined {
		if err := n.SetBottom(style.Bottom); err != nil {
			return err
		}
	}
	if style.Left.unit != UnitUndefined {
		if err := n.SetLeft(style.Left); err != nil {
			return err
		}
	}

	if err := n.SetOverflow(style.Overflow); err != nil {
		return err
	}
	if err := n.SetBoxSizing(style.BoxSizing); err != nil {
		return err
	}

	return nil
}

type StyleGap struct {
	node C.YGNodeRef
}

func (g *StyleGap) GetRow() Value {
	return fromYGValue(C.YGNodeStyleGetGap(g.node, C.YGGutterRow))
}

func (g *StyleGap) GetColumn() Value {
	return fromYGValue(C.YGNodeStyleGetGap(g.node, C.YGGutterColumn))
}

func (g *StyleGap) GetAll() Value {
	return fromYGValue(C.YGNodeStyleGetGap(g.node, C.YGGutterAll))
}

type styleEdgeType int

const (
	styleEdgePadding styleEdgeType = iota
	styleEdgeMargin
	styleEdgeBorder
)

type StyleEdges struct {
	typ  styleEdgeType
	node C.YGNodeRef
}

func (e *StyleEdges) GetTop() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeTop))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeTop))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeTop))}
	}

	return Undefined()
}

func (e *StyleEdges) GetRight() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeRight))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeRight))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeRight))}
	}

	return Undefined()
}

func (e *StyleEdges) GetBottom() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeBottom))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeBottom))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeBottom))}
	}

	return Undefined()
}

func (e *StyleEdges) GetLeft() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeLeft))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeLeft))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeLeft))}
	}

	return Undefined()
}

func (e *StyleEdges) GetStart() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeStart))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeStart))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeStart))}
	}

	return Undefined()
}

func (e *StyleEdges) GetEnd() Value {
	switch e.typ {
	case styleEdgePadding:
		return fromYGValue(C.YGNodeStyleGetPadding(e.node, C.YGEdgeEnd))
	case styleEdgeMargin:
		return fromYGValue(C.YGNodeStyleGetMargin(e.node, C.YGEdgeEnd))
	case styleEdgeBorder:
		return Value{unit: UnitPoint, value: float32(C.YGNodeStyleGetBorder(e.node, C.YGEdgeEnd))}
	}
	return Undefined()
}
