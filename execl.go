package util

import (
	"bytes"
	"fmt"
	"go/ast"
	"reflect"

	"github.com/xuri/excelize/v2"
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
 * 	Tmp        int64  `excelize:"-"`
 * }
 * tbody := []ReportExcel{
 * 	{"2006-01-02", 1, 2},
 * }
 * buf, err := ExportExcel(ReportExcel{}, tbody)
 * if err != nil {
 * 	//
 * }
 * ioutil.WriteFile("test.xlsx", buf.Bytes(), 0644)
 */
func ExportExcel(table interface{}, tbody interface{}) (*bytes.Buffer, error) {
	// make sure 'table' is a Struct
	tableVal := reflect.ValueOf(table)
	if tableVal.Kind() == reflect.Ptr {
		tableVal = tableVal.Elem()
	}
	if tableVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("table not struct")
	}
	// make sure 'tbody' is a Slice
	tbodyVal := reflect.ValueOf(tbody)
	if tbodyVal.Kind() != reflect.Slice {
		return nil, fmt.Errorf("tbody not slice")
	}
	var (
		tableFieldNum = tableVal.NumField()
		tbodyLen      = tbodyVal.Len()
		sheet         = "Sheet1"
		headCol       = 1
	)

	xlsx := excelize.NewFile()
	index, err := xlsx.NewSheet(sheet)
	if err != nil {
		return nil, fmt.Errorf("new sheet error: %s", err)
	}

	// write table head
	for i := 0; i < tableFieldNum; i++ {
		t := tableVal.Type().Field(i)

		if !ast.IsExported(t.Name) {
			continue
		}
		tags := ParseTagSetting(t.Tag, "excelize")
		// is ignored field
		if _, ok := tags["-"]; ok {
			continue
		}
		head := t.Name
		if t, ok := tags["HEAD"]; ok && len(head) > 0 {
			head = t
		}
		axis, err := excelize.CoordinatesToCellName(headCol, 1)
		if err != nil {
			return nil, fmt.Errorf("head to cell name error: %s", err)
		}
		headCol++
		xlsx.SetCellValue(sheet, axis, head)
	}

	// write tbody
	for i := 0; i < tbodyLen; i++ {
		v := tbodyVal.Index(i)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("tbody element not struct")
		}
		num_field := v.NumField() // number of fields in struct
		bodyCol := 1

		for j := 0; j < num_field; j++ {
			axis, err := excelize.CoordinatesToCellName(bodyCol, i+2)
			if err != nil {
				return nil, fmt.Errorf("tbody to cell name error: %s", err)
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
			bodyCol++
			xlsx.SetCellValue(sheet, axis, f.Interface())
		}
	}
	xlsx.SetActiveSheet(index)

	buf, err := xlsx.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("write to buffer error: %s", err)
	}
	return buf, nil
}
