package exporter

import (
	"github.com/kodernubie/keireport/core"
	"github.com/kodernubie/keireport/exporter/pdf"
)

func init() {

	core.RegisterExporter("pdf", &pdf.PDFExporter{})
}
