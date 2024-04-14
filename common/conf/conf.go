package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

func Unmarshal(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}
