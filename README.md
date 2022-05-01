# KeiReport
Golang based reporting engine. Generating PDF, HTML, Docx and more using simple template. Getting data from database, csv and more

## Features
  
  - JSON based template
  - Simple band system : title - header - detail - footer
  - Datasource : DB, CSV
  - Label, Rect, Image
  - Parameterized value in text, image

View our roadmap for upcoming feature

## Related Project

  - [KeiReport Designer](https://github.com/kodernubie/keireport-designer)
  - [Keireport Server](https://github.com/kodernubie/keireport-server)

## Installation

To install the package to your system, run :

```
go get github.com/kodernubie/keireport
```

## Quick Start

Following code load report template from file and generate to pdf

```
import "github.com/kodernubie/keireport"

...

rpt, err := keireport.LoadFromFile("simple.krpt")

if err == nil {

    rpt.GenToFile("simple.pdf")
}

```

For more usage examples you can check "example" directory  

## License

`kodernubie/keireport` is released under the MIT License.

## Roadmap

- Generate to DocX
- Specialized template for spreadsheet
- Barcode, QR code