package keireport

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {

	rpt, err := LoadFromFile("example/simple.krpt")
	rpt.Debug = true

	if err == nil {

		err = rpt.GenToFile("example/simple.pdf")

		if err != nil {

			fmt.Println("Error generate :", err.Error())
		}
	} else {

		fmt.Println("Error load template :", err.Error())
	}
}

func TestFont(t *testing.T) {

	rpt, err := LoadFromFile("example/customFont.krpt")
	rpt.Debug = true

	if err == nil {

		err = rpt.GenToFile("example/customFont.pdf")

		if err != nil {

			fmt.Println("Error generate :", err.Error())
		}
	} else {

		fmt.Println("Error load template :", err.Error())
	}
}

func TestVariable(t *testing.T) {

	rpt, err := LoadFromFile("example/variable.krpt")
	rpt.Debug = true

	if err == nil {

		err = rpt.GenToFile("example/variable.pdf")

		if err != nil {

			fmt.Println("Error generate :", err.Error())
		}
	} else {

		fmt.Println("Error load template :", err.Error())
	}
}
