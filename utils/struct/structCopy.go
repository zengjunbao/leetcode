package _struct

import (
	"encoding/json"
	"errors"
	"reflect"
)

type OrderA struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type OrderB struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Code    string `json:"code"`
}

// CopyStruct oldObj 和 newObj 需要传入指针，通过json tag对应
func CopyStruct2(oldObj interface{}, newObj interface{}) error {
	data,err := json.Marshal(oldObj)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data,newObj); err != nil {
		return err
	}
	return nil
}

// StructCopy Struct拷贝
// DstStructPtr type interface  目标结构体
// SrcStructPtr type interface  源结构体
func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) error {
	srcV := reflect.ValueOf(SrcStructPtr)
	dstV := reflect.ValueOf(DstStructPtr)
	srcT := reflect.TypeOf(SrcStructPtr)
	dstT := reflect.TypeOf(DstStructPtr)
	if srcT.Kind() != reflect.Ptr || dstT.Kind() != reflect.Ptr ||
		srcT.Elem().Kind() == reflect.Ptr || dstT.Elem().Kind() == reflect.Ptr {
		return errors.New("参数类型必须是指针类型")
	}
	if srcV.IsNil() || dstV.IsNil() {
		return errors.New("参数不能为nil")
	}
	srcVv := srcV.Elem()
	dstVv := dstV.Elem()
	srcFields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcFields {
		if v.Anonymous {
			continue
		}
		dst := dstVv.FieldByName(v.Name)
		src := srcVv.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return nil
}

// DeepFields 深度遍历
func DeepFields(iFaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < iFaceType.NumField(); i++ {
		v := iFaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}