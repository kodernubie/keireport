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

		border := ""

		if label.Border.Left {

			border += "L"
		}

		if label.Border.Top {

			border += "T"
		}

		if label.Border.Right {

			border += "R"
		}

		if label.Border.Bottom {

			border += "B"
		}

		align := ""

		switch label.AlignHor {
		case "left":
			align += "L"
		case "right":
			align += "R"
		case "center":
			align += "C"
		}

		switch label.AlignVer {
		case "top":
			align += "T"
		case "bottom":
			align += "B"
		case "middle":
			align += "M"
		}

		fill := false
		exporter.pdf.CellFormat(label.Width, label.Height, label.Value, border, 1, align, fill, 0, "")
	}

	return err
}

func init() {

	RegisterComponent("label", &LabelExporter{})
}
