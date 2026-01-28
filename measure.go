package tess

/*
#include <yoga/Yoga.h>
#include <stdint.h>

extern YGSize goMeasureCallback(YGNodeConstRef node, float width, YGMeasureMode widthMode, float height, YGMeasureMode heightMode);

static inline void tessSetMeasureFunc(YGNodeRef node) {
	YGNodeSetMeasureFunc(node, goMeasureCallback);
}

static inline void tessUnsetMeasureFunc(YGNodeRef node) {
	YGNodeSetMeasureFunc(node, NULL);
}

static inline void tessSetNodeContext(YGNodeRef node, uintptr_t handle) {
	YGNodeSetContext(node, (void*)handle);
}

static inline uintptr_t tessGetNodeContext(YGNodeRef node) {
	return (uintptr_t)YGNodeGetContext(node);
}
*/
import "C"
import "runtime/cgo"

type MeasureFunc func(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size

type measureContext struct {
	node *Node
	fn   MeasureFunc
}

//export goMeasureCallback
func goMeasureCallback(node C.YGNodeConstRef, width C.float, widthMode C.YGMeasureMode, height C.float, heightMode C.YGMeasureMode) C.YGSize {
	ctxPtr := C.tessGetNodeContext(C.YGNodeRef(node))
	if ctxPtr == 0 {
		return C.YGSize{width: 0, height: 0}
	}

	ctx := cgo.Handle(ctxPtr).Value().(*measureContext)
	size := ctx.fn(ctx.node, float32(width), fromYGMeasureMode(widthMode), float32(height), fromYGMeasureMode(heightMode))

	return C.YGSize{width: C.float(size.Width), height: C.float(size.Height)}
}

func (n *Node) SetMeasureFunc(fn MeasureFunc) {
	if n.HasMeasureFunc() {
		n.UnsetMeasureFunc()
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	ctx := &measureContext{node: n, fn: fn}
	handle := cgo.NewHandle(ctx)
	C.tessSetNodeContext(n.node, C.uintptr_t(handle))
	C.tessSetMeasureFunc(n.node)
}

func (n *Node) HasMeasureFunc() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return bool(C.YGNodeHasMeasureFunc(n.node))
}

func (n *Node) UnsetMeasureFunc() {
	n.mu.Lock()
	defer n.mu.Unlock()

	ctx := C.tessGetNodeContext(n.node)
	if ctx != 0 {
		handle := cgo.Handle(ctx)
		C.tessSetNodeContext(n.node, 0)
		handle.Delete()
	}

	C.tessUnsetMeasureFunc(n.node)
}

func (n *Node) getMeasureFunc() MeasureFunc {
	ctxPtr := C.tessGetNodeContext(n.node)
	if ctxPtr == 0 {
		return nil
	}

	handle := cgo.Handle(ctxPtr)
	ctx, ok := handle.Value().(*measureContext)
	if !ok {
		return nil
	}

	return ctx.fn
}

func (n *Node) reregisterMeasureFunc(fn MeasureFunc) {
	n.clearContext()

	if fn == nil {
		C.tessUnsetMeasureFunc(n.node)
		return
	}

	ctx := &measureContext{node: n, fn: fn}
	handle := cgo.NewHandle(ctx)
	C.tessSetNodeContext(n.node, C.uintptr_t(handle))
	C.tessSetMeasureFunc(n.node)
}

func (n *Node) clearContext() {
	C.tessSetNodeContext(n.node, 0)
}
