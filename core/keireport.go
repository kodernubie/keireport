package core

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PaesslerAG/gval"
	"github.com/kodernubie/keireport/util"
)

type Keireport struct {
	BaseDir     string
	SQLDB       *sql.DB
	Debug       bool
	UnitLength  string
	PageSize    string
	PageWidth   float64
	PageHeight  float64
	Orientation string
	Margin      *Margin
	Params      []*Parameter
	Vars        []*Variable
	MaxHeight   float64
	Template    map[string]interface{}
	CurrRow     map[string]interface{}
	Fonts       map[string]string
	DataSource  DataSource
	Pages       []*Page
	CurrentPage *Page
	Resources   map[string]ResourceItem
}

func (o *Keireport) GetResourceFileName(fileName string) string {

	return filepath.Join(o.BaseDir, fileName)
}

func (o *Keireport) GetResource(name string) ([]byte, string) {

	res, ok := o.Resources[name]

	if ok {

		return res.Data, res.Mime
	}

	return nil, ""
}

func (o *Keireport) AddResource(name, mime string, data []byte) {

	o.Resources[name] = ResourceItem{
		Mime: mime,
		Data: data,
	}
}

func (o *Keireport) LoadFromString(templateString, baseDir string) error {

	var err error

	o.BaseDir = baseDir

	b := []byte(templateString)

	err = json.Unmarshal(b, &o.Template)

	if err != nil {

		if o.Debug {
			fmt.Println("Error loading design from file :", err)
		}

		return err
	}

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

	o.Fonts = map[string]string{}
	fontList := util.GetMap("fonts", o.Template)

	for name, target := range fontList {

		targetS, ok := target.(string)

		if ok {
			o.Fonts[name] = targetS
		}
	}

	o.Params = []*Parameter{}
	paramList := util.GetArr("params", o.Template)

	for _, item := range paramList {

		target, ok := item.(map[string]interface{})

		if ok {

			par := &Parameter{}
			par.Init(target)

			o.Params = append(o.Params, par)
		}
	}

	o.Vars = []*Variable{}
	varList := util.GetArr("vars", o.Template)

	for _, item := range varList {

		target, ok := item.(map[string]interface{})

		if ok {
			varO := &Variable{}
			varO.Init(target)

			o.Vars = append(o.Vars, varO)
		}
	}

	o.Resources = map[string]ResourceItem{}

	return err
}

func (o *Keireport) LoadFromFile(fileName string) error {

	var err error

	b, err := ioutil.ReadFile(fileName)

	if err != nil {

		// try relative
		path, _ := os.Getwd()

		fileName = path + "/" + fileName
		b, err = ioutil.ReadFile(fileName)
	}

	if err == nil {

		baseDir, _ := filepath.Abs(filepath.Dir(fileName))
		err = o.LoadFromString(string(b), baseDir)
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

			targetComp, err := builder.Build(compData, o)

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

		err = o.updateVar("page")

		if err != nil {
			return err
		}

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

			err = errors.New("datasource config is not defined in template")
		}
	} else {

		if dsTemplate != nil {

			// update config incase supplied datasource is not complete
			o.DataSource.SetConfig(dsTemplate)
		}
	}

	if err == nil {
		o.Pages = []*Page{}
		o.NewPage()

		o.CurrRow, err = o.DataSource.Next(o)

		empty := o.CurrRow == nil

		if err == nil || empty {

			errV := o.updateVar("row")

			if errV != nil {
				return errV
			}

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

									o.CurrRow, err = o.DataSource.Next(o)

									if err == nil {

										errV := o.updateVar("row")

										if errV != nil {
											return errV
										}
									}
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

		// util.PrettyPrint(o.Params)
		// util.PrettyPrint(o.Vars)
	}

	return err
}

func (o *Keireport) SetDBConn(db *sql.DB) {

	o.SQLDB = db
}

func (o *Keireport) SetParam(name string, value interface{}) {

	for _, par := range o.Params {

		if par.Name == name {

			par.Value = value
			break
		}
	}
}

func (o *Keireport) ReplaceString(data string) string {

	target := data

	if regexField.Match([]byte(target)) {

		if o.CurrRow == nil {

			target = regexField.ReplaceAllString(target, "")
		} else {

			for key, val := range o.CurrRow {

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
	}

	if regexVar.Match([]byte(target)) {

		for _, val := range o.Vars {

			valStr := ""

			switch val.Type {
			case "float":
				valStr = fmt.Sprintf("%f", val.GetFloat())
			// case "time":
			// 	valStr = val.GetTime().Format("2006-01-02")
			default:
				valStr = fmt.Sprintf("%v", val.Value)
			}

			target = strings.ReplaceAll(target, "$V{"+val.Name+"}", valStr)
		}
	}

	if regexParam.Match([]byte(target)) {

		for _, val := range o.Params {

			valStr := ""

			switch val.Type {
			case "float":
				valStr = fmt.Sprintf("%f", val.GetFloat())
			// case "time":
			// 	valStr = val.GetTime().Format("2006-01-02")
			default:
				valStr = fmt.Sprintf("%v", val.Value)
			}

			target = strings.ReplaceAll(target, "$P{"+val.Name+"}", valStr)
		}
	}

	return target
}

func (o *Keireport) ExecScript(script string) (interface{}, error) {

	ret := script
	retParam := map[string]interface{}{}
	vars := map[string]interface{}{}
	params := map[string]interface{}{}
	fields := map[string]interface{}{}

	// variable
	for _, varI := range o.Vars {

		ret = strings.ReplaceAll(ret, "$V{"+varI.Name+"}", "vars[\""+varI.Name+"\"]")
		vars[varI.Name] = varI.Value
	}

	// parameter
	for _, varI := range o.Params {

		ret = strings.ReplaceAll(ret, "$P{"+varI.Name+"}", "vars[\""+varI.Name+"\"]")
		params[varI.Name] = varI.Value
	}

	// fields
	for name, val := range o.CurrRow {

		ret = strings.ReplaceAll(ret, "$F{"+name+"}", "fields[\""+name+"\"]")
		fields[name] = val
	}

	retParam["vars"] = vars
	retParam["params"] = params
	retParam["fields"] = fields

	return gval.Evaluate(ret, retParam)
}

func (o *Keireport) updateVar(updateOn string) error {

	var err error

	for _, varO := range o.Vars {

		if varO.ExecuteOn == updateOn {

			var ret interface{}

			ret, err = o.ExecScript(varO.Expression)

			if err == nil {

				varO.Value = ret
			} else {

				fmt.Println(varO.Expression, varO.Name, "===>ERROR EXECUTE==>", err.Error())
				break
			}
		}
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
