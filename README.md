<h1 align="center"><code>tess</code></h1>

<p align="center">A layout engine in Go powered by <a href="https://yogalayout.dev">Yoga</a>.</p>

```go
// Create a root Node
root, err := tess.NewNode()
root.SetWidth(tess.Point(400))
root.SetHeight(tess.Point(300))
root.SetPadding(tess.Edges{All: tess.Point(20)})

// Add a child to the root Node
child, err := tess.NewNode()
child.SetWidth(tess.Percent(100))
child.SetHeight(tess.Point(150))
root.AddChild(child)

// Compute the layout to get final positions and sizes
root.ComputeLayout(tess.Container{})

// Get computed layout for a child
layout := root.GetChild(0).GetLayout()
// Use these values to render your element on screen
layout.Left()
layout.Top()
layout.Width()
layout.Height()
```

## Introduction

`tess` determines the position and size of UI element using browser-style flexbox and style attributes.

It can be used to build terminal UIs, game interfaces, PDF layouts, or anything else where you need flexible box-model positioning.

### Core concepts

**Nodes** - are the building blocks of your layout tree. Each node represents a UI element that needs positioning and sizing.

**Styles** - controls how nodes behave (width, height, padding, flex direction, etc). Think CSS properties, but for any rendering system.

**Layouts** - contains the final calculated positions and sizes of nodes that you can use to render elements on screen.

## Usage

### Installation

```bash
go get github.com/AnatoleLucet/tess
```

### Getting started

As a starter project, let's center a ~~div~~ node! (I know right!?)

```go
package main

import (
	"fmt"

	"github.com/AnatoleLucet/tess"
)

func main() {
  // Create a 500x500 container with children aligned and justified to center
  root, err := tess.NewNode()
  root.SetWidth(tess.Point(500)) // Points are the base unit for all measurements (think pixels but with a scale factor applied)
  root.SetHeight(tess.Point(500))
  root.SetJustifyContent(tess.JustifyCenter)
  root.SetAlignItems(tess.AlignCenter)

  // Add a child of 200x200 to the container
  child, err := tess.NewNode()
  child.SetWidth(tess.Point(200))
  child.SetHeight(tess.Point(200))
  root.AddChild(child)

  root.ComputeLayout(tess.Container{})

  layout := child.GetLayout()
  layout.Left() // 150 (child is centered!)
  layout.Top() // 150
}
```

### Styling

Apply styles to nodes either by calling setter methods or passing a `Style` struct at creation.

```go
root, err := tess.NewNode()
root.SetWidth(tess.Point(500))
root.SetHeight(tess.Point(500))
// or
root, err := tess.NewNode(&tess.Style{
  Width: tess.Point(500),
  Height: tess.Point(500),
})
```

#### Units

`tess` supports multiple unit types for sizing and positioning:

```go
node.SetWidth(tess.Point(500))    // Absolute points
node.SetWidth(tess.Percent(100))  // Percentage of parent

node.SetWidth(tess.Auto())        // Automatic sizing
node.SetWidth(tess.Stretch())     // Stretch to fill
node.SetWidth(tess.MaxContent())  // Size to content
node.SetWidth(tess.FitContent())  // Fit content with constraints

node.SetWidth(tess.Undefined())   // Unset (use default behavior)
```

#### Edges

```go
node.SetMargin(tess.Edges{Horizontal: tess.Point(10)})
node.SetMargin(tess.Edges{Start: tess.Point(10), End: tess.Point(10)})

node.SetPadding(tess.Edges{All: tess.Point(20)})

node.SetBorder(tess.Edges{Bottom: tess.Point(2)})
```

### Computing layouts

After creating your node tree and setting styles, call `ComputeLayout()` to calculate positions and sizes:

```go
err := root.ComputeLayout(tess.Container{})
```

The Container parameter specifies available space constraints. Pass an empty container to use the root node's dimensions, or specify width/height to constrain the layout:

```go
err := root.ComputeLayout(tess.Container{
  Width: 1920,
  Height: 1080,

  Direction: tess.RTL, // You can also pass a direction for LTR support!
})
```

Call `ComputeLayout()` again whenever you change styles or the node tree to recalculate positions.

### Configuring

Configs control global layout behavior. By default, all nodes share the same config.

```go
config := tess.NewConfig()

// Controls pixel density for layout rounding - set to your display's DPI scale
config.SetPointScaleFactor(2.0) // e.g. 2x supersampling for high-DPI displays

// Apply config to a node (and optionally its children if they share this config)
root, err := tess.NewNode()
root.SetConfig(config)
```

## TODOs

- [x] measurement callbacks
- [ ] baseline callbacks
- [ ] more examples

## Credits

- [yoga](https://yogalayout.dev) for powering everything under the hood (_so ~~much~~ little salt in this bad boy_)
