package html

import (
	"github.com/kodernubie/keireport/component"
	"github.com/kodernubie/keireport/core"
)

type LabelExporter struct {
}

func (o *LabelExporter) Export(report *core.Keireport, comp interface{}) string {

	ret := ""

	label, _ := comp.(*component.Label)

	if label != nil {

		ret = `<div style='position: absolute; ` +
			WriteAttr("left", label.Left, report.UnitLength) + `;` +
			WriteAttr("top", label.Top, report.UnitLength) + `;` +
			WriteAttr("width", label.Width, report.UnitLength) + `;` +
			WriteAttr("height", label.Height, report.UnitLength) + `;` +
			`'>` +
			label.Value +
			`</div>`
	}

	return ret
}

func init() {

	RegisterComponent("label", &LabelExporter{})
}
