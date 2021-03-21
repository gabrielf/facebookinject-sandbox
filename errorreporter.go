package main

import "fmt"

type ErrorReporter interface {
	ReportError(error)
}

type StdoutErrorReporter struct {
}

func (ser StdoutErrorReporter) ReportError(err error) {
	fmt.Printf("Reporting error: %s\n", err.Error())
}
