package identicon

import (
	"image"
	"crypto/md5"
	"github.com/llgcode/draw2d/draw2dimg"
	"testing"
)

func Test(t *testing.T)  {
	dest := image.NewRGBA(image.Rect(0, 0, 300, 300))
	renderer := NewImageRenderer(dest)
	hash := string(md5.New().Sum([]byte("Jeremiah")))
	config := Config{}
	identicon, err := NewIconGenerator(renderer, hash, 0, 0, 300, 8, config)
	if err != nil {
		t.Error(err)
	}
	identicon.Generator()
	draw2dimg.SaveToPngFile("./test.png", dest)

}


