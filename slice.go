package util

import (
	"math/rand"
	"reflect"
	"time"
)

// InSlice
func InSlice(val interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == false {
			continue
		}
		index = i
		exists = true
		return
	}
	return
}

// SliceDiff
func SliceDiff(slice1, slice2 interface{}) (r []interface{}) {
	if reflect.TypeOf(slice1).Kind() != reflect.Slice || reflect.TypeOf(slice2).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice1)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice2); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	s = reflect.ValueOf(slice2)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice1); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

// SliceUnique
func SliceUnique(slice interface{}) (r []interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), r); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

// SliceUniqueInt
func SliceUniqueInt(slice []int) (r []int) {
	for _, v := range slice {
		if exists, _ := InSlice(v, r); exists {
			continue
		}
		r = append(r, v)
	}
	return
}

// SliceUniqueInt32
func SliceUniqueInt32(slice []int32) (r []int32) {
	for _, v := range slice {
		if exists, _ := InSlice(v, r); exists {
			continue
		}
		r = append(r, v)
	}
	return
}

// SliceShuffle
func SliceShuffle(slice interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	var ran = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := s.Len() - 1; i > 0; i-- {
		k := ran.Intn(i + 1)
		tmp := s.Index(k).Interface()
		s.Index(k).Set(s.Index(i))
		s.Index(i).Set(reflect.ValueOf(tmp))
	}
	return
}
