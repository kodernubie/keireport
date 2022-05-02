package component

import (
	"regexp"

	"github.com/kodernubie/keireport/util"
)

var regexField *regexp.Regexp

func init() {

	regexField, _ = regexp.Compile(`\$F\{[A-Za-z1-90_]+\}`)
}

type Base struct {
	Type    string
	Left    float64
	Top     float64
	Width   float64
	Height  float64
	PrintOn string
}

func (o *Base) SetData(config map[string]interface{}) {

	o.Type = util.GetString("type", config)
	o.Left = util.GetFloat("left", config, 0)
	o.Top = util.GetFloat("top", config, 0)
	o.Width = util.GetFloat("width", config, 20)
	o.Height = util.GetFloat("height", config, 20)
	o.PrintOn = util.GetString("printOn", config, "now")
}

func (o *Base) GetType() string {

	return o.Type
}

func (o *Base) GetLeft() float64 {

	return o.Left
}

func (o *Base) GetTop() float64 {

	return o.Top
}

func (o *Base) GetWidth() float64 {

	return o.Width
}

func (o *Base) GetHeight() float64 {

	return o.Height
}

func (o *Base) GetPrintOn() string {

	return o.PrintOn
}

//-------------------------------------

type Font struct {
	Name       string
	Size       float64
	Bold       bool
	Underscore bool
	Italic     bool
	Strikeout  bool
}

func (o *Font) Init(config map[string]interface{}) {

	if config != nil {

		o.Name = util.GetString("name", config, o.Name)
		o.Size = util.GetFloat("size", config, o.Size)
		o.Bold = util.GetBool("bold", config, o.Bold)
		o.Underscore = util.GetBool("underscore", config, o.Underscore)
		o.Italic = util.GetBool("italic", config, o.Italic)
		o.Strikeout = util.GetBool("strikeout", config, o.Strikeout)
	}
}

//-------------------------------------

type Border struct {
	Width  float64
	Color  string
	Left   bool
	Top    bool
	Right  bool
	Bottom bool
}

func (o *Border) Init(config map[string]interface{}) {

	if config != nil {

		o.Width = util.GetFloat("width", config, o.Width)
		o.Color = util.GetString("color", config, o.Color)
		o.Left = util.GetBool("left", config, o.Left)
		o.Top = util.GetBool("top", config, o.Top)
		o.Right = util.GetBool("right", config, o.Right)
		o.Bottom = util.GetBool("bottom", config, o.Bottom)
	}
}

//-------------------------------------

type Fill struct {
	Type  string
	Color string
}

func (o *Fill) Init(config map[string]interface{}) {

	if config != nil {

		o.Type = util.GetString("type", config, o.Type)
		o.Color = util.GetString("color", config, o.Color)
	}
}
