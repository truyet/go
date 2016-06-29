package identicon

import "image/color"

func correctedHLS(h, l, s float64) HLS{
	correctors := [...]float64{0.55, 0.5, 0.5, 0.46, 0.6, 0.55, 0.55}
	corrector := correctors[int32(h * 6 + 0.5)]
	if l < 0.5 {
		l = l * corrector * 2
	} else {
		l = corrector + (l - 0.5) *  (1 - corrector) * 2
	}

	return HLS{h, l, s}
}

func colorTheme(hue float64, c Config) []color.Color {
	var colors = []color.Color{}
	colors = append(colors, HLS{0, c.grayScaleLightness(0), 0})
	colors = append(colors, correctedHLS(hue, c.colorLightness(0.5), c.Saturation))
	colors = append(colors, HLS{0, c.grayScaleLightness(1), 0})
	colors = append(colors, correctedHLS(hue, c.colorLightness(1), c.Saturation))
	colors = append(colors, correctedHLS(hue, c.colorLightness(0), c.Saturation))

	return colors
}