package main

import (
	"os"

	_ "json-to-go/internal/backend"
)

func main() {
	runQtApp(len(os.Args), os.Args)
}
