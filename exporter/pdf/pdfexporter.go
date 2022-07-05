package pdf

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/kodernubie/keireport/core"
)

type PDFExporter struct {
	pdf        *fpdf.Fpdf
	tempDir    string
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

	o.tempDir, _ = os.MkdirTemp("", "keireport")

	o.pdf = fpdf.New(report.Orientation, report.UnitLength, report.PageSize, o.tempDir)
	o.pdf.SetFont("Arial", "", 12)
	var err error

	// process font
	for name, target := range report.Fonts {

		targetPath := report.GetResourceFileName(target)

		mapFile := filepath.Join(o.tempDir, "cp1252.map")
		os.WriteFile(mapFile, []byte(encStr["cp1252"]), 0644)
		err := fpdf.MakeFont(targetPath, mapFile, o.tempDir, nil, true)

		if err == nil {

			jsonFile := filepath.Base(targetPath)

			if pos := strings.LastIndexByte(jsonFile, '.'); pos != -1 {
				jsonFile = jsonFile[:pos]
			}

			jsonFile += ".json"

			o.pdf.AddFont(name, "", jsonFile)
			o.pdf.AddFont(name, "B", jsonFile)
			o.pdf.AddFont(name, "I", jsonFile)
			o.pdf.AddFont(name, "BI", jsonFile)
		}
	}

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
	defer os.RemoveAll(o.tempDir)

	if err == nil {

		err = o.pdf.OutputFileAndClose(fileName)
	}

	return err
}

func (o *PDFExporter) Export(report *core.Keireport) ([]byte, error) {

	var ret []byte
	err := o.doExport(report)
	defer os.RemoveAll(o.tempDir)

	if err == nil {

		buff := bytes.Buffer{}
		writer := bufio.NewWriter(&buff)

		o.pdf.Output(writer)

		writer.Flush()
		ret = buff.Bytes()
	}

	return ret, err
}

func RegisterExporter(name string, component PDFCompExporter) {

	PDFComponentMap[name] = component
}
