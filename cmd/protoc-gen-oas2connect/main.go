package main

import (
	"os"

	o2c "github.com/handlename/protoc-gen-oas2connect"
)

func main() {
	os.Exit(o2c.Run())
}
