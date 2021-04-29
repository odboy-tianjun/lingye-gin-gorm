package util

import (
	"lingye-gin/src/base"
	"reflect"
	"strconv"
)

// 修正分页参数
func FixPage(page uint, pageSize uint) (p uint, ps uint) {
	if page == 0 {
		p = 1
	}
	if pageSize == 0 {
		ps = 15
	}
	return p, ps
}

// source  有数据的结构体
// target  空的结构体
func StructCopy(source interface{}, target interface{}) {
	sourceVal := reflect.ValueOf(source).Elem() // 获取reflect.Type类型
	targetVal := reflect.ValueOf(target).Elem() // 获取reflect.Type类型
	vTypeOfT := sourceVal.Type()
	for i := 0; i < sourceVal.NumField(); i++ {
		// 在要空的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := targetVal.FieldByName(name).IsValid(); ok {
			targetVal.FieldByName(name).Set(reflect.ValueOf(sourceVal.Field(i).Interface()))
		}
	}
}

// 字符串转uint
func StringToUInt64(number string) uint64 {
	result, _ := strconv.Atoi(number)
	return uint64(result)
}

// 构成分页对象
func BuildPageData(data interface{}, total int) base.SimplePageData {
	var pageData base.SimplePageData
	pageData.Data = data
	pageData.Total = total
	return pageData
}
