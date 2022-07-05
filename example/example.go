package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/kodernubie/keireport"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	//Simple()
	//CustomFont()
	//ParameterAndVariable()
	//CustomConnection()
	EmptyConnection()
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

var imgStr string = "iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAACXBIWXMAAAsTAAALEwEAmpwYAAADIUlEQVRoge3Wa0hTYRgH8EXSt6KskAxFBUPRJEnCLCTQEjRJNClT0jIMs4uK1ZfCbmRhGkqaKGgpVl4SyTSpTDMTTTJveSm3ubmbubOLbmam/tt5KTOYkLjpTp0H/rB353D2/OB9nx0Ohy222Pp/6tNGjnu/FQfGiDAqqESr1VoyFiBKiKxTqVR9RkUYEzAYG16nVqthcARFUVZKpdKWzkBUsB/XxkywkPTbmA3qBZwMJYCfiF6DIXQPk/x6sCGilIi0+gDiU4dqZ99HIzQazQbGAPju1k2SK/H1syM8HvzI4ABKIkZ3QQ668u5C1t0JWV8PPt7LQnd+NlnT93CrytGVm4meh3l/DdCXz1acKYMDxK0tKNlkjrb0GyhxWIfeonxU+Wwj61InC/BrqvE8yAsNccfQfifZRAG6xjtz0mcA9dGh5FrzhVg0no0mgDcxh9GRdds0AcX2q/EhLQnSzjb0lRbOAGrC/PE++TIBdGSmQNhQZ5qAMlfrmTUNKHOxROVuN9SE+EIhlRDAwOuX8zrEiwbQF7lQABVF/W5SLodapWIOwBBj1JiAp7ofbqajFItahl9VdS0k8ori1kUFzC5jvguxABYwR3jOa6BIvQT5tQRmAoYSjkLgYQvNiycQ7HJgHuDL+Sjwt6zH2NsaSMJ9mQcQBWwnmRyWQRS4k3kArt0KjDwuAKanwdtszjwAndHKEnzr7WTmIe63XoYJbh9Gyx8wEyAK8ABdI0W5EHo7Mw+gLszG19YmcG3NdGchHzzHVcwB8F3WYkozSkYpvRbssMNoRRHE+z0hjQrEUFw4ZKdDIdzjYpoAKjWRAOgmqVsXIb9+DsOJZzDWWEv+E+ipJPRygrb2GfnedAA2y0mDkwo5FGlXIfC0/+M6z3EllBlJUN3PwFB8BMbb3kGZedM0ALKYg2TfTwh50NZVkyk05xZztYD0iD8G/dz03re4AF0DyuwUsl0kYT6YpIbB37ph3qNzyQCyEwcgDvEm2+W7SABJxN4FNb/oAGnkPjImx9tbyOeFNr8kZ4A+nPQBNkTzSzuFWAAL+EcBbLHFlunVDzXOPAb1wxnfAAAAAElFTkSuQmCC"

/*
	Report without connection to database
*/
func EmptyConnection() {

	rpt, err := keireport.LoadFromFile("empty_conn.krpt")

	if err == nil {

		rpt.SetParam("name", "Koder Nubie")
		rpt.SetParam("positon", "Senior Golang Developer")

		byt, _ := base64.RawStdEncoding.DecodeString(imgStr)
		rpt.AddResource("logo.png", "image/png", byt)

		rpt.GenToFile("empty_conn.pdf")

		fmt.Println("generated to : empty_conn.pdf")
	} else {

		fmt.Println("Error : ", err)
	}
}
