package pkg

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadFile(dest interface{}, filePath string) error {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, dest)
	if err != nil {
		return err
	}
	return nil
}
