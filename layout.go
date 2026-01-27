package tess

/*
#include <yoga/Yoga.h>
#include "tess_ext.h"
*/
import "C"
import (
	"fmt"
	"math"
	"sync"
)

type layerEdgeType int

const (
	layerEdgePadding layerEdgeType = iota
	layerEdgeMargin  layerEdgeType = iota
	layerEdgeBorder  layerEdgeType = iota
)

type Layout struct {
	mu   *sync.RWMutex
	node C.YGNodeRef
}

func (l *Layout) Width() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetWidth(l.node))
}
func (l *Layout) Height() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetHeight(l.node))
}
func (l *Layout) RawWidth() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetRawWidth(l.node))
}
func (l *Layout) RawHeight() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetRawHeight(l.node))
}

func (l *Layout) Top() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetTop(l.node))
}
func (l *Layout) Right() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetRight(l.node))
}
func (l *Layout) Bottom() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetBottom(l.node))
}
func (l *Layout) Left() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return float32(C.YGNodeLayoutGetLeft(l.node))
}

func (l *Layout) Padding() *LayoutEdges {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return &LayoutEdges{l.mu, layerEdgePadding, l.node}
}
func (l *Layout) Margin() *LayoutEdges {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return &LayoutEdges{l.mu, layerEdgeMargin, l.node}
}
func (l *Layout) Border() *LayoutEdges {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return &LayoutEdges{l.mu, layerEdgeBorder, l.node}
}

func (l *Layout) Direction() DirectionType {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return fromYGDirection(C.YGNodeLayoutGetDirection(l.node))
}

func (l *Layout) HadOverflow() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return bool(C.YGNodeLayoutGetHadOverflow(l.node))
}

func (l *Layout) AbsoluteTop() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var top float32

	node := &Node{node: l.node}
	for node != nil {
		top += node.GetLayout().Top()
		node = node.GetParent()
	}

	return top
}

func (l *Layout) AbsoluteLeft() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var left float32

	node := &Node{node: l.node}
	for node != nil {
		left += node.GetLayout().Left()
		node = node.GetParent()
	}

	return left
}

func (l *Layout) AbsoluteBottom() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var bottom float32

	node := &Node{node: l.node}
	for node != nil {
		bottom += node.GetLayout().Bottom()
		node = node.GetParent()
	}

	return bottom
}

func (l *Layout) AbsoluteRight() float32 {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var right float32

	node := &Node{node: l.node}
	for node != nil {
		right += node.GetLayout().Right()
		node = node.GetParent()
	}

	return right
}

type LayoutEdges struct {
	mu *sync.RWMutex

	typ  layerEdgeType
	node C.YGNodeRef
}

func (e *LayoutEdges) Top() float32 {
	e.mu.RLock()
	defer e.mu.RUnlock()

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
	e.mu.RLock()
	defer e.mu.RUnlock()

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
	e.mu.RLock()
	defer e.mu.RUnlock()

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
	e.mu.RLock()
	defer e.mu.RUnlock()

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
	n.mu.RLock()
	defer n.mu.RUnlock()

	return &Layout{mu: &n.mu, node: n.node}
}

func (n *Node) HasNewLayout() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return bool(C.YGNodeGetHasNewLayout(n.node))
}

func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	C.YGNodeSetHasNewLayout(n.node, C.bool(hasNewLayout))
}

func (n *Node) IsDirty() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return bool(C.YGNodeIsDirty(n.node))
}

// MarkDirty marks the node as dirty. This will cause the next layout
// computation to recalculate this node and its children.
//
// Note: this method should only be used for nodes with a custom measure function.
// Nodes without a measure function are marked dirty automatically when their
// style properties are changed.
func (n *Node) MarkDirty() {
	n.mu.Lock()
	defer n.mu.Unlock()

	C.YGNodeMarkDirty(n.node)
}

// SetDirty directly sets the dirty flag on a node without propagating to parent.
func (n *Node) SetDirty(dirty bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	C.YGNodeSetDirtyExt(n.node, C.bool(dirty))
}

type Container struct {
	Width     float32
	Height    float32
	Direction DirectionType
}

func (n *Node) ComputeLayout(container Container) error {
	n.lockChildren()
	n.mu.Lock()
	defer func() {
		n.mu.Unlock()
		n.unlockChildren()
	}()

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

func (n *Node) lockChildren() {
	for child := range n.Children() {
		child.lockChildren()
		child.mu.Lock()
	}
}

func (n *Node) unlockChildren() {
	for child := range n.Children() {
		child.mu.Unlock()
		child.unlockChildren()
	}
}
