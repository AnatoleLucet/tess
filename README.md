<h1 align="center"><code>tess</code></h1>

<p align="center">A layout engine in Go powered by <a href="https://yogalayout.dev">Yoga</a>.</p>

```go
// Create a root Node
root := tess.NewNode()
root.SetWidth(tess.Point(400))
root.SetHeight(tess.Point(300))
root.SetPadding(tess.Edges{All: tess.Point(20)})

// Add a child to the root Node
child := tess.NewNode()
child.SetWidth(tess.Percent(100))
child.SetHeight(tess.Point(150))
root.AddChild(child)

// Compute the layout to get positions and sizes
root.ComputeLayout(tess.Container{})

// Get computed layout for a child
layout := root.GetChild(0).GetLayout()
layout.Left()
layout.Top()
layout.Width()
layout.Height()
```

## Installation

```bash
go get github.com/AnatoleLucet/tess
```

## Usage

...

## TODOs

- [ ] better error handling (avoid panics)
- [ ] measurement callbacks
- [ ] baseline callbacks
- [ ] more examples
