package identicon


type Graphics struct {
	renderer	Renderer
	transform	Transform
}

func NewGraphics(r Renderer) Graphics {
	return Graphics{renderer:r, transform:NoTransform()}
}

func (g *Graphics) addPolygon(points []Point, invert bool) {
	di := 1
	i := 0
	if invert {
		di = -1
		i = len(points) - 1
	}

	transformPoints := []Point{}
	for ; i < len(points) && i >= 0; i +=di {
		transformPoints = append(transformPoints, g.transform.transformPoint(points[i].x, points[i].y, 0, 0))
	}
	g.renderer.addPolygon(transformPoints)
}

func (g *Graphics) addCircle(x float64, y float64, size float64, invert bool) {
	point := g.transform.transformPoint(x, y, size, size)
	g.renderer.addCircle(point, size, invert)
}

func (g *Graphics) addRectangle(x, y, w, h float64, invert bool) {
	points := []Point{Point{x,y}, Point{x+w,y}, Point{x+w,y+h}, Point{x, y+h}}
	g.addPolygon(points, invert)
}

func (g *Graphics) addTriangle(x, y, w, h, r float64, invert bool) {
	points := [...]Point{Point{x+w,y}, Point{x+w,y+h}, Point{x, y+h}, Point{x,y}}
	newPoints := []Point{}
	idx := (int(r/2) % 4) * 2
	for _, v := range points[:idx] {
		newPoints = append(newPoints, v)
	}

	for _, v := range points[idx+1:] {
		newPoints = append(newPoints, v)
	}

	g.addPolygon(newPoints, invert)
}

func (g *Graphics) addRhombus(x, y, w, h float64, invert bool) {
	points := []Point{Point{x+w/2,y}, Point{x+w,y+h/2}, Point{x+w/2,y+h}, Point{x, y+h/2}}
	g.addPolygon(points, invert)
}