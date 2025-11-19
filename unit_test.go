package tess

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {
	t.Run("Value String method", func(t *testing.T) {
		v1 := Value{unit: UnitPoint, value: 12.34}
		assert.Equal(t, "12.34pt", v1.String())

		v2 := Value{unit: UnitPercent, value: 56.78}
		assert.Equal(t, "56.78%", v2.String())

		v3 := Value{unit: UnitAuto}
		assert.Equal(t, "auto", v3.String())

		v4 := Value{unit: UnitMaxContent}
		assert.Equal(t, "max-content", v4.String())

		v5 := Value{unit: UnitFitContent}
		assert.Equal(t, "fit-content", v5.String())

		v6 := Value{unit: UnitUndefined}
		assert.Equal(t, "undefined", v6.String())

		v7 := Value{unit: UnitStretch}
		assert.Equal(t, "stretch", v7.String())
	})
}

func TestUnit(t *testing.T) {
	t.Run("Unit String method", func(t *testing.T) {
		assert.Equal(t, "undefined", UnitUndefined.String())
		assert.Equal(t, "point", UnitPoint.String())
		assert.Equal(t, "percent", UnitPercent.String())
		assert.Equal(t, "max-content", UnitMaxContent.String())
		assert.Equal(t, "fit-content", UnitFitContent.String())
		assert.Equal(t, "auto", UnitAuto.String())
		assert.Equal(t, "stretch", UnitStretch.String())
		assert.Equal(t, "unknown", Unit(999).String())
	})
}
