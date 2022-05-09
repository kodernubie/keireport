package component

import (
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

func (o *LabelBuilder) Build(template map[string]interface{}, rpt *core.Keireport) (core.Component, error) {

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

		o.Update(ret, rpt)
	}

	return ret, nil
}

func (o *LabelBuilder) Update(comp interface{}, rpt *core.Keireport) error {

	var ret error

	label, ok := comp.(*Label)

	if ok {

		label.Value = rpt.ReplaceString(label.Expression)
	}

	return ret
}

func init() {

	core.RegisterComponent("label", &LabelBuilder{})
}
