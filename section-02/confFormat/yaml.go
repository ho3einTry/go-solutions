package confFormat

import (
	"bytes"
	"github.com/go-yaml/yaml"
)

type YamlData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func (t *YamlData) ToYaml() (*bytes.Buffer, error) {

	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

func (t *YamlData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}
