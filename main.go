package main

import (
	"fmt"
	"log"
)

func main() {
	cv := Converter{InlineStruct: false}
	res, err := cv.Convert(smartyStreets)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
