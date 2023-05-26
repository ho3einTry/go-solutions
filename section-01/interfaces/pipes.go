package interfaces

import (
	"io"
	"os"
	"time"
)

func PipeExample() error {

	r, w := io.Pipe()
	go func() {

		w.Write([]byte("Start\n"))
		for i := 0; i < 100; i++ {

			time.Sleep(50 * time.Millisecond)
			w.Write([]byte("test\n"))

		}
		w.Write([]byte("End"))
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
