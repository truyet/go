package identicon

import (
	"image/color"
	"math"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d"
	"fmt"
	"golang.org/x/image/draw"
)

type Renderer interface  {
	beginShape(color color.Color)
	endShape()
	addPolygon( points []Point)
	addCircle(point Point, diameter float64, counterClockwise bool)
}

type ImageRenderer struct {
	ctx 		draw2d.GraphicContext

}

func NewImageRenderer(dest draw.Image) *ImageRenderer {
	return &ImageRenderer{ctx:draw2dimg.NewGraphicContext(dest)}
}

func (renderer *ImageRenderer) beginShape(col color.Color) {
	r,g,b,a := col.RGBA()
	fmt.Printf("\n%v,%v,%v,%v", r, g, b, a)
	renderer.ctx.SetFillColor(col)
	//renderer.ctx.BeginPath()
}

func (renderer *ImageRenderer) endShape() {
	renderer.ctx.Fill()
	//renderer.ctx.Save()
}

func (renderer *ImageRenderer) addPolygon( points []Point) {
	fmt.Printf("\n Polygon %v", points)
	renderer.ctx.MoveTo(float64(points[0].x), points[0].y)
	for _, point := range points {
		renderer.ctx.LineTo(point.x, point.y)
	}
	renderer.ctx.Close()

}

func (renderer *ImageRenderer) addCircle( point Point, diameter float64, counterClockwise bool) {
	fmt.Printf("\n Circle, diameter %v,%v ", point, diameter)
	radius := diameter / 2
	angle := math.Pi * 2
	if !counterClockwise {
		angle = -angle
	}
	renderer.ctx.ArcTo(point.x + radius, point.y + radius, radius, radius, 0, angle)
	renderer.ctx.Close()
}
