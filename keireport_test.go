package keireport

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	rpt, err := LoadFromFile("example/simple.krpt")

	if err == nil {

		err = rpt.GenToFile("example/simple.pdf")

		if err != nil {

			fmt.Println("Error generate :", err.Error())
		}
	} else {

		fmt.Println("Error load template :", err.Error())
	}
}
