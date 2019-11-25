package main

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/buger/jsonparser"
)

// GeneratedName is the default name for generated type
const GeneratedName = "AutoGenerated"

// Converter is object to convert JSON to Go type
type Converter struct {
	GeneratedName string
	InlineStruct  bool
}

// Convert runs the converter.
func (cv Converter) Convert(jsonValue string) (string, error) {
	// Set default name
	generatedName := cv.GeneratedName
	if generatedName == "" {
		generatedName = GeneratedName
	}

	// Parse JSON value
	value, dataType, _, err := jsonparser.Get([]byte(jsonValue))
	if err != nil {
		return "", fmt.Errorf("failed to parse json: %w", err)
	}

	// Convert value depending on type
	var result string
	switch dataType {
	case jsonparser.String:
		result = "type " + generatedName + " string"
	case jsonparser.Boolean:
		result = "type " + generatedName + " bool"
	case jsonparser.Number:
		numberType := cv.getNumberType(value)
		result = "type " + generatedName + " " + numberType
	case jsonparser.Array:
		arrayType := cv.getArrayType("", value)
		result = "type " + generatedName + " " + arrayType
	case jsonparser.Object:
		structDecl, err := cv.createStruct(value)
		if err != nil {
			return "", fmt.Errorf("failed to create struct: %w", err)
		}
		result = "type " + generatedName + " " + structDecl
	default:
		return "type " + generatedName + " interface{}", nil
	}

	// Format the result
	result, err = cv.formatCode(result)
	if err != nil {
		return "", fmt.Errorf("failed to format output: %w", err)
	}

	return result, nil
}

func (cv Converter) createStruct(data []byte) (string, error) {
	result := "struct {\n"
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
			fieldType = cv.getNumberType(value)
		case jsonparser.Array:
			fieldType = cv.getArrayType(fieldName, value)
		case jsonparser.Boolean:
			fieldType = "bool"
		case jsonparser.Object:
			if !cv.InlineStruct {
				fieldType = fieldName
			} else {
				structDecl, err := cv.createStruct(value)
				if err != nil {
					return fmt.Errorf("failed to parse %s: %w", jsonName, err)
				}
				fieldType = structDecl
			}
		}

		result += fmt.Sprintf("%s %s `json:\"%s\"`\n", fieldName, fieldType, jsonName)
		return nil
	}

	err := jsonparser.ObjectEach(data, handler)
	if err != nil {
		return "", err
	}

	result += "}"
	return result, nil
}

func (cv Converter) getNumberType(value []byte) string {
	// If value contain decimal separator, it's float
	if bytes.ContainsRune(value, '.') {
		return "float64"
	}
	return "int"
}

func (cv Converter) getArrayType(name string, value []byte) string {
	// Get all type in this array
	mapType := make(map[string]struct{})
	handler := func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch dataType {
		default:
			mapType["interface{}"] = struct{}{}
		case jsonparser.String:
			mapType["string"] = struct{}{}
		case jsonparser.Number:
			numberType := cv.getNumberType(value)
			mapType[numberType] = struct{}{}
		case jsonparser.Array:
			arrayType := cv.getArrayType(name, value)
			mapType[arrayType] = struct{}{}
		case jsonparser.Boolean:
			mapType["bool"] = struct{}{}
		case jsonparser.Object:
			if !cv.InlineStruct && name != "" {
				mapType[name] = struct{}{}
			} else {
				structDecl, err1 := cv.createStruct(value)
				if err1 != nil {
					err = err1
				}

				mapType[structDecl] = struct{}{}
			}
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

func (cv Converter) formatCode(src string) (string, error) {
	bt, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}

	return string(bt), nil
}
