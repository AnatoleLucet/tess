package main

import (
	"fmt"
	"math"

	"github.com/AnatoleLucet/tess"
)

func main() {
	root := tess.NewNode()
	defer root.Free()

	root.SetWidth(100)
	root.SetHeight(100)
	root.SetFlexDirection(tess.FlexDirectionRow)
	root.SetJustifyContent(tess.JustifyCenter)

	child := tess.NewNode()
	defer child.Free()

	child.SetWidth(50)
	child.SetHeight(50)

	root.InsertChild(child, 0)
	root.CalculateLayout(float32(math.NaN()), float32(math.NaN()))

	rootLayout := root.GetLayout()
	childLayout := child.GetLayout()

	fmt.Printf("Root layout: left=%.2f, top=%.2f, width=%.2f, height=%.2f\n",
		rootLayout.Left, rootLayout.Top, rootLayout.Width, rootLayout.Height)
	fmt.Printf("Child layout: left=%.2f, top=%.2f, width=%.2f, height=%.2f\n",
		childLayout.Left, childLayout.Top, childLayout.Width, childLayout.Height)
}
