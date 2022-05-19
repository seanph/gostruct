package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

const (
	structHdr  = "type NewObject struct {\n"
	structLine = "\t%s\t\t%s `json:\"%s\"`\n"
	structEnd  = "}\n"
)

func attemptJsonDecode(b []byte) (map[string]interface{}, error) {
	obj := map[string]interface{}{}
	// TODO Consider using Decode instead directly on the input streams.
	err := json.Unmarshal(b, &obj)
	return obj, err
}

// TODO Fine for now, might be nice to support structures other than
// `map[string]interface`.
func mapToStructDef(obj map[string]interface{}) string {
	output := structHdr
	titleCaser := cases.Title(language.Und)

	for jsonFieldName, fieldValue := range obj {
		goFieldName := ""

		// Convert json_snake_case to GoTitleCase.
		// TODO Handle other separators like '-'.
		parts := strings.Split(jsonFieldName, "_")
		for i := range parts {
			goFieldName += titleCaser.String(parts[i])
		}

		// Default to `interface{}` for `nil` fields.
		fieldType := "interface{}"
		if fieldValue != nil {
			// Remove spaces since `TypeOf().String()` will return for example
			// `map[string]interface {}` otherwise.
			fieldType = strings.Replace(reflect.TypeOf(fieldValue).String(), " ", "", -1)
		}

		// TODO Handle alternative struct tags.
		// TODO Better formatting.
		output += fmt.Sprintf(structLine, goFieldName, fieldType, jsonFieldName)
	}

	return output + structEnd
}
