package pdf

import (
	"github.com/kodernubie/keireport/component"
	"github.com/kodernubie/keireport/core"
)

type LabelExporter struct {
}

func (o *LabelExporter) Export(report *core.Keireport, exporter *PDFExporter, comp interface{}) error {

	var err error

	label, _ := comp.(*component.Label)

	if label != nil {

		style := ""

		if label.Font.Bold {

			style += "B"
		}

		if label.Font.Italic {

			style += "I"
		}

		if label.Font.Underscore {

			style += "U"
		}

		if label.Font.Strikeout {

			style += "S"
		}

		exporter.pdf.SetFont(label.Font.Name, style, label.Font.Size)
		exporter.pdf.SetXY(report.Margin.Left+label.Left, exporter.curBandTop+label.Top)
		exporter.pdf.Cell(label.Width, label.Height, label.Value)
	}

	return err
}

func init() {

	RegisterComponent("label", &LabelExporter{})
}
