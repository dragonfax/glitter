package squares

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var _ ElementWidget = Text{}
var _ Element = &TextElement{}

const CHARACTER_WIDTH = 10
const CHARACTER_HEIGHT = 10

type Text struct {
	Text string
}

func (t Text) createElement() Element {
	te := &TextElement{}
	te.widget = t
	return te
}

type TextElement struct {
	elementData
}

func (t *TextElement) layout(c Constraints) error {
	cWidth := float64(len(t.widget.(Text).Text) * CHARACTER_WIDTH)
	cHeight := 1.0 * CHARACTER_HEIGHT

	t.size = Size{Width: math.Max(cWidth, c.minWidth), Height: math.Max(cHeight, c.minHeight)}

	return nil
}

func (t *TextElement) render(offset Offset, renderer *sdl.Renderer) {
	ux := int32(offset.x)
	uy := int32(offset.y)
	surface, err := font.RenderUTF8Blended(t.widget.(Text).Text, sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		panic(err)
	}
	defer surface.Free()
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	renderer.Copy(texture, nil, &sdl.Rect{X: ux, Y: uy, W: surface.W, H: surface.H})
}
