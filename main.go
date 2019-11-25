package main

import (
	"fmt"
	"log"

	"json-2-go/converter"
)

func main() {
	cv := converter.Converter{InlineStruct: true}
	res, err := cv.Convert(sample)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
