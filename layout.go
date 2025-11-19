package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

// Layout holds the computed layout of a node
type Layout struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

// GetLayout returns the computed layout of the node
func (n *Node) GetLayout() Layout {
	return Layout{
		Left:   float32(C.YGNodeLayoutGetLeft(n.node)),
		Top:    float32(C.YGNodeLayoutGetTop(n.node)),
		Width:  float32(C.YGNodeLayoutGetWidth(n.node)),
		Height: float32(C.YGNodeLayoutGetHeight(n.node)),
	}
}

// GetLeft returns the computed left position
func (n *Node) GetLeft() float32 {
	return float32(C.YGNodeLayoutGetLeft(n.node))
}

// GetTop returns the computed top position
func (n *Node) GetTop() float32 {
	return float32(C.YGNodeLayoutGetTop(n.node))
}

// GetRight returns the computed right position
func (n *Node) GetRight() float32 {
	return float32(C.YGNodeLayoutGetRight(n.node))
}

// GetBottom returns the computed bottom position
func (n *Node) GetBottom() float32 {
	return float32(C.YGNodeLayoutGetBottom(n.node))
}

// GetLayoutWidth returns the computed width
func (n *Node) GetLayoutWidth() float32 {
	return float32(C.YGNodeLayoutGetWidth(n.node))
}

// GetLayoutHeight returns the computed height
func (n *Node) GetLayoutHeight() float32 {
	return float32(C.YGNodeLayoutGetHeight(n.node))
}

// GetLayoutDirection returns the computed direction
func (n *Node) GetLayoutDirection() int {
	return int(C.YGNodeLayoutGetDirection(n.node))
}

// GetHadOverflow returns true if the node had overflow during layout
func (n *Node) GetHadOverflow() bool {
	return bool(C.YGNodeLayoutGetHadOverflow(n.node))
}

// GetLayoutMargin returns the computed margin for the given edge
func (n *Node) GetLayoutMargin(edge int) float32 {
	return float32(C.YGNodeLayoutGetMargin(n.node, C.YGEdge(edge)))
}

// GetLayoutBorder returns the computed border for the given edge
func (n *Node) GetLayoutBorder(edge int) float32 {
	return float32(C.YGNodeLayoutGetBorder(n.node, C.YGEdge(edge)))
}

// GetLayoutPadding returns the computed padding for the given edge
func (n *Node) GetLayoutPadding(edge int) float32 {
	return float32(C.YGNodeLayoutGetPadding(n.node, C.YGEdge(edge)))
}

// GetRawWidth returns the measured width before layout rounding
func (n *Node) GetRawWidth() float32 {
	return float32(C.YGNodeLayoutGetRawWidth(n.node))
}

// GetRawHeight returns the measured height before layout rounding
func (n *Node) GetRawHeight() float32 {
	return float32(C.YGNodeLayoutGetRawHeight(n.node))
}
