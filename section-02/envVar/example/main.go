package main

import (
	"bytes"
	"fmt"
	"github.com/ho3eintry/go-solutions/section-02/envVar"
	"io/ioutil"
	"os"
)

type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error

	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	secrets := `{
        "secret": "so so secret"
    }`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	if err = os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err = envVar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))
	fmt.Printf("Final Config: %#v\n", c)
}
