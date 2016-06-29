package identicon


type Config struct {
	Saturation 	float64
	LightnessConfig map[string][]float64
}

func DefaultConfig() *Config {
	return &Config{Saturation:0.5, LightnessConfig:make(map[string][]float64)}
}

func (c *Config) lightness(configName string, defaulMin, defaultMax float64) []float64  {
	if c.LightnessConfig[configName] != nil {
		return c.LightnessConfig[configName]
	} else {
		return []float64{defaulMin, defaultMax}
	}
}

func (c *Config) colorLightness(value float64) float64 {
	lightnessRange := c.lightness("color", 0.4, 0.8)
	result := lightnessRange[0] + value * (lightnessRange[1] - lightnessRange[0])
	if result < 0 {
		result = 0;
	} else if result > 1 {
		result = 1
	}

	return result
}

func (c *Config) grayScaleLightness(value float64) float64 {
	lightnessRange := c.lightness("grayscale", 0.3, 0.9)
	result := lightnessRange[0] + value * (lightnessRange[1] - lightnessRange[0])
	if result < 0 {
		result = 0;
	} else if result > 1 {
		result = 1
	}

	return result
}

