package utils

import (
	"reflect"
)

// Assign
func Assign(origin, target interface{}, excludes ...string) {
	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		if !val_target.FieldByName(val_origin.Type().Field(i).Name).IsValid() {
			continue
		}
		is_exclude := false
		for _, col := range excludes {
			if val_origin.Type().Field(i).Name == col {
				is_exclude = true
				break
			}
		}
		if is_exclude {
			continue
		}
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetInt(val_origin.Field(i).Int())
		case reflect.String:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetString(val_origin.Field(i).String())
		}
	}
}
