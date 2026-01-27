package tess

/*
#include <yoga/Yoga.h>
*/
import "C"
import (
	"iter"
	"slices"
	"sync"
)

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
	mu sync.RWMutex

	node   C.YGNodeRef
	config *Config

	parent   *Node
	children []*Node

	layout *Layout
}

func NewNode(styles ...*Style) (*Node, error) {
	config := getDefaultConfig()
	ygnode := C.YGNodeNewWithConfig(config.config)

	node := newNode(config, ygnode)
	for _, style := range styles {
		if err := node.Apply(style); err != nil {
			return nil, err
		}
	}

	return node, nil
}

func newNode(config *Config, ygnode C.YGNodeRef) *Node {
	node := &Node{node: ygnode, config: config}
	node.layout = newLayout(node)
	return node
}

func (n *Node) Clone() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	clone := newNode(n.config, C.YGNodeClone(n.node))

	if n.HasMeasureFunc() {
		clone.reregisterMeasureFunc(n.getMeasureFunc())
	} else {
		clone.clearContext()
	}

	return clone
}

func (n *Node) CloneRecursive() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	clone := newNode(n.config, C.YGNodeClone(n.node))
	clone.RemoveAllChildren()

	if n.HasMeasureFunc() {
		clone.reregisterMeasureFunc(n.getMeasureFunc())
	} else {
		clone.clearContext()
	}

	for i := 0; i < n.GetChildCount(); i++ {
		child := n.GetChild(i)
		clonedChild := child.CloneRecursive()
		clonedChild.SetParent(clone)
		clone.AppendChild(clonedChild)
	}

	return clone
}

// Snapshot creates a deep clone of the node tree while preserving the dirty state.
// Unlike CloneRecursive which marks cloned nodes as dirty, Snapshot maintains
// the original dirty status, allowing cached layout computations to be reused.
func (n *Node) Snapshot() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.snapshot(n.IsDirty())
}

func (n *Node) snapshot(parentWasDirty bool) *Node {
	wasDirty := n.IsDirty()

	clone := newNode(n.config, C.YGNodeClone(n.node))
	clone.RemoveAllChildren()

	if n.HasMeasureFunc() {
		clone.reregisterMeasureFunc(n.getMeasureFunc())
	} else {
		clone.clearContext()
	}

	for i := 0; i < n.GetChildCount(); i++ {
		child := n.GetChild(i)
		clonedChild := child.snapshot(wasDirty)
		clonedChild.SetParent(clone)
		clone.AppendChild(clonedChild)
	}

	if !wasDirty && !parentWasDirty {
		clone.SetDirty(false)
	}

	return clone
}

func (n *Node) Free() {
	if n.node != nil {
		if n.HasMeasureFunc() {
			n.UnsetMeasureFunc()
		}

		n.mu.Lock()
		C.YGNodeFree(n.node)
		n.node = nil
		n.mu.Unlock()
	}
}

func (n *Node) FreeRecursive() {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.node != nil {
		C.YGNodeFreeRecursive(n.node)
		n.node = nil
	}
}

func (n *Node) Reset() {
	n.mu.Lock()
	defer n.mu.Unlock()

	C.YGNodeReset(n.node)
}

func (n *Node) GetNodeType() NodeType {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return fromYGNodeType(C.YGNodeGetNodeType(n.node))
}

func (n *Node) SetNodeType(nodeType NodeType) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	ygNodeType, err := toYGNodeType(nodeType)
	if err != nil {
		return err
	}
	C.YGNodeSetNodeType(n.node, ygNodeType)
	return nil
}

func (n *Node) GetChildCount() int {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return len(n.children)
}

func (n *Node) GetChild(index int) *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if index < 0 || index >= len(n.children) {
		return nil
	}

	return n.children[index]
}

func (n *Node) Children() iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		count := n.GetChildCount()

		n.mu.RLock()
		defer n.mu.RUnlock()

		for i := range count {
			child := n.GetChild(i)
			if !yield(child) {
				return
			}
		}
	}
}

func (n *Node) AppendChild(child *Node) {
	count := n.GetChildCount()

	child.SetParent(n)

	n.mu.Lock()
	defer n.mu.Unlock()
	C.YGNodeInsertChild(n.node, child.node, C.size_t(count))
	n.children = append(n.children, child)
}

func (n *Node) SetChildren(children []*Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if len(children) == 0 {
		C.YGNodeRemoveAllChildren(n.node)
		n.children = n.children[:0]
		return
	}

	cChildren := make([]C.YGNodeRef, len(children))
	n.children = n.children[:0]
	for i, child := range children {
		child.SetParent(n)

		cChildren[i] = child.node
		n.children = append(n.children, child)
	}

	C.YGNodeSetChildren(n.node, &cChildren[0], C.size_t(len(children)))

}

func (n *Node) InsertChild(child *Node, index int) {
	n.mu.Lock()
	defer n.mu.Unlock()

	child.SetParent(n)

	C.YGNodeInsertChild(n.node, child.node, C.size_t(index))
	n.children = slices.Insert(n.children, index, child)
}

func (n *Node) RemoveChild(child *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	child.SetParent(nil)

	C.YGNodeRemoveChild(n.node, child.node)
	for i, c := range n.children {
		if c == child {
			n.children = append(n.children[:i], n.children[i+1:]...)
			break
		}
	}
}

func (n *Node) RemoveAllChildren() {
	n.mu.Lock()
	defer n.mu.Unlock()

	for _, child := range n.children {
		child.SetParent(nil)
	}

	C.YGNodeRemoveAllChildren(n.node)
	n.children = n.children[:0]
}

func (n *Node) GetParent() *Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.parent
}

func (n *Node) SetParent(parent *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.parent = parent
}

func (n *Node) SetConfig(config *Config) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.config = config
	C.YGNodeSetConfig(n.node, config.config)
}

func (n *Node) GetConfig() *Config {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.config
}
