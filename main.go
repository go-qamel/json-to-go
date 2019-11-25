package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"

	"github.com/buger/jsonparser"
)

// GeneratedName is the name for generated type
const GeneratedName = "AutoGenerated"

func main() {
	res, err := jsonToGoCode([]byte(smartyStreets), false)
	if err != nil {
		log.Fatalln(err)
	}

	bt, err := format.Source([]byte(res))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bt))
}

func jsonToGoCode(data []byte, inlineStruct bool) (string, error) {
	// Parse JSON value
	value, dataType, _, err := jsonparser.Get(data)
	if err != nil {
		return "", fmt.Errorf("failed to parse json: %w", err)
	}

	// If it's a primitive value, return as it is
	switch dataType {
	default:
		return "type " + GeneratedName + " interface{}", nil
	case jsonparser.String:
		return "type " + GeneratedName + " string", nil
	case jsonparser.Boolean:
		return "type " + GeneratedName + " bool", nil
	case jsonparser.Number:
		numberType := getNumberType(value)
		return "type " + GeneratedName + " " + numberType, nil
	case jsonparser.Array:
		arrayType := getArrayType(value)
		return "type " + GeneratedName + " " + arrayType, nil
	case jsonparser.Object:
	}

	// At this point, the json value is object that need additional care to convert.
	structDecl, subStructs, err := createStruct(GeneratedName, value, inlineStruct)
	if err != nil {
		return "", fmt.Errorf("failed to create struct: %w", err)
	}

	// If this is inline struct, add struct name to struct declaration
	if inlineStruct {
		structDecl = "type " + GeneratedName + " " + structDecl
	}

	// Merge struct declaration and sub structs to one string.
	// Map the sub structs to prevent duplicate.
	finalResult := structDecl
	mapSubStructs := make(map[string]struct{})
	for _, subStruct := range subStructs {
		if _, exist := mapSubStructs[subStruct]; exist {
			continue
		}

		mapSubStructs[subStruct] = struct{}{}
		finalResult += "\n\n" + subStruct
	}

	return finalResult, nil
}

func createStruct(name string, data []byte, inlineStruct bool) (string, []string, error) {
	var result string
	if inlineStruct {
		result = "struct {\n"
	} else {
		result = "type " + name + " struct {\n"
	}

	var subStructs []string
	handler := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		// Convert key to camel case
		jsonName := string(key)
		fieldName := toCamelCase(jsonName)

		// Get field type
		var fieldType string
		switch dataType {
		default:
			fieldType = "interface{}"
		case jsonparser.String:
			fieldType = "string"
		case jsonparser.Number:
			fieldType = getNumberType(value)
		case jsonparser.Array:
			fieldType = getArrayType(value)
		case jsonparser.Boolean:
			fieldType = "bool"
		case jsonparser.Object:
			structDecl, subs, err := createStruct(fieldName, value, inlineStruct)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", jsonName, err)
			}

			if inlineStruct {
				fieldType = structDecl
			} else {
				fieldType = fieldName
				subStructs = append(subStructs, structDecl)
				subStructs = append(subStructs, subs...)
			}
		}

		result += fmt.Sprintf("%s %s `json:\"%s\"`\n", fieldName, fieldType, jsonName)
		return nil
	}

	err := jsonparser.ObjectEach(data, handler)
	if err != nil {
		return "", nil, err
	}

	result += "}"
	return result, subStructs, nil
}

func getNumberType(value []byte) string {
	// If value contain decimal separator, it's float
	if bytes.ContainsRune(value, '.') {
		return "float64"
	}

	return "int"
}

func getArrayType(value []byte) string {
	// Get all type in this array
	mapType := make(map[string]struct{})
	handler := func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch dataType {
		default:
			mapType["interface{}"] = struct{}{}
		case jsonparser.String:
			mapType["string"] = struct{}{}
		case jsonparser.Number:
			numberType := getNumberType(value)
			mapType[numberType] = struct{}{}
		case jsonparser.Array:
			arrayType := getArrayType(value)
			mapType[arrayType] = struct{}{}
		case jsonparser.Boolean:
			mapType["bool"] = struct{}{}
		}
	}

	jsonparser.ArrayEach(value, handler)

	// Convert map to list
	var listType []string
	for tp := range mapType {
		listType = append(listType, tp)
	}

	// If there is only one type, return it
	if len(listType) == 1 {
		return "[]" + listType[0]
	}

	// If there are only two type, and both of them is number
	// (i.e int and float64), float64 win
	_, intExist := mapType["int"]
	_, floatExist := mapType["float64"]
	if len(listType) == 2 && intExist && floatExist {
		return "[]float64"
	}

	// When all fail, just return it as []interface{}
	return "[]interface{}"
}
