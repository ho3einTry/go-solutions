package confFormat

import "fmt"

func MarshalAll() error {
	t := TomlData{
		Name: "Name1",
		Age:  20,
	}

	j := JsonData{
		Name: "Name2",
		Age:  30,
	}

	y := YamlData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.ToToml()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJson()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYaml()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}
