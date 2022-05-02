package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kodernubie/keireport/util"
)

var ErrEndOfRow = errors.New("End of row")

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
	Next() (map[string]interface{}, error)
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

		o.Left = util.GetFloat("name", config, 25.4)
		o.Top = util.GetFloat("top", config, 25.4)
		o.Right = util.GetFloat("right", config, 25.4)
		o.Bottom = util.GetFloat("bottom", config, 25.4)
	}
}

//-------------------------------------------

type Keireport struct {
	BaseDir      string
	Debug        bool
	TemplateFile string
	UnitLength   string
	PageSize     string
	PageWidth    float64
	PageHeight   float64
	Orientation  string
	Margin       *Margin
	MaxHeight    float64
	Template     map[string]interface{}
	CurrRow      map[string]interface{}
	Fonts        map[string]string
	DataSource   DataSource
	Pages        []*Page
	CurrentPage  *Page
}

type ComponentBuilder interface {
	Build(template map[string]interface{}, fields map[string]interface{}) (Component, error)
	Update(comp interface{}, fields map[string]interface{}) error
}

type Exporter interface {
	IsHandling(fileName string) bool
	ExportToFile(report *Keireport, fileName string) error
	Export(report *Keireport) ([]byte, error)
}

var builderMap map[string]ComponentBuilder = map[string]ComponentBuilder{}
var exporterMap map[string]Exporter = map[string]Exporter{}
var datasourceMap map[string]DataSourceBuilder = map[string]DataSourceBuilder{}

// Keireport --------------------------------------------------------------

func (o *Keireport) GetResource(fileName string) string {

	baseDir := o.BaseDir

	if o.TemplateFile != "" {

		baseDir, _ = filepath.Abs(filepath.Dir(o.TemplateFile))
	} else if baseDir == "" {

		baseDir, _ = os.Getwd()
		baseDir, _ = filepath.Abs(baseDir)
	}

	return filepath.Join(baseDir, fileName)
}

func (o *Keireport) LoadFromFile(fileName string) error {

	var err error

	b, err := ioutil.ReadFile(fileName)
	o.TemplateFile = fileName

	if err != nil {

		// try relative
		path, _ := os.Getwd()

		b, err = ioutil.ReadFile(path + "/" + fileName)

		o.TemplateFile = path + "/" + fileName
	}

	if err == nil {

		err = json.Unmarshal(b, &o.Template)

		if err != nil {

			fmt.Println("Error loading design from file :", err)
		}
	} else {

		fmt.Println("Error loading design from file :", err)
	}

	return err
}

func (o *Keireport) NewPage() {

	newPage := &Page{
		Bands: []*Band{},
	}

	o.Pages = append(o.Pages, newPage)
	o.CurrentPage = newPage
}

func (o *Keireport) BuildBand(bandTemplate map[string]interface{}) error {

	var err error

	band := &Band{}
	band.Components = []Component{}

	band.Height = util.GetFloat("height", bandTemplate)
	band.AutoSize = util.GetBool("autoSize", bandTemplate)

	if len(o.CurrentPage.Bands) == 0 {

		band.Top = o.Margin.Top
	} else {

		lastBand := o.CurrentPage.Bands[len(o.CurrentPage.Bands)-1]
		band.Top = lastBand.Top + lastBand.Height
	}

	var maxHeight float64 = 0

	comps := util.GetArr("components", bandTemplate)

	for _, comp := range comps {

		compData := comp.(map[string]interface{})
		compType := util.GetString("type", compData)

		builder := builderMap[compType]

		if builder != nil {

			targetComp, err := builder.Build(compData, o.CurrRow)

			if err == nil {

				band.Components = append(band.Components, targetComp)
			} else {

				fmt.Print("[Error] Build comp error : ", err)
			}
		} else {

			fmt.Println("[Error] Comp builder not found", compType)
		}
	}

	if band.AutoSize && band.Height <= maxHeight {

		band.Height = maxHeight
	}

	if band.Top+band.Height > o.MaxHeight {

		o.NewPage()

		bandList, _ := o.Template["bands"].(map[string]interface{})

		// title
		bandTemplate, _ := bandList["header"].(map[string]interface{})

		if band != nil {

			err = o.BuildBand(bandTemplate)
		}

		if len(o.CurrentPage.Bands) == 0 {

			band.Top = o.Margin.Top
		} else {

			lastBand := o.CurrentPage.Bands[len(o.CurrentPage.Bands)-1]
			band.Top = lastBand.Top + lastBand.Height
		}
	}

	if err == nil {

		o.CurrentPage.Bands = append(o.CurrentPage.Bands, band)
	}

	return err
}

func (o *Keireport) Build() error {

	var err error

	// global config
	o.PageSize = util.GetString("pageSize", o.Template, "A4")
	o.Orientation = util.GetString("orientation", o.Template, "P")
	o.UnitLength = util.GetString("unitLength", o.Template, "mm")

	switch o.PageSize {
	case "A4":
		switch o.UnitLength {
		case "mm":
			o.PageWidth = 210
			o.PageHeight = 297
		}
	case "A5":
		switch o.UnitLength {
		case "mm":
			o.PageWidth = 148
			o.PageHeight = 210
		}
	}

	o.Margin = &Margin{
		Left:   25.4,
		Top:    25.4,
		Right:  25.4,
		Bottom: 25.4,
	}

	o.Margin.Init(util.GetMap("margin", o.Template))

	o.MaxHeight = o.PageHeight - o.Margin.Top - o.Margin.Bottom

	// data source
	dsTemplate, _ := o.Template["datasource"].(map[string]interface{})

	if o.DataSource == nil {

		// not supplied from programmatic, try getting from template
		if dsTemplate != nil {

			dsType, _ := dsTemplate["type"].(string)

			dsBuilder := datasourceMap[dsType]

			if dsBuilder != nil {

				o.DataSource, err = dsBuilder.Build(dsTemplate)

			} else {

				err = errors.New("Datasource Builder is not found : " + dsType)
			}
		} else {

			err = errors.New("Datasource config is not defined in template")
		}
	} else {

		if dsTemplate != nil {

			// update config incase supplied datasource is not complete
			o.DataSource.SetConfig(dsTemplate)
		}
	}

	o.Fonts = map[string]string{}
	fontList := util.GetMap("fonts", o.Template)

	for name, target := range fontList {

		targetS, ok := target.(string)

		if ok {
			o.Fonts[name] = targetS
		}
	}

	if err == nil {
		o.Pages = []*Page{}
		o.NewPage()

		o.CurrRow, err = o.DataSource.Next()

		empty := o.CurrRow == nil

		if err == nil || empty {

			bandList, _ := o.Template["bands"].(map[string]interface{})

			if bandList != nil {

				// title
				band, _ := bandList["title"].(map[string]interface{})

				if band != nil {

					err = o.BuildBand(band)
				}

				if err == nil {

					// header
					band, _ := bandList["header"].(map[string]interface{})

					if band != nil {

						if empty {

							if util.GetBool("printWhenEmpty", band, false) {
								err = o.BuildBand(band)
							}
						} else {

							err = o.BuildBand(band)
						}
					}
				}

				if err == nil {

					//detail
					band, _ := bandList["detail"].(map[string]interface{})

					if band != nil {

						if empty {

							if util.GetBool("printWhenEmpty", band, false) {
								err = o.BuildBand(band)
							}
						} else {

							for err == nil {

								err = o.BuildBand(band)

								if err == nil {

									o.CurrRow, err = o.DataSource.Next()
								}
							}

							if errors.Is(err, ErrEndOfRow) {

								err = nil
							}
						}
					}
				}

				if err == nil {

					// footer
					band, _ := bandList["footer"].(map[string]interface{})

					if band != nil {

						if empty {

							if util.GetBool("printWhenEmpty", band, false) {
								err = o.BuildBand(band)
							}
						} else {

							err = o.BuildBand(band)
						}
					}
				}
			}
		}
	}

	if o.Debug {

		//util.PrettyPrint(o.Pages)
	}

	return err
}

func (o *Keireport) Generate(format string) ([]byte, error) {

	exporter, _ := exporterMap[format]

	if exporter == nil {

		return nil, errors.New("Exporter not available for format : " + format)
	}

	var ret []byte
	err := o.Build()

	if err == nil {

		ret, err = exporter.Export(o)
	}

	return ret, err
}

func (o *Keireport) GenToFile(fileName string) error {

	var err error

	var exporter Exporter

	for _, tmp := range exporterMap {

		if tmp.IsHandling(fileName) {

			exporter = tmp
			break
		}
	}

	if exporter != nil {

		err = o.Build()

		if err == nil {

			err = exporter.ExportToFile(o, fileName)
		}
	} else {

		err = errors.New("Exporter not found")
	}

	return err
}

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
