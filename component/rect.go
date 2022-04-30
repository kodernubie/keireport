package component

import (
	"github.com/kodernubie/keireport/core"
	"github.com/kodernubie/keireport/util"
)

type Rect struct {
	Base
	Fill   *Fill
	Border *Border
}

type RectBuilder struct {
}

func (o *RectBuilder) Build(template map[string]interface{}, fields map[string]interface{}) (core.Component, error) {

	ret := &Rect{}

	ret.Base.SetData(template)

	ret.Border = &Border{
		Width:  0.2,
		Color:  "0x000000",
		Left:   true,
		Top:    true,
		Right:  true,
		Bottom: true,
	}

	ret.Border.Init(util.GetMap("border", template))

	ret.Fill = &Fill{
		Type:  "transparent",
		Color: "#FFFFFF",
	}

	ret.Fill.Init(util.GetMap("fill", template))

	if ret.PrintTime == 0 {

		o.Update(ret, fields)
	}

	return ret, nil
}

func (o *RectBuilder) Update(comp interface{}, fields map[string]interface{}) error {

	var ret error

	return ret
}

func init() {

	core.RegisterComponent("rect", &RectBuilder{})
}
