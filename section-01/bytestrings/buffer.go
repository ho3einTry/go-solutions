package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {
	rawbytes := []byte(rawString)
	var b = new(bytes.Buffer)
	b.Write(rawbytes)

	b = bytes.NewBuffer(rawbytes)
	b = bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
