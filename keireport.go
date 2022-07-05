package keireport

import (
	_ "github.com/kodernubie/keireport/component"
	"github.com/kodernubie/keireport/core"
	_ "github.com/kodernubie/keireport/datasource"
	_ "github.com/kodernubie/keireport/exporter"
)

func LoadFromFile(fileName string) (*core.Keireport, error) {

	ret := &core.Keireport{}
	err := ret.LoadFromFile(fileName)

	return ret, err
}

func LoadFromString(templateString, baseDir string) (*core.Keireport, error) {

	ret := &core.Keireport{}
	err := ret.LoadFromString(templateString, baseDir)

	return ret, err
}
