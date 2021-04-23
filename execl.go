package util

import (
	"bytes"
	"fmt"
	"go/ast"
	"reflect"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

/**
 * ExportExcel 生成excel
 *
 * Example:
 *
 * type ReportExcel struct {
 * 	Dt         string `excelize:"head:日期;"`
 * 	NewUsers   int64  `excelize:"head:注册人数;"`
 * 	LoginUsers int64  `excelize:"head:登录人数;"`
 * }
 * tbody := []ReportExcel{
 * 	{"2006-01-02", 1, 2},
 * }
 * buf, err := ExportExcel(tbody)
 * if err != nil {
 * 	//
 * }
 * ioutil.WriteFile("test.txt", buf.Bytes(), 0644)
 */

func ExportExcel(tbody interface{}) (r *bytes.Buffer, err error) {
	// make sure 'tbody' is a Slice
	val_tbody := reflect.ValueOf(tbody)
	if val_tbody.Kind() != reflect.Slice {
		err = fmt.Errorf("tbody not slice")
		return
	}
	var (
		// it's a slice, so open up its values
		n     = val_tbody.Len()
		thead = []string{}
		sheet = "Sheet1"
	)

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet(sheet)

	// write tbody
	for i := 0; i < n; i++ {
		v := val_tbody.Index(i)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			continue
		}
		num_field := v.NumField() // number of fields in struct

		for j := 0; j < num_field; j++ {
			var axis string
			axis, err = excelize.CoordinatesToCellName(j+1, i+2)
			if err != nil {
				err = fmt.Errorf("tbody to cell name error: %s", err)
				return
			}
			f := v.Field(j)
			t := v.Type().Field(j)

			if !ast.IsExported(t.Name) {
				continue
			}
			tags := ParseTagSetting(t.Tag, "excelize")
			// is ignored field
			if _, ok := tags["-"]; ok {
				continue
			}
			if i == 0 {
				head := t.Name
				if t, ok := tags["HEAD"]; ok && len(head) > 0 {
					head = t
				}
				thead = append(thead, head)
			}
			xlsx.SetCellValue(sheet, axis, f.Interface())
		}
	}
	// write thead
	for k, v := range thead {
		var axis string
		axis, err = excelize.CoordinatesToCellName(k+1, 1)
		if err != nil {
			err = fmt.Errorf("thead to cell name error: %s", err)
			return
		}
		xlsx.SetCellValue(sheet, axis, v)
	}
	xlsx.SetActiveSheet(index)

	r, err = xlsx.WriteToBuffer()
	if err != nil {
		err = fmt.Errorf("write to buffer error: %s", err)
		return
	}
	return
}
