package main

import (
	"encoding/json"
	"io/ioutil"
)

type PackageConfig struct {
	data map[string]interface{}
	path string
}

func NewPackageConfig(path string) PackageConfig {
	var ret PackageConfig
	if path == "" {
		path = "./cpackage.json"
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &ret.data)
	if err != nil {
		panic(err)
	}

	return ret
}

func (s *PackageConfig) getProjects() []string {
	var ret []string
	value, ok := s.data["projects"]
	if !ok {
		return nil
	}

	switch v := value.(type) {
	case map[string]interface{}:
		for i, _ := range v {
			ret = append(ret, i)
		}
	}

	return ret
}
