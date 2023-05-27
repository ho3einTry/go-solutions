package temp_files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() error {
	t, err := ioutil.TempDir("/home/hossein", "hossein")
	if err != nil {
		return err
	}
	// os temp directory
	fmt.Println("temp directory: ", os.TempDir())

	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "hosseinFile")
	fmt.Println(tf.Name())
	tf2, err := ioutil.TempFile(t, "parvinFile*.log")
	fmt.Println(tf2.Name())
	return err
}
