package tess

/*
#include <yoga/Yoga.h>
*/
import "C"
import (
	"fmt"
	"math"
)

type layerEdgeType int

const (
	layerEdgePadding layerEdgeType = iota
	layerEdgeMargin  layerEdgeType = iota
	layerEdgeBorder  layerEdgeType = iota
)

type Layout struct {
	node C.YGNodeRef
}

func (l *Layout) Width() float32     { return float32(C.YGNodeLayoutGetWidth(l.node)) }
func (l *Layout) Height() float32    { return float32(C.YGNodeLayoutGetHeight(l.node)) }
func (l *Layout) RawWidth() float32  { return float32(C.YGNodeLayoutGetRawWidth(l.node)) }
func (l *Layout) RawHeight() float32 { return float32(C.YGNodeLayoutGetRawHeight(l.node)) }

func (l *Layout) Top() float32    { return float32(C.YGNodeLayoutGetTop(l.node)) }
func (l *Layout) Right() float32  { return float32(C.YGNodeLayoutGetRight(l.node)) }
func (l *Layout) Bottom() float32 { return float32(C.YGNodeLayoutGetBottom(l.node)) }
func (l *Layout) Left() float32   { return float32(C.YGNodeLayoutGetLeft(l.node)) }

func (l *Layout) Padding() *LayoutEdges { return &LayoutEdges{layerEdgePadding, l.node} }
func (l *Layout) Margin() *LayoutEdges  { return &LayoutEdges{layerEdgeMargin, l.node} }
func (l *Layout) Border() *LayoutEdges  { return &LayoutEdges{layerEdgeBorder, l.node} }

func (l *Layout) Direction() DirectionType {
	return fromYGDirection(C.YGNodeLayoutGetDirection(l.node))
}

func (l *Layout) HadOverflow() bool {
	return bool(C.YGNodeLayoutGetHadOverflow(l.node))
}

type LayoutEdges struct {
	typ  layerEdgeType
	node C.YGNodeRef
}

func (e *LayoutEdges) Top() float32 {
	switch e.typ {
	case layerEdgePadding:
		return float32(C.YGNodeLayoutGetPadding(e.node, C.YGEdgeTop))
	case layerEdgeMargin:
		return float32(C.YGNodeLayoutGetMargin(e.node, C.YGEdgeTop))
	case layerEdgeBorder:
		return float32(C.YGNodeLayoutGetBorder(e.node, C.YGEdgeTop))
	}

	return 0
}

func (e *LayoutEdges) Right() float32 {
	switch e.typ {
	case layerEdgePadding:
		return float32(C.YGNodeLayoutGetPadding(e.node, C.YGEdgeRight))
	case layerEdgeMargin:
		return float32(C.YGNodeLayoutGetMargin(e.node, C.YGEdgeRight))
	case layerEdgeBorder:
		return float32(C.YGNodeLayoutGetBorder(e.node, C.YGEdgeRight))
	}

	return 0
}

func (e *LayoutEdges) Bottom() float32 {
	switch e.typ {
	case layerEdgePadding:
		return float32(C.YGNodeLayoutGetPadding(e.node, C.YGEdgeBottom))
	case layerEdgeMargin:
		return float32(C.YGNodeLayoutGetMargin(e.node, C.YGEdgeBottom))
	case layerEdgeBorder:
		return float32(C.YGNodeLayoutGetBorder(e.node, C.YGEdgeBottom))
	}

	return 0
}

func (e *LayoutEdges) Left() float32 {
	switch e.typ {
	case layerEdgePadding:
		return float32(C.YGNodeLayoutGetPadding(e.node, C.YGEdgeLeft))
	case layerEdgeMargin:
		return float32(C.YGNodeLayoutGetMargin(e.node, C.YGEdgeLeft))
	case layerEdgeBorder:
		return float32(C.YGNodeLayoutGetBorder(e.node, C.YGEdgeLeft))
	}

	return 0
}

func (n *Node) GetLayout() *Layout {
	return &Layout{node: n.node}
}

func (n *Node) HasNewLayout() bool {
	return bool(C.YGNodeGetHasNewLayout(n.node))
}

func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	C.YGNodeSetHasNewLayout(n.node, C.bool(hasNewLayout))
}

func (n *Node) IsDirty() bool {
	return bool(C.YGNodeIsDirty(n.node))
}

// MarkDirty marks the node as dirty. This will cause the next layout
// computation to recalculate this node and its children.
//
// Note: this method should only be used for nodes with a custom measure function.
// Nodes without a measure function are marked dirty automatically when their
// style properties are changed.
func (n *Node) MarkDirty() {
	C.YGNodeMarkDirty(n.node)
}

type Container struct {
	Width     float32
	Height    float32
	Direction DirectionType
}

func (n *Node) ComputeLayout(container Container) error {
	w := C.float(container.Width)
	h := C.float(container.Height)
	if container.Width == 0 || math.IsNaN(float64(container.Width)) {
		w = C.float(math.NaN())
	}
	if container.Height == 0 || math.IsNaN(float64(container.Height)) {
		h = C.float(math.NaN())
	}

	d, err := toYGDirection(container.Direction)
	if err != nil {
		return fmt.Errorf("Failed to compute layout: %w", err)
	}

	C.YGNodeCalculateLayout(n.node, w, h, d)
	return nil
}
