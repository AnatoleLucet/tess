package main

import (
	"fmt"
	"math"

	"github.com/AnatoleLucet/tess"
)

func main() {
	// Create root container
	root := tess.NewNode()
	defer root.Free()

	root.SetWidth(400)
	root.SetHeight(300)
	root.SetFlexDirection(tess.FlexDirectionRow)
	root.SetJustifyContent(tess.JustifySpaceBetween)
	root.SetAlignItems(tess.AlignCenter)
	root.SetPadding(tess.EdgeAll, 20)
	root.SetGap(tess.GutterColumn, 10)

	// Create first child with fixed dimensions
	child1 := tess.NewNode()
	defer child1.Free()
	child1.SetWidth(100)
	child1.SetHeight(100)
	child1.SetMargin(tess.EdgeAll, 5)

	// Create second child with percentage-based width
	child2 := tess.NewNode()
	defer child2.Free()
	child2.SetWidthPercent(30)
	child2.SetHeight(120)
	child2.SetAlignSelf(tess.AlignFlexStart)

	// Create third child with flex grow
	child3 := tess.NewNode()
	defer child3.Free()
	child3.SetFlexGrow(1)
	child3.SetHeight(80)
	child3.SetMargin(tess.EdgeLeft, 10)

	// Add children to root
	root.InsertChild(child1, 0)
	root.InsertChild(child2, 1)
	root.InsertChild(child3, 2)

	// Calculate layout
	root.CalculateLayout(float32(math.NaN()), float32(math.NaN()), tess.DirectionLTR)

	// Print results
	fmt.Println("=== Tess Layout Example ===")
	fmt.Printf("\nRoot: %.2f x %.2f\n", root.GetLayoutWidth(), root.GetLayoutHeight())
	fmt.Printf("  Padding: left=%.2f, top=%.2f, right=%.2f, bottom=%.2f\n",
		root.GetLayoutPadding(tess.EdgeLeft),
		root.GetLayoutPadding(tess.EdgeTop),
		root.GetLayoutPadding(tess.EdgeRight),
		root.GetLayoutPadding(tess.EdgeBottom))

	for i := 0; i < root.GetChildCount(); i++ {
		child := root.GetChild(i)
		layout := child.GetLayout()
		fmt.Printf("\nChild %d: left=%.2f, top=%.2f, width=%.2f, height=%.2f\n",
			i+1, layout.Left, layout.Top, layout.Width, layout.Height)
	}
}
