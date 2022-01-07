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

func init() {
	ReadingConfiguration()
}

func ReadingConfiguration() Conf {
	bytes, err := ioutil.ReadFile(CONFIGURATION)
	if err != nil {
		panic(err)
	}
	conf := Conf{}
	err = yaml.Unmarshal(bytes, &conf)

	if err != nil {
		panic(err)
	}

	fmt.Println(conf.Logconf.Leve)
	return conf
}

func getLevel(conf *Conf) int8 {
	return conf.Logconf.Leve
}
