package tess

/*
#include <yoga/Yoga.h>
*/
import "C"

type NodeType int

const (
	NodeTypeDefault NodeType = iota
	NodeTypeText
)

func (t NodeType) String() string {
	switch t {
	case NodeTypeDefault:
		return "default"
	case NodeTypeText:
		return "text"
	}

	return "unknown"
}

type Node struct {
	node   C.YGNodeRef
	config *Config
}

func NewNode(styles ...*Style) *Node {
	config := getDefaultConfig()

	node := &Node{
		node:   C.YGNodeNewWithConfig(config.config),
		config: config,
	}
	for _, style := range styles {
		node.Apply(style)
	}

	return node
}

func (n *Node) Clone() *Node {
	return &Node{node: C.YGNodeClone(n.node)}
}

func (n *Node) Free() {
	if n.node != nil {
		C.YGNodeFree(n.node)
		n.node = nil
	}
}

func (n *Node) FreeRecursive() {
	if n.node != nil {
		C.YGNodeFreeRecursive(n.node)
		n.node = nil
	}
}

func (n *Node) Reset() {
	C.YGNodeReset(n.node)
}

func (n *Node) GetNodeType() NodeType {
	return fromYGNodeType(C.YGNodeGetNodeType(n.node))
}

func (n *Node) SetNodeType(nodeType NodeType) {
	ygNodeType, _ := toYGNodeType(nodeType)
	C.YGNodeSetNodeType(n.node, ygNodeType)
}

func (n *Node) GetChildCount() int {
	return int(C.YGNodeGetChildCount(n.node))
}

func (n *Node) GetChild(index int) *Node {
	child := C.YGNodeGetChild(n.node, C.size_t(index))
	if child == nil {
		return nil
	}

	return &Node{node: child}
}

func (n *Node) AddChild(child *Node) {
	count := n.GetChildCount()
	C.YGNodeInsertChild(n.node, child.node, C.size_t(count))
}

func (n *Node) SetChildren(children []*Node) {
	if len(children) == 0 {
		C.YGNodeRemoveAllChildren(n.node)
		return
	}

	cChildren := make([]C.YGNodeRef, len(children))
	for i, child := range children {
		cChildren[i] = child.node
	}

	C.YGNodeSetChildren(n.node, &cChildren[0], C.size_t(len(children)))
}

func (n *Node) InsertChild(child *Node, index int) {
	C.YGNodeInsertChild(n.node, child.node, C.size_t(index))
}

func (n *Node) SwapChild(child *Node, index int) {
	C.YGNodeSwapChild(n.node, child.node, C.size_t(index))
}

func (n *Node) RemoveChild(child *Node) {
	C.YGNodeRemoveChild(n.node, child.node)
}

func (n *Node) RemoveAllChildren() {
	C.YGNodeRemoveAllChildren(n.node)
}

func (n *Node) GetParent() *Node {
	parent := C.YGNodeGetParent(n.node)
	if parent == nil {
		return nil
	}
	return &Node{node: parent}
}

func (n *Node) SetConfig(config *Config) {
	n.config = config
	C.YGNodeSetConfig(n.node, config.config)
}

func (n *Node) GetConfig() *Config {
	return n.config
}
