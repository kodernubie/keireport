package pdf

import (
	"bytes"

	"github.com/go-pdf/fpdf"
	"github.com/kodernubie/keireport/component"
	"github.com/kodernubie/keireport/core"
)

type ImageExporter struct {
}

func (o *ImageExporter) Export(report *core.Keireport, exporter *PDFExporter, comp interface{}) error {

	var err error

	image, _ := comp.(*component.Image)

	if image != nil {

		opt := fpdf.ImageOptions{}

		targetName := image.Value
		res, mime := report.GetResource(image.Value)

		if res != nil {

			rdr := bytes.NewReader(res)

			tp := exporter.pdf.ImageTypeFromMime(mime)
			exporter.pdf.RegisterImageReader(targetName, tp, rdr)
		} else {

			targetName = report.GetResourceFileName(image.Value)
		}

		exporter.pdf.ImageOptions(targetName,
			report.Margin.Left+image.Left, exporter.curBandTop+image.Top, image.Width, image.Height,
			false, opt, 0, "")
	}

	return err
}

func init() {

	RegisterExporter("image", &ImageExporter{})
}
