package util

import (
	"math/rand"
	"reflect"
	"time"
)

// InSlice 函数搜索数组中是否存在指定的值
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

// InSliceF 函数搜索数组中是否存在指定的值
func InSliceF[T comparable](val T, slice []T) (bool, int) {
	for k, v := range slice {
		if val == v {
			return true, k
		}
	}
	return false, -1
}

// SliceDiff 函数用于比较两个数组的值，并返回差集
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

// SliceUnique 函数用于移除数组中重复的值
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

// SliceShuffle 函数把数组中的元素按随机顺序重新排列
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

// SliceChunk 函数把一个数组分割为新的数组块
func SliceChunk(slice []int, size int) (r [][]int) {
	for size < len(slice) {
		r, slice = append(r, slice[0:size:size]), slice[size:]
	}
	r = append(r, slice)
	return
}

// SliceChunkInt64
func SliceChunkInt64(slice []int64, size int) (r [][]int64) {
	for size < len(slice) {
		r, slice = append(r, slice[0:size:size]), slice[size:]
	}
	r = append(r, slice)
	return
}
