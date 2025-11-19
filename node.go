package tess

/*
#include <yoga/Yoga.h>
*/
import "C"
import (
	"math"
	"unsafe"
)

// Node represents a Yoga layout node
type Node struct {
	node C.YGNodeRef
}

// NewNode creates a new Yoga node with default settings
func NewNode() *Node {
	return &Node{
		node: C.YGNodeNew(),
	}
}

// Clone creates a mutable copy of the node with the same context and children
func (n *Node) Clone() *Node {
	return &Node{
		node: C.YGNodeClone(n.node),
	}
}

// Free frees the Yoga node, disconnecting it from its owner and children
func (n *Node) Free() {
	if n.node != nil {
		C.YGNodeFree(n.node)
		n.node = nil
	}
}

// FreeRecursive frees the subtree of Yoga nodes rooted at this node
func (n *Node) FreeRecursive() {
	if n.node != nil {
		C.YGNodeFreeRecursive(n.node)
		n.node = nil
	}
}

// Reset resets the node to its default state
func (n *Node) Reset() {
	C.YGNodeReset(n.node)
}

// CalculateLayout calculates the layout of the tree rooted at this node
func (n *Node) CalculateLayout(width, height float32, direction int) {
	w := C.float(width)
	h := C.float(height)
	if math.IsNaN(float64(width)) {
		w = C.float(math.NaN())
	}
	if math.IsNaN(float64(height)) {
		h = C.float(math.NaN())
	}
	C.YGNodeCalculateLayout(n.node, w, h, C.YGDirection(direction))
}

// HasNewLayout returns whether the node may have new layout results
func (n *Node) HasNewLayout() bool {
	return bool(C.YGNodeGetHasNewLayout(n.node))
}

// SetHasNewLayout sets whether a node's layout is considered new
func (n *Node) SetHasNewLayout(hasNewLayout bool) {
	C.YGNodeSetHasNewLayout(n.node, C.bool(hasNewLayout))
}

// IsDirty returns whether the node's layout results are dirty
func (n *Node) IsDirty() bool {
	return bool(C.YGNodeIsDirty(n.node))
}

// MarkDirty marks a node with custom measure function as dirty
func (n *Node) MarkDirty() {
	C.YGNodeMarkDirty(n.node)
}

// InsertChild inserts a child node at the given index
func (n *Node) InsertChild(child *Node, index int) {
	C.YGNodeInsertChild(n.node, child.node, C.size_t(index))
}

// SwapChild replaces the child node at a given index with a new one
func (n *Node) SwapChild(child *Node, index int) {
	C.YGNodeSwapChild(n.node, child.node, C.size_t(index))
}

// RemoveChild removes the given child node
func (n *Node) RemoveChild(child *Node) {
	C.YGNodeRemoveChild(n.node, child.node)
}

// RemoveAllChildren removes all child nodes
func (n *Node) RemoveAllChildren() {
	C.YGNodeRemoveAllChildren(n.node)
}

// SetChildren sets children according to the given list of nodes
func (n *Node) SetChildren(children []*Node) {
	if len(children) == 0 {
		C.YGNodeRemoveAllChildren(n.node)
		return
	}

	// Convert Go slice to C array
	cChildren := make([]C.YGNodeRef, len(children))
	for i, child := range children {
		cChildren[i] = child.node
	}
	C.YGNodeSetChildren(n.node, &cChildren[0], C.size_t(len(children)))
}

// GetChild returns the child node at the given index
func (n *Node) GetChild(index int) *Node {
	child := C.YGNodeGetChild(n.node, C.size_t(index))
	if child == nil {
		return nil
	}
	return &Node{node: child}
}

// GetChildCount returns the number of child nodes
func (n *Node) GetChildCount() int {
	return int(C.YGNodeGetChildCount(n.node))
}

// GetParent returns the parent node
func (n *Node) GetParent() *Node {
	parent := C.YGNodeGetParent(n.node)
	if parent == nil {
		return nil
	}
	return &Node{node: parent}
}

// GetOwner returns the owner node
func (n *Node) GetOwner() *Node {
	owner := C.YGNodeGetOwner(n.node)
	if owner == nil {
		return nil
	}
	return &Node{node: owner}
}

// SetContext sets extra data on the node which may be read from during callbacks
func (n *Node) SetContext(context unsafe.Pointer) {
	C.YGNodeSetContext(n.node, context)
}

// GetContext returns the context or nil if no context has been set
func (n *Node) GetContext() unsafe.Pointer {
	return C.YGNodeGetContext(n.node)
}

// SetNodeType sets whether a leaf node's layout results may be truncated during layout rounding
func (n *Node) SetNodeType(nodeType int) {
	C.YGNodeSetNodeType(n.node, C.YGNodeType(nodeType))
}

// GetNodeType returns the node type
func (n *Node) GetNodeType() int {
	return int(C.YGNodeGetNodeType(n.node))
}

// SetAlwaysFormsContainingBlock makes this node always form a containing block for descendants
func (n *Node) SetAlwaysFormsContainingBlock(alwaysFormsContainingBlock bool) {
	C.YGNodeSetAlwaysFormsContainingBlock(n.node, C.bool(alwaysFormsContainingBlock))
}

// GetAlwaysFormsContainingBlock returns whether the node always forms a containing block
func (n *Node) GetAlwaysFormsContainingBlock() bool {
	return bool(C.YGNodeGetAlwaysFormsContainingBlock(n.node))
}
