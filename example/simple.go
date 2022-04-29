package main

import (
	"github.com/kodernubie/keireport"
)

/*
	Simple load template and generate to pdf
	Database connection using config in template file

*/
func main() {

	rpt, err := keireport.LoadFromFile("simple.krpt")

	if err == nil {

		rpt.GenToFile("simple.pdf")
	}
}
