package component

import (
	"fmt"
	"strings"
	"time"

	"github.com/kodernubie/keireport/core"
	"github.com/kodernubie/keireport/util"
)

type Label struct {
	Base
	Font       *Font
	AlignVer   string
	AlignHor   string
	Value      string
	Expression string
	Border     *Border
}

type LabelBuilder struct {
}

func (o *LabelBuilder) Build(template map[string]interface{}, fields map[string]interface{}) (core.Component, error) {

	ret := &Label{}

	ret.Base.SetData(template)
	ret.Expression = util.GetString("expression", template)
	ret.AlignHor = util.GetString("alignHor", template, "left")
	ret.AlignVer = util.GetString("alignVer", template, "top")

	ret.Font = &Font{
		Name:       "Arial",
		Size:       12,
		Bold:       false,
		Underscore: false,
		Italic:     false,
		Strikeout:  false,
	}

	ret.Font.Init(util.GetMap("font", template))

	ret.Border = &Border{
		Width:  0.2,
		Color:  "0x000000",
		Left:   false,
		Top:    false,
		Right:  false,
		Bottom: false,
	}

	ret.Border.Init(util.GetMap("border", template))

	if ret.PrintOn == "now" {

		o.Update(ret, fields)
	}

	return ret, nil
}

func (o *LabelBuilder) Update(comp interface{}, fields map[string]interface{}) error {

	var ret error

	label, ok := comp.(*Label)

	if ok {

		target := label.Expression

		if fields == nil {

			target = regexField.ReplaceAllString(target, "")
		} else {

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
		}

		label.Value = target
	}

	return ret
}

func init() {

	core.RegisterComponent("label", &LabelBuilder{})
}
