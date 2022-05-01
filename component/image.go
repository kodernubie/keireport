package component

import (
	"fmt"
	"strings"
	"time"

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

func (o *ImageBuilder) Build(template map[string]interface{}, fields map[string]interface{}) (core.Component, error) {

	ret := &Image{}

	ret.Base.SetData(template)
	ret.Src = util.GetString("src", template)

	if ret.PrintTime == 0 {

		o.Update(ret, fields)
	}

	return ret, nil
}

func (o *ImageBuilder) Update(comp interface{}, fields map[string]interface{}) error {

	var ret error

	image, ok := comp.(*Image)

	if ok {

		target := image.Src

		for key, val := range fields {

			valStr := ""

			switch val.(type) {
			case float64:
				valStr = fmt.Sprintf("%f", val.(float64))
			case float32:
				valStr = fmt.Sprintf("%f", val.(float32))
			case time.Time:
				valStr = val.(time.Time).Format("2006-01-02")
			default:
				valStr = fmt.Sprintf("%v", val)
			}

			target = strings.ReplaceAll(target, "$F{"+key+"}", valStr)
		}

		image.Value = target
	}

	return ret
}

func init() {

	core.RegisterComponent("image", &ImageBuilder{})
}
