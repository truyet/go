package identicon


type Shape func(Graphics, float64, int32)

type Shapes struct {
	CenterShapes	[]Shape
	OuterShapes	[]Shape
}

func NewShapes() Shapes {
	centerShapes := []Shape{
		func (g Graphics, cell float64, index int32) {
			k := cell * 0.42
			points := []Point{Point{0, 0}, Point{cell, 0}, Point{cell, cell - k * 2},
			Point{cell - k, cell}, Point{0, cell}}
			g.addPolygon(points, false)
		},
		func (g Graphics, cell float64, index int32) {
			w := (cell * 0.5)
			h := (cell * 0.8)
			g.addTriangle(cell - w, 0, w, h, 2, false)
		},
		func (g Graphics, cell float64, index int32) {
			s := (cell / 3)
			g.addRectangle(s, s, cell - s, cell - s, false);
		},
		func (g Graphics, cell float64, index int32) {
			inner := (cell * 0.1)
			outer := (cell * 0.25)
			g.addRectangle(outer, outer, cell - inner - outer, cell - inner - outer, false);
		},
		func (g Graphics, cell float64, index int32) {
			m := (cell * 0.15)
			s := (cell * 0.5)
			g.addCircle(cell - s - m, cell - s - m, s, false);
		},
		func (g Graphics, cell float64, index int32) {
			inner := cell * 0.1
			outer := inner * 4

			g.addRectangle(0, 0, cell, cell, false);
			g.addPolygon([]Point{Point{outer,outer}, Point{cell - inner, outer},
				Point{outer + (cell - outer - inner) / 2, cell - inner}}, true);
		},
		func (g Graphics, cell float64, index int32) {
			g.addPolygon([]Point{Point{0, 0}, Point{cell, 0}, Point{cell, cell * 0.7},
				Point{cell * 0.4, cell * 0.4}, Point{cell * 0.7, cell}, Point{0, cell}}, false);
		},
		func (g Graphics, cell float64, index int32) {
			g.addTriangle(cell/2, cell/2, cell/2, cell/2, 3, false)
		},
		func (g Graphics, cell float64, index int32) {
			g.addRectangle(0, 0, cell, cell / 2, false)
			g.addRectangle(0, cell / 2, cell / 2, cell / 2, false)
			g.addTriangle(cell / 2, cell / 2, cell / 2, cell / 2, 1, false)
		},
		func (g Graphics, cell float64, index int32) {
			inner := (cell * 0.14)
			outer := (cell * 0.35)
			g.addRectangle(0, 0, cell, cell, false)
			g.addRectangle(outer, outer, cell - outer - inner, cell - outer - inner, true)
		},
		func (g Graphics, cell float64, index int32) {
			inner := cell * 0.12
			outer := inner * 3

			g.addRectangle(0, 0, cell, cell, false)
			g.addCircle(outer, outer, cell - inner - outer, true)
		},
		func (g Graphics, cell float64, index int32) {
			g.addTriangle(cell / 2, cell / 2, cell / 2, cell / 2, 3, false);
		},
		func (g Graphics, cell float64, index int32) {
			m := cell * 0.25
			g.addRectangle(0, 0, cell, cell, false)
			g.addRhombus(m, m, cell - m, cell - m, true)
		},
		func (g Graphics, cell float64, index int32) {
			m := cell * 0.4
			s := cell * 1.2
			if !(index < 0) {
				g.addCircle(m, m, s, false);
			}
		}}

	outerShapes := []Shape{
		func (g Graphics, cell float64, index int32) {
			g.addTriangle(0, 0, cell, cell, 0, false);
		},
		func (g Graphics, cell float64, index int32) {
			g.addTriangle(0, cell / 2, cell, cell / 2, 0, false);
		},
		func (g Graphics, cell float64, index int32) {
			g.addRhombus(0, 0, cell, cell, false);
		},
		func (g Graphics, cell float64, index int32) {
			m := cell / 6;
			g.addCircle(m, m, cell - 2 * m, false);
		}}

	return Shapes{CenterShapes:centerShapes, OuterShapes:outerShapes}
}


