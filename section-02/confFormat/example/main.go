package example

import "github.com/ho3eintry/go-solutions/section-02/confFormat"

func main() {
	if err := confFormat.MarshalAll(); err != nil {
		panic(err)
	}

	if err := confFormat.UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := confFormat.JsonToMap(); err != nil {
		panic(err)
	}
}
