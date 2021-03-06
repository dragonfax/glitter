package squares

import (
	"github.com/veandco/go-sdl2/sdl"
)

var _ ElementWidget = Divider{}
var _ Element = &DividerElement{}

type Divider struct {
}

func (d Divider) createElement() Element {
	de := &DividerElement{}
	de.widget = d
	return de
}

type DividerElement struct {
	elementData
}

func (ce *DividerElement) layout(c Constraints) error {

	ce.size = Size{Width: c.maxWidth, Height: 16}

	return nil
}

func (ce *DividerElement) render(offset Offset, renderer *sdl.Renderer) {

	renderer.SetDrawColor(200, 200, 200, 255)
	ux := int32(offset.x)
	uy := int32(offset.y + ce.size.Height/2)
	renderer.DrawLine(ux+0, uy, ux+int32(ce.size.Width), uy)
}
