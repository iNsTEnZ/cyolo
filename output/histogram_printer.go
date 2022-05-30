package output

import (
	"bytes"
	"cyolo-exercise/model"
	"fmt"
)

type Printer interface {
	Print(data interface{}) string
}

type HistogramPrinter struct{}

func NewPrinter() *HistogramPrinter {
	return &HistogramPrinter{}
}

func (srv *HistogramPrinter) Print(data interface{}) string {
	var buffer bytes.Buffer

	if val, ok := data.([]model.Pair); ok {
		for _, pair := range val {
			buffer.WriteString(fmt.Sprintf("%s\t%d\n", pair.Key, pair.Value))
		}

		return buffer.String()
	}

	return ""
}
