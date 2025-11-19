package tess

import "fmt"

type Value struct {
	unit  Unit
	value float32
}

func (v Value) Unit() Unit {
	return v.unit
}

func (v Value) Value() float32 {
	return v.value
}

func (v Value) String() string {
	switch v.unit {
	case UnitUndefined:
		return "undefined"
	case UnitPoint:
		return fmt.Sprintf("%.2fpt", v.value)
	case UnitPercent:
		return fmt.Sprintf("%.2f%%", v.value)
	case UnitAuto:
		return "auto"
	case UnitMaxContent:
		return "max-content"
	case UnitFitContent:
		return "fit-content"
	case UnitStretch:
		return "stretch"
	}

	return "unknown"
}

type Unit int

const (
	UnitUndefined Unit = iota
	UnitPoint
	UnitPercent
	UnitMaxContent
	UnitFitContent
	UnitAuto
	UnitStretch
)

func (u Unit) String() string {
	switch u {
	case UnitUndefined:
		return "undefined"
	case UnitPoint:
		return "point"
	case UnitPercent:
		return "percent"
	case UnitMaxContent:
		return "max-content"
	case UnitFitContent:
		return "fit-content"
	case UnitAuto:
		return "auto"
	case UnitStretch:
		return "stretch"
	}

	return "unknown"
}

func Point(points float32) Value     { return Value{unit: UnitPoint, value: points} }
func Percent(percents float32) Value { return Value{unit: UnitPercent, value: percents} }
func Undefined() Value               { return Value{unit: UnitUndefined} }
func MaxContent() Value              { return Value{unit: UnitMaxContent} }
func FitContent() Value              { return Value{unit: UnitFitContent} }
func Auto() Value                    { return Value{unit: UnitAuto} }
func Stretch() Value                 { return Value{unit: UnitStretch} }

// Edges represents values for the edges of a box.
// Note: Start and End will always be prioritized over Left and Right when both are set.
type Edges struct {
	Top, Bottom, Left, Right Value
	Start, End               Value
	Horizontal, Vertical     Value
	All                      Value
}

// Gap represents the gap values for rows and columns.
type Gap struct {
	Row    Value
	Column Value
	All    Value
}
