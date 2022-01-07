package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml"
)

type Cfg struct {
	// Name      string `yaml:name`
	// Auto      bool   `yaml:auto`
	// Port      int    `yaml:port`
	// Blackip   []string
	// Clusterip []string
	Logconf Logconf
}

type Logconf struct {
	Path string
	// Cmd     string
	Leve string
	// Timeout string
	// Disable bool
}

func main21212() {

	// file, err := os.Open("TestConf.yaml")

	// if err != nil {
	// 	panic(err)
	// }

	// bytes, err := ioutil.ReadAll(file)
	bytes, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}

	cfg := Cfg{}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
