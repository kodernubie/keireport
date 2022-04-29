package pdf

import (
	"fmt"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/kodernubie/keireport/core"
)

type PDFExporter struct {
	pdf        *fpdf.Fpdf
	curBandTop float64
}

type PDFCompExporter interface {
	Export(report *core.Keireport, exporter *PDFExporter, comp interface{}) error
}

var PDFComponentMap map[string]PDFCompExporter = map[string]PDFCompExporter{}

func (o *PDFExporter) IsHandling(fileName string) bool {

	ret := false

	fileName = strings.ToLower(strings.TrimSpace(fileName))

	if strings.HasSuffix(fileName, ".pdf") {

		ret = true
	}

	return ret
}

func (o *PDFExporter) doExport(report *core.Keireport) error {

	o.pdf = fpdf.New(report.Orientation, report.UnitLength, report.PageSize, "")
	o.pdf.SetFont("Arial", "", 12)
	var err error

	for _, page := range report.Pages {

		o.pdf.AddPage()
		o.curBandTop = report.Margin.Top

		for _, band := range page.Bands {

			for _, comp := range band.Components {

				exporter, _ := PDFComponentMap[comp.GetType()]

				if exporter != nil {

					exporter.Export(report, o, comp)
				}
			}

			o.curBandTop += band.Height
		}
	}

	return err
}

func (o *PDFExporter) ExportToFile(report *core.Keireport, fileName string) error {

	err := o.doExport(report)

	if err == nil {

		err = o.pdf.OutputFileAndClose(fileName)
	}

	return err
}

func (o *PDFExporter) Export(report *core.Keireport) ([]byte, error) {

	var ret []byte
	err := o.doExport(report)

	if err == nil {

		fmt.Println("generate byte array")
	}

	return ret, err

}

func RegisterComponent(name string, component PDFCompExporter) {

	PDFComponentMap[name] = component
}
