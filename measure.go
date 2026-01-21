package tess

/*
#include <yoga/Yoga.h>

extern YGSize goMeasureCallback(YGNodeConstRef node, float width, YGMeasureMode widthMode, float height, YGMeasureMode heightMode);

static inline void tessSetMeasureFunc(YGNodeRef node) {
	YGNodeSetMeasureFunc(node, goMeasureCallback);
}

static inline void tessUnsetMeasureFunc(YGNodeRef node) {
	YGNodeSetMeasureFunc(node, NULL);
}
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

type MeasureFunc func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size

//export goMeasureCallback
func goMeasureCallback(node C.YGNodeConstRef, width C.float, widthMode C.YGMeasureMode, height C.float, heightMode C.YGMeasureMode) (result C.YGSize) {
	ctx := C.YGNodeGetContext(C.YGNodeRef(node))
	if ctx == nil {
		return C.YGSize{width: 0, height: 0}
	}

	// Handle race condition where the handle may be deleted between
	// checking ctx and calling handle.Value()
	defer func() {
		if r := recover(); r != nil {
			result = C.YGSize{width: 0, height: 0}
		}
	}()

	handle := cgo.Handle(uintptr(ctx))
	fn := handle.Value().(MeasureFunc)

	goNode := &Node{node: C.YGNodeRef(node)}
	size := fn(goNode, float32(width), fromYGMeasureMode(widthMode), float32(height), fromYGMeasureMode(heightMode))

	return C.YGSize{width: C.float(size.Width), height: C.float(size.Height)}
}

func (n *Node) SetMeasureFunc(fn MeasureFunc) {
	if n.HasMeasureFunc() {
		n.UnsetMeasureFunc()
	}

	handle := cgo.NewHandle(fn)
	C.YGNodeSetContext(n.node, unsafe.Pointer(handle))
	C.tessSetMeasureFunc(n.node)
}

func (n *Node) HasMeasureFunc() bool {
	return bool(C.YGNodeHasMeasureFunc(n.node))
}

func (n *Node) UnsetMeasureFunc() {
	ctx := C.YGNodeGetContext(n.node)
	if ctx != nil {
		handle := cgo.Handle(uintptr(ctx))
		// Set context to nil BEFORE deleting handle to prevent race condition
		// where another goroutine sees non-nil context but handle is deleted
		C.YGNodeSetContext(n.node, nil)
		handle.Delete()
	}
	C.tessUnsetMeasureFunc(n.node)
}
