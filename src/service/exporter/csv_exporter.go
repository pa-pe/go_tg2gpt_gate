package exporter

import (
	"encoding/csv"
	"net/http"
)

type csvExporter struct {
	writer *csv.Writer
}

func (csvExporter) SetDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Disposition", "attachment; filename="+"result"+".csv")
	w.Header().Set("Content-Type", "text/csv")
}
func (c *csvExporter) SetWriter(w http.ResponseWriter) {
	c.writer = csv.NewWriter(w)
}
func (c *csvExporter) Write(fields []string) error {

	if err := c.writer.Write(fields); err != nil {
		return err
	}
	c.writer.Flush()
	if err := c.writer.Error(); err != nil {
		return err
	}
	return nil
}
func (csvExporter) Close() error {
	return nil
}
