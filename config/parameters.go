package config

import (
	"github.com/go-ini/ini"
	"fmt"
)

var Parameters ParametersStruct

func init() {
	config, err := ini.Load("parameters.ini")

	if (err != nil) {
		panic("parameters.ini does not exist.")
	}

	Parameters = ParametersStruct{
		config: config,
	}
}

type ParametersStruct struct {
	config *ini.File
}

func (parameters *ParametersStruct) GetDsn() string {
	mysqlSection, err := parameters.config.GetSection("mysql")

	if (err != nil) {
		panic("No database section in config.")
	}

	return fmt.Sprintf("%s:%s@/%s", mysqlSection.Key("username"), mysqlSection.Key("passoword"), mysqlSection.Key("database"))
}

func (parameters *ParametersStruct) GetGoogleKey() string {
	google, err := parameters.config.GetSection("google")

	if (err != nil) {
		panic("Google credentials not set.")
	}

	return google.Key("key").Value()
}
