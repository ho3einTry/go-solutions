package confFormat

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JsonData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t *JsonData) ToJson() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

func (t *JsonData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

func JsonToMap() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}
	fmt.Println("We can unmarshal into a map instead of a struct:", res)
	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams:", res)
	return nil
}
