package component

import "github.com/kodernubie/keireport/util"

type Base struct {
	Type      string
	Left      float64
	Top       float64
	Width     float64
	Height    float64
	PrintTime int
}

func (o *Base) SetData(config map[string]interface{}) {

	o.Type = util.GetString("type", config)
	o.Left = util.GetFloat("left", config)
	o.Top = util.GetFloat("top", config)
	o.Width = util.GetFloat("width", config)
	o.Height = util.GetFloat("height", config)
	o.PrintTime = util.GetInt("printTime", config)
}

func (o *Base) GetType() string {

	return "label"
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

func (o *Base) GetPrintTime() int {

	return o.PrintTime
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

		o.Name = util.GetString("name", config, "Arial")
		o.Size = util.GetFloat("size", config, 12)
		o.Bold = util.GetBool("bold", config, false)
		o.Underscore = util.GetBool("underscore", config, false)
		o.Italic = util.GetBool("italic", config, false)
		o.Strikeout = util.GetBool("strikeout", config, false)
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

		o.Width = util.GetFloat("width", config, 0.2)
		o.Color = util.GetString("color", config, "#000000")
		o.Left = util.GetBool("left", config, false)
		o.Top = util.GetBool("top", config, false)
		o.Right = util.GetBool("right", config, false)
		o.Bottom = util.GetBool("bottom", config, false)
	}
}
