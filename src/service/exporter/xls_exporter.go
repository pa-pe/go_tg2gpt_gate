package exporter

import (
	"github.com/tealeg/xlsx"
	"net/http"
)

type xlsExporter struct {
	writer     *xlsx.StreamFileBuilder
	streamFile *xlsx.StreamFile
}

func (xlsExporter) SetDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Disposition", "attachment; filename=result.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
}
func (c *xlsExporter) SetWriter(w http.ResponseWriter) {
	c.writer = xlsx.NewStreamFileBuilder(w)
}
func (c *xlsExporter) Write(fields []string) error {
	if c.streamFile == nil {
		err := c.writer.AddSheet("Sheet", fields, []*xlsx.CellType{})
		if err != nil {
			return err
		}
		sf, err := c.writer.Build()
		if err != nil {
			return err
		}
		c.streamFile = sf

		return nil
	}

	err := c.streamFile.Write(fields)
	if err != nil {
		return err
	}

	return nil
}
func (c *xlsExporter) Close() error {
	if c.streamFile != nil {
		err := c.streamFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
