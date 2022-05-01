package main

import (
	"github.com/kodernubie/keireport"
)

func main() {

	simple()
	customFont()
}

/*
	Embed custom TTF font. See "fonts" field for detail
	Using empty datasource
	You can change band printin behavior when datasource is empty
	by changing "printWhenEmpty" field
*/
func customFont() {

	rpt, err := keireport.LoadFromFile("customFont.krpt")

	if err == nil {

		rpt.GenToFile("customFont.pdf")
	}
}

/*
	Simple load template and generate to pdf
	Database connection using config in template file
*/
func simple() {

	rpt, err := keireport.LoadFromFile("simple.krpt")

	if err == nil {

		rpt.GenToFile("simple.pdf")
	}
}
