package tess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("set point scale factor", func(t *testing.T) {
		config := NewConfig()
		defer config.free()

		config.SetPointScaleFactor(2.0)
		scale := config.GetPointScaleFactor()
		assert.Equal(t, float32(2.0), scale)
	})

	t.Run("default config point scale factor", func(t *testing.T) {
		config := getDefaultConfig()
		scale := config.GetPointScaleFactor()
		assert.Equal(t, float32(1.0), scale) // Default is 1.0
	})

	t.Run("multiple default config calls return same instance", func(t *testing.T) {
		config1 := getDefaultConfig()
		config2 := getDefaultConfig()
		assert.Equal(t, config1, config2)
	})
}
