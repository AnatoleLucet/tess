<h1 align="center"><code>tess</code></h1>

<p align="center">A layout engine in Go powered by <a href="https://yogalayout.dev">Yoga</a>.</p>

```go
// Create a root Node
root := tess.NewNode()
root.SetWidth(400) // should probably be able to pass it to NewNode directly
root.SetHeight(300)
root.SetPadding(tess.EdgeAll, 20)

// Add a child to the root Node
child := tess.NewNode()
child.SetWidthPercent(100)
child.SetHeight(150)
root.InsertChild(child, 0)

root.CalculateLayout(float32(math.NaN()), float32(math.NaN()), tess.DirectionLTR) // this should be improved

// Get computed layout size and position
root.GetLayoutWidth()
root.GetLayoutHeight()

// Get computed layout for a child
layout := root.GetChild(0).GetLayout()
layout.Left
layout.Top
layout.Width
layout.Height
```

## Installation

```bash
go get github.com/AnatoleLucet/as
```

## Usage
