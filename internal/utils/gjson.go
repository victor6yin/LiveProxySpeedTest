package utils

import "github.com/tidwall/gjson"

func GJsonGetDefault(getValue gjson.Result, defaultValue interface{}) string {
	if !getValue.Exists() {
		switch value := defaultValue.(type) {
		case string:
			return value
		case gjson.Result:
			return value.String()
		}
	}
	return getValue.String()
}
