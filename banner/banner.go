package banner

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var (
	size = image.Rect(0, 0, 200, 400)
)

type Banner struct {
	base   color.RGBA
	layers []Layer

	img          *image.RGBA
	currentColor color.RGBA
}

type Layer struct {
	Color   color.RGBA
	Pattern string
}

func NewBanner(baseColor color.RGBA, layers []Layer) *Banner {
	return &Banner{
		base:         baseColor,
		layers:       layers,
		img:          image.NewRGBA(size),
		currentColor: color.RGBA{A: 0xff},
	}
}

func (b *Banner) fillRect(x, y, width, height int) {
	col := b.currentColor
	alpha := col.A
	col.A = 0xff
	draw.DrawMask(b.img, image.Rect(x, y, x+width, y+height),
		image.NewUniform(col), image.Point{}, image.NewUniform(color.Alpha{A: alpha}), image.Point{}, draw.Over)
}

func (b *Banner) setColor(col color.RGBA) {
	b.currentColor = col
}

func (b *Banner) Draw() {
	b.setColor(b.base)
	b.fillRect(0, 0, 200, 400)
	b.setColor(Mask)
	b.fillRect(0, 0, 200, 10)
	b.fillRect(0, 390, 200, 10)
	b.fillRect(0, 10, 10, 380)
	b.fillRect(190, 10, 10, 380)

	for _, layer := range b.layers {
		b.setColor(layer.Color)
		if drawFunc, ok := patterns[layer.Pattern]; ok {
			drawFunc(b, layer.Color)
		} else {
			log.Printf("Draw: couldn't find a draw func for %v\n", layer.Pattern)
		}
	}
}

func (b *Banner) Output(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	return png.Encode(f, b.img)
}
