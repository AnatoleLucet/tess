package main

import (
	"fmt"

	"github.com/AnatoleLucet/tess"
)

func main() {
	// Create root container
	root, _ := tess.NewNode(&tess.Style{
		Width:  tess.Point(400),
		Height: tess.Point(300),

		FlexDirection:  tess.Row,
		JustifyContent: tess.JustifySpaceBetween,
		AlignItems:     tess.AlignCenter,

		Padding: tess.Edges{All: tess.Point(20)},
		Gap:     tess.Gap{Column: tess.Point(20)},
	})
	defer root.Free()

	// Create first child with fixed dimensions
	child1, _ := tess.NewNode(&tess.Style{
		Width:  tess.Point(100),
		Height: tess.Point(100),
		Margin: tess.Edges{All: tess.Point(5)},
	})
	defer child1.Free()

	// Create second child with percentage-based width
	child2, _ := tess.NewNode(&tess.Style{
		Width:  tess.Percent(30),
		Height: tess.Point(120),

		AlignSelf: tess.AlignStart,
	})
	defer child2.Free()

	// Create third child with flex grow
	child3, _ := tess.NewNode(&tess.Style{
		FlexGrow: 1,
		Height:   tess.Point(80),

		Margin: tess.Edges{Left: tess.Point(10)},
	})
	defer child3.Free()

	// Add children to root
	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)

	// Calculate layout
	_ = root.ComputeLayout(tess.Container{Direction: tess.LTR})

	// Print results
	fmt.Println("=== Tess Layout Example ===")
	fmt.Printf("\nRoot: %.2f x %.2f\n", root.GetLayout().Width(), root.GetLayout().Height())
	fmt.Printf("  Padding: left=%.2f, top=%.2f, right=%.2f, bottom=%.2f\n",
		root.GetLayout().Padding().Left(),
		root.GetLayout().Padding().Top(),
		root.GetLayout().Padding().Right(),
		root.GetLayout().Padding().Bottom())

	for i := 0; i < root.GetChildCount(); i++ {
		child := root.GetChild(i)
		layout := child.GetLayout()
		fmt.Printf("\nChild %d: left=%.2f, top=%.2f, width=%.2f, height=%.2f\n",
			i+1, layout.Left(), layout.Top(), layout.Width(), layout.Height())
	}
}
