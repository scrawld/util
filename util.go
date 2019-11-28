package util

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
		tmp_origin := val_origin.Field(i)
		tmp_target := val_target.FieldByName(val_origin.Type().Field(i).Name)
		if reflect.TypeOf(tmp_origin.Interface()) != reflect.TypeOf(tmp_target.Interface()) {
			continue
		}
		tmp_target.Set(tmp_origin)
	}
}
