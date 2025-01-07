package config

import (
	"github.com/bsthun/gut"
	"gopkg.in/yaml.v3"
	"os"
	"project1/common"
	"project1/type/cmn"
)

func Init() {
	// * declare struct
	config := new(cmn.Config)

	// * load configurations to struct
	yml, err := os.ReadFile("config.yml")
	if err != nil {
		gut.Fatal("unable to read configuration file", err)
	}
	if err := yaml.Unmarshal(yml, config); err != nil {
		gut.Fatal("unable to parse configuration file", err)
	}

	// * validate configurations
	if err := gut.Validator.Struct(config); err != nil {
		gut.Fatal("invalid configuration file", err)
	}

	// * Set global config
	common.Config = config
}
