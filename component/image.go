package component

import (
	"github.com/kodernubie/keireport/core"
	"github.com/kodernubie/keireport/util"
)

type Image struct {
	Base
	Value string
	Src   string
}

type ImageBuilder struct {
}

func (o *ImageBuilder) Build(template map[string]interface{}, rpt *core.Keireport) (core.Component, error) {

	ret := &Image{}

	ret.Base.SetData(template)
	ret.Src = util.GetString("src", template)

	if ret.PrintOn == "now" {

		o.Update(ret, rpt)
	}

	return ret, nil
}

func (o *ImageBuilder) Update(comp interface{}, rpt *core.Keireport) error {

	var ret error

	image, ok := comp.(*Image)

	if ok {

		image.Value = rpt.ReplaceString(image.Src)
	}

	return ret
}

func init() {

	core.RegisterComponent("image", &ImageBuilder{})
}
