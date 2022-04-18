package viewmodel

import "github.com/pluralsight/webservice/model"

type Base struct {
	Title  string
	Active string
	Record *model.Record
}

func NewBase() Base {
	return Base{
		Title:  "Stopwatch",
		Active: "stopwatch",
	}
}
