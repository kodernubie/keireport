package pdf

import (
	"regexp"
	"strconv"

	"github.com/kodernubie/keireport/component"
	"github.com/kodernubie/keireport/core"
)

type RectExporter struct {
}

func (o *RectExporter) Export(report *core.Keireport, exporter *PDFExporter, comp interface{}) error {

	var err error

	rect, _ := comp.(*component.Rect)

	if rect != nil {

		var r int64 = 0
		var g int64 = 0
		var b int64 = 0

		regex, _ := regexp.Compile(`^#[A-Fa-f0-9]{6}$`)

		switch rect.Fill.Type {

		case "solid":
			if regex.MatchString(rect.Fill.Color) {

				//#001122
				//0123456
				r, _ = strconv.ParseInt(rect.Fill.Color[1:3], 16, 64)
				g, _ = strconv.ParseInt(rect.Fill.Color[3:3], 16, 64)
				b, _ = strconv.ParseInt(rect.Fill.Color[1:3], 16, 64)
			}

			exporter.pdf.SetFillColor(int(r), int(g), int(b))
			exporter.pdf.Rect(report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top,
				rect.Width, rect.Height, "F")
		}

		if regex.MatchString(rect.Border.Color) {

			//#001122
			//0123456
			r, _ = strconv.ParseInt(rect.Border.Color[1:3], 16, 64)
			g, _ = strconv.ParseInt(rect.Border.Color[3:3], 16, 64)
			b, _ = strconv.ParseInt(rect.Border.Color[1:3], 16, 64)
		} else {

			r = 0
			g = 0
			b = 0
		}

		exporter.pdf.SetLineWidth(rect.Border.Width)
		exporter.pdf.SetDrawColor(int(r), int(g), int(b))

		exporter.pdf.SetXY(report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top)

		if rect.Border.Top {

			exporter.pdf.Line(report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top,
				report.Margin.Left+rect.Left+rect.Width, exporter.curBandTop+rect.Top)
		}

		if rect.Border.Right {

			exporter.pdf.Line(report.Margin.Left+rect.Left+rect.Width, exporter.curBandTop+rect.Top,
				report.Margin.Left+rect.Left+rect.Width, exporter.curBandTop+rect.Top+rect.Height)
		}

		if rect.Border.Bottom {

			exporter.pdf.Line(report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top+rect.Height,
				report.Margin.Left+rect.Left+rect.Width, exporter.curBandTop+rect.Top+rect.Height)
		}

		if rect.Border.Left {

			exporter.pdf.Line(report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top,
				report.Margin.Left+rect.Left, exporter.curBandTop+rect.Top+rect.Height)
		}
	}

	return err
}

func init() {

	RegisterExporter("rect", &RectExporter{})
}
