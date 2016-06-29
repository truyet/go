package identicon

type Transform struct {
	x 		float64
	y 		float64
	size 		float64
	rotation 	float64
}

func (t *Transform) transformPoint (x float64, y float64, w float64, h float64) Point {
	var point 	Point
	right := t.x + t.size
	bottom := t.y + t.size
	if t.rotation == 1 {
		point = NewPoint(right - y - h, t.y + x)
	} else if t.rotation == 2 {
		point = NewPoint(right - x - h, bottom - y - h)
	} else if t.rotation == 3 {
		point = NewPoint(t.x + y, bottom - x - w)
	} else {
		point = NewPoint(t.x + x, t.y + y)
	}

	return point

}

func NoTransform() Transform {
	return Transform{x:0, y:0, size:0, rotation: 0}
}

func NewTransform(x, y, size float64, rotation float64) Transform {
	return Transform{x, y, size, rotation}
}
