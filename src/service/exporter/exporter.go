package exporter

import (
	"errors"
	"net/http"
)

type ExpType string

const (
	TypeCsv  ExpType = "csv"
	TypeXlsx ExpType = "xlsx"
)

type Exporter interface {
	SetDefaultHeaders(w http.ResponseWriter)
	SetWriter(w http.ResponseWriter)
	Write([]string) error
	Close() error
}

func NewExporter(exporterType string) (Exporter, error) {
	var exp Exporter
	if exporterType == string(TypeCsv) {
		exp = getCsvExporter()
	} else if exporterType == string(TypeXlsx) {
		exp = getXlsExporter()
	} else {
		return nil, errors.New("wrong exporter type ")
	}

	return exp, nil
}

func getCsvExporter() *csvExporter {
	return &csvExporter{}
}

func getXlsExporter() *xlsExporter {
	return &xlsExporter{}
}
