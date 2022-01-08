package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml"
)

type Conf struct {
	Logconf Logconf
}

type Logconf struct {
	Path string
	Leve int8
}

const (
	CONFIGURATION = "conf.yaml"
)

// func init() {
// 	conf := Conf{}
// 	conf.ReadingConfiguration()
// }

func (cfg *Conf) ReadingConfiguration() *Conf {
	bytes, err := ioutil.ReadFile(CONFIGURATION)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, cfg)

	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Logconf.Leve)
	return cfg
}
