package core

import (
	"errors"
	"regexp"

	"github.com/kodernubie/keireport/util"
)

var ErrEndOfRow = errors.New("End of row")
var regexField *regexp.Regexp
var regexParam *regexp.Regexp
var regexVar *regexp.Regexp

func init() {

	regexField, _ = regexp.Compile(`\$F\{[A-Za-z1-90_]+\}`)
	regexParam, _ = regexp.Compile(`\$P\{[A-Za-z1-90_]+\}`)
	regexVar, _ = regexp.Compile(`\$V\{[A-Za-z1-90_]+\}`)
}

type Component interface {
	GetType() string
	GetLeft() float64
	GetTop() float64
	GetWidth() float64
	GetHeight() float64
	GetPrintOn() string
}

type Band struct {
	Top        float64
	Height     float64
	AutoSize   bool
	Components []Component
}

type Page struct {
	Bands []*Band
}

type DataSource interface {
	SetConfig(data map[string]interface{}) error
	Next(rpt *Keireport) (map[string]interface{}, error)
}

type DataSourceBuilder interface {
	Build(data map[string]interface{}) (DataSource, error)
}

//------------------------------------------------------

type Margin struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (o *Margin) Init(config map[string]interface{}) {

	if config != nil {

		o.Left = util.GetFloat("left", config, 25.4)
		o.Top = util.GetFloat("top", config, 25.4)
		o.Right = util.GetFloat("right", config, 25.4)
		o.Bottom = util.GetFloat("bottom", config, 25.4)
	}
}

//-------------------------------------------

type Parameter struct {
	Name       string
	Type       string
	DefaultVal interface{}
	Value      interface{}
}

func (o *Parameter) Init(data map[string]interface{}) {

	o.Name = util.GetString("name", data)
	o.Type = util.GetString("type", data)
	o.DefaultVal = data["defaultVal"]

	if o.DefaultVal == nil {
		o.DefaultVal = ""
	}

	o.Value = o.DefaultVal
}

func (o *Parameter) GetFloat(defValue ...float64) float64 {

	ret, ok := o.Value.(float64)

	if !ok {

		var ret32 float32

		ret32, ok = o.Value.(float32)

		if ok {

			ret = float64(ret32)
		}
	}

	if !ok {
		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Parameter) GetInt(defValue ...int) int {

	ret, ok := o.Value.(int)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Parameter) GetString(defValue ...string) string {

	ret, ok := o.Value.(string)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Parameter) GetBool(defValue ...bool) bool {

	ret, ok := o.Value.(bool)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

//-------------------------------------------

type Variable struct {
	Name       string
	Type       string
	InitialVal interface{}
	Value      interface{}
	Expression string
	ExecuteOn  string
}

func (o *Variable) Init(data map[string]interface{}) {

	o.Name = util.GetString("name", data)
	o.Type = util.GetString("type", data, "string")
	o.Expression = util.GetString("expression", data, "")
	o.ExecuteOn = util.GetString("executeOn", data, "row")

	o.InitialVal = data["initialVal"]

	if o.InitialVal == nil {
		o.InitialVal = ""
	}

	o.Value = o.InitialVal
}

func (o *Variable) GetInt(defValue ...int) int {

	ret, ok := o.Value.(int)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Variable) GetFloat(defValue ...float64) float64 {

	ret, ok := o.Value.(float64)

	if !ok {

		var ret32 float32

		ret32, ok = o.Value.(float32)

		if ok {

			ret = float64(ret32)
		}
	}

	if !ok {
		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Variable) GetString(defValue ...string) string {

	ret, ok := o.Value.(string)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

func (o *Variable) GetBool(defValue ...bool) bool {

	ret, ok := o.Value.(bool)

	if !ok {

		if len(defValue) > 0 {

			ret = defValue[0]
		}
	}

	return ret
}

//-------------------------------------------

type ComponentBuilder interface {
	Build(template map[string]interface{}, rpt *Keireport) (Component, error)
	Update(comp interface{}, rpt *Keireport) error
}

type Exporter interface {
	IsHandling(fileName string) bool
	ExportToFile(report *Keireport, fileName string) error
	Export(report *Keireport) ([]byte, error)
}

var builderMap map[string]ComponentBuilder = map[string]ComponentBuilder{}
var exporterMap map[string]Exporter = map[string]Exporter{}
var datasourceMap map[string]DataSourceBuilder = map[string]DataSourceBuilder{}

// Register --------------------------------------------------------------

func RegisterComponent(name string, builder ComponentBuilder) {

	builderMap[name] = builder
}

func RegisterExporter(name string, exporter Exporter) {

	exporterMap[name] = exporter
}

func RegisterDatasource(name string, ds DataSourceBuilder) {

	datasourceMap[name] = ds
}
