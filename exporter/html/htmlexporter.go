package html

import (
	"fmt"
	"strings"

	"github.com/kodernubie/keireport/core"
	"github.com/yosssi/gohtml"
)

type HTMLExporter struct {
}

type HTMLCompExporter interface {
	Export(report *core.Keireport, comp interface{}) string
}

var HTMLComponentMap map[string]HTMLCompExporter = map[string]HTMLCompExporter{}

func (o *HTMLExporter) IsHandling(fileName string) bool {

	ret := false

	fileName = strings.ToLower(strings.TrimSpace(fileName))

	if strings.HasSuffix(fileName, ".htm") || strings.HasSuffix(fileName, ".html") {

		ret = true
	}

	return ret
}

func (o *HTMLExporter) ExportToFile(report *core.Keireport, fileName string) error {

	return nil
}

func (o *HTMLExporter) Export(report *core.Keireport) ([]byte, error) {

	var ret []byte
	var err error

	htmlStr := `<html>
<head>
</head>
<body>	
`

	for _, page := range report.Pages {

		pageInnerHTML := ""

		for _, band := range page.Bands {

			bandInnerHTML := ""

			for _, comp := range band.Components {

				exporter, _ := HTMLComponentMap[comp.GetType()]

				if exporter != nil {

					bandInnerHTML += exporter.Export(report, comp)
				}
			}

			pageInnerHTML += `<div stye='position: relative; ` +
				WriteAttr("height", band.Height, report.UnitLength) + `;` +
				`'>` +
				bandInnerHTML +
				`</div>`
		}

		htmlStr += `<div stye='position: relative; ` +
			WriteAttr("width", report.PageWidth, report.UnitLength) + `;` +
			WriteAttr("height", report.PageHeight, report.UnitLength) + `;` +
			WriteAttr("margin-left", report.Margin.Left, report.UnitLength) + `;` +
			WriteAttr("margin-top", report.Margin.Top, report.UnitLength) + `;` +
			WriteAttr("margin-right", report.Margin.Right, report.UnitLength) + `;` +
			WriteAttr("margin-bottom", report.Margin.Bottom, report.UnitLength) + `;` +
			`'>` +
			pageInnerHTML +
			`</div>`
	}

	htmlStr += `</body>`

	fmt.Println(gohtml.Format(htmlStr))
	ret = []byte(htmlStr)

	return ret, err
}

func RegisterComponent(name string, component HTMLCompExporter) {

	HTMLComponentMap[name] = component
}

func WriteAttr(name string, data interface{}, unit ...string) string {

	var ret string = name + ":"
	var typ = fmt.Sprintf("%T\n", data)

	switch typ {
	case "float64":
	case "float32":
		ret += fmt.Sprintf("%.2f", data)
	default:
		ret += fmt.Sprintf("%v", data)
	}

	if len(unit) > 0 {

		ret += unit[0]
	}

	return ret
}
