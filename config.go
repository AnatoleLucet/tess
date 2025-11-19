package tess

/*
#include <yoga/Yoga.h>
*/
import "C"
import (
	"runtime"
	"sync"
)

type Config struct {
	config C.YGConfigRef
}

var (
	defaultConfig     *Config
	defaultConfigOnce sync.Once
)

func getDefaultConfig() *Config {
	defaultConfigOnce.Do(func() {
		config := C.YGConfigNew()

		// use web defaults for every tess node
		C.YGConfigSetUseWebDefaults(config, true)

		defaultConfig = &Config{config: config}
	})

	return defaultConfig
}

func NewConfig() *Config {
	c := &Config{config: C.YGConfigNew()}
	runtime.SetFinalizer(c, (*Config).free)

	return c
}

func (c *Config) free() {
	if c.config != nil {
		C.YGConfigFree(c.config)
		c.config = nil
	}
}

func (c *Config) SetPointScaleFactor(scale float32) {
	C.YGConfigSetPointScaleFactor(c.config, C.float(scale))
}

func (c *Config) GetPointScaleFactor() float32 {
	return float32(C.YGConfigGetPointScaleFactor(c.config))
}
