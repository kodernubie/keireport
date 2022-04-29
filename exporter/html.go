package exporter

import (
	"github.com/kodernubie/keireport/core"
	"github.com/kodernubie/keireport/exporter/html"
)

func init() {

	core.RegisterExporter("html", &html.HTMLExporter{})
}
