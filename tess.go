package tess

/*
#cgo CFLAGS: -I${SRCDIR}/etc/include
#cgo CXXFLAGS: -I${SRCDIR}/etc/include
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_amd64 -lyogacore -lstdc++ -lm
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_arm64 -lyogacore -lstdc++ -lm
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_amd64 -lyogacore -lc++
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_arm64 -lyogacore -lc++
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/etc/lib/windows_amd64/yogacore.lib
#cgo windows,arm64 LDFLAGS: ${SRCDIR}/etc/lib/windows_arm64/yogacore.lib
#include <yoga/Yoga.h>
*/
import "C"
import "math"

type Node struct {
	node C.YGNodeRef
}

func NewNode() *Node {
	return &Node{
		node: C.YGNodeNew(),
	}
}

func (n *Node) Free() {
	if n.node != nil {
		C.YGNodeFree(n.node)
		n.node = nil
	}
}

func (n *Node) SetWidth(width float32) {
	C.YGNodeStyleSetWidth(n.node, C.float(width))
}

func (n *Node) SetHeight(height float32) {
	C.YGNodeStyleSetHeight(n.node, C.float(height))
}

func (n *Node) SetFlexDirection(direction int) {
	C.YGNodeStyleSetFlexDirection(n.node, C.YGFlexDirection(direction))
}

func (n *Node) SetJustifyContent(justify int) {
	C.YGNodeStyleSetJustifyContent(n.node, C.YGJustify(justify))
}

func (n *Node) InsertChild(child *Node, index int) {
	C.YGNodeInsertChild(n.node, child.node, C.size_t(index))
}

func (n *Node) CalculateLayout(width, height float32) {
	w := C.float(width)
	h := C.float(height)
	if math.IsNaN(float64(width)) {
		w = C.float(math.NaN())
	}
	if math.IsNaN(float64(height)) {
		h = C.float(math.NaN())
	}
	C.YGNodeCalculateLayout(n.node, w, h, C.YGDirectionLTR)
}

type Layout struct {
	Left   float32
	Top    float32
	Width  float32
	Height float32
}

func (n *Node) GetLayout() Layout {
	return Layout{
		Left:   float32(C.YGNodeLayoutGetLeft(n.node)),
		Top:    float32(C.YGNodeLayoutGetTop(n.node)),
		Width:  float32(C.YGNodeLayoutGetWidth(n.node)),
		Height: float32(C.YGNodeLayoutGetHeight(n.node)),
	}
}

const (
	FlexDirectionColumn = int(C.YGFlexDirectionColumn)
	FlexDirectionRow    = int(C.YGFlexDirectionRow)
	JustifyCenter       = int(C.YGJustifyCenter)
)
