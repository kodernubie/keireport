package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/kodernubie/keireport"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	//Simple()
	//CustomFont()
	//ParameterAndVariable()
	CustomConnection()
}

/*
	Simple load template and generate to pdf
	Database connection using config in template file
*/
func Simple() {

	rpt, err := keireport.LoadFromFile("simple.krpt")

	if err == nil {

		rpt.GenToFile("simple.pdf")

		fmt.Println("generated to : simple.pdf")
	}
}

/*
	Embed external TTF font. See "fonts" field for detail
	Using empty datasource
	You can change band printing behavior when datasource is empty
	by changing "printWhenEmpty" field
*/
func CustomFont() {

	rpt, err := keireport.LoadFromFile("customFont.krpt")

	if err == nil {

		rpt.GenToFile("customFont.pdf")

		fmt.Println("generated to : customFont.pdf")
	}
}

/*
	Set parameter to report so the report can generate dynamic content based on parameter
	The template contains variable usage example too
*/
func ParameterAndVariable() {

	rpt, err := keireport.LoadFromFile("variable.krpt")

	if err == nil {

		rpt.SetParam("trxId", 2)
		rpt.GenToFile("variable.pdf")

		fmt.Println("generated to : variable.pdf")
	}
}

/*
	Supply connection programmatically
*/
func CustomConnection() {

	rpt, err := keireport.LoadFromFile("variable.krpt")

	if err == nil {

		db, err := sql.Open("pgx", "user=postgres password=admin host=localhost dbname=keisample2 port=5432 sslmode=disable TimeZone=Asia/Jakarta")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		rpt.SetDBConn(db)
		rpt.GenToFile("variable.pdf")

		fmt.Println("generated to : variable.pdf")
	}
}
