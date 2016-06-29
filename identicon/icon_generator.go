package identicon

import (
	"errors"
	"regexp"
	"image/color"
	"math"
	"fmt"
)

type IconGenerator struct  {
	renderer 		Renderer
	graphics		Graphics
	availableColors		[]color.Color
	hash			string
	x, y			float64
	size, padding		float64
	config 			Config
}

func fromHexChar(c byte) (int32, bool) {
	switch {
	case '0' <= c && c <= '9':
		return int32(c - '0'), true
	case 'a' <= c && c <= 'f':
		return int32(c - 'a' + 10), true
	case 'A' <= c && c <= 'F':
		return int32(c - 'A' + 10), true
	}

	return 0, false
}

func hexStringToInt(str string) int32 {
	var intVal int32
	src := []byte(str)
	pow := len(src) - 1
	for _, v := range src {
		hexCharToInt, _ := fromHexChar(v)
		intVal += hexCharToInt * int32(math.Pow(16, float64(pow)))
		pow--
	}
	return intVal
}

func NewIconGenerator (renderer Renderer, hash string, x, y float64, size, padding float64, config Config) (*IconGenerator, error) {
	if padding < 0 {
		padding = 0.08
	}

	padding = (size * padding)
	size -= padding * 2

	if size < 30 {
		return nil, errors.New("Gdenticon cannot render identicons smaller than 30 pixels.")
	}

	validHash := regexp.MustCompile("^[0-9a-f]{11,}$")
	if !validHash.MatchString(hash) {
		return nil, errors.New("Invalid hash passed to Gdenticon.")
	}

	graphics := NewGraphics(renderer)

	// Calculate cell size and ensure it is an integer
	var cell = int32(size / 4);

	// Since the cell size is integer based, the actual icon will be slightly smaller than specified => center icon
	x += padding + size / 2 - float64(cell) * 2;
	y += padding + size / 2 - float64(cell) * 2;

	// AVAILABLE COLORS
	hue := float64(hexStringToInt(hash[len(hash)-7:])) / 0xfffffff

	// Available colors for this icon
	availableColors := colorTheme(hue, config)
	fmt.Printf("%v", availableColors)
	fmt.Printf("\nhash, x, y, size, padding %v,%v,%v,%v,%v", hash, x, y, size, padding)

	return &IconGenerator{renderer, graphics, availableColors, hash, x, y, size, padding, config}, nil
}

func (icg *IconGenerator) Generator() {
	shapes := NewShapes()
	//Sides
	icg.renderShape(0, shapes.OuterShapes, 2, 3 , []Point{Point{1, 0}, Point{2, 0}, Point{2,3}, Point{1,3}, Point{0,1},
	Point{3,1}, Point{3,2}, Point{0,2}})
	// Corners
	icg.renderShape(1, shapes.OuterShapes, 4, 5, []Point{Point{0, 0}, Point{3, 0}, Point{3,3}, Point{0,3}})
	//// Center
	icg.renderShape(2, shapes.CenterShapes, 1, -1, []Point{Point{1, 1}, Point{2, 1}, Point{2, 2}, Point{1, 2}})
}

//func (icg *IconGenerator) Close() {
//	icg.graphics.renderer.
//}

func (icg *IconGenerator) renderShape(colorIndex float64, shapes []Shape, index float64, rotationIndex float64,
positions []Point) {
	var r int32
	if rotationIndex >= 0 {
		r = hexStringToInt(icg.hash[int(rotationIndex): int(rotationIndex)+1])
	}

	cell := icg.size / 4;
	selectedColorIndexes := []int32{}
	var i int32
	for i=0; i < 3; i++ {
		cIndex := hexStringToInt(string(icg.hash[8+i])) % int32(len(icg.availableColors))
		if isDuplicate(selectedColorIndexes, []int32{0, 4}, cIndex) ||
			isDuplicate(selectedColorIndexes, []int32{2, 3}, cIndex) {
			cIndex = 1;
		}
		selectedColorIndexes = append(selectedColorIndexes, cIndex)
	}
	shape_index := hexStringToInt(string(icg.hash[int(index)])) % int32(len(shapes))
	fmt.Printf("\nselectedColorIndexes, shape_index %v, %v", selectedColorIndexes, shape_index)
	icg.renderer.beginShape(icg.availableColors[selectedColorIndexes[int(colorIndex)]])
	for i, v := range positions {
		icg.graphics.transform = NewTransform(icg.x + v.x * cell, icg.y + v.y * cell, cell, float64(r % 4))
		r++
		shapes[shape_index](icg.graphics, cell, int32(i))
	}

	icg.renderer.endShape()
}

func isDuplicate(src []int32, values []int32, value int32) bool {
	if (indexOf(values, value) >= 0) {
		for i := 0; i < len(values); i++ {
			if (indexOf(src, values[i]) >= 0) {
				return true;
			}
		}
	}
	return false
}

func indexOf(values []int32, value int32) int32 {
	for i, v := range values {
		if v == value {
			return int32(i)
		}
	}

	return -1;
}

