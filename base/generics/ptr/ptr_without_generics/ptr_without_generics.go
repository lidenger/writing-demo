package ptr_without_generics

import "reflect"

func Ptr(v any) any {
	rv := reflect.New(reflect.TypeOf(v))
	rv.Elem().Set(reflect.ValueOf(v))
	return rv.Interface()
}

func PtrInt8(v int8) *int8 {
	return &v
}

func PtrInt16(v int16) *int16 {
	return &v
}

func PtrInt32(v int32) *int32 {
	return &v
}

func PtrInt(v int) *int {
	return &v
}

func PtrInt64(v int64) *int64 {
	return &v
}

func PtrFloat32(v float32) *float32 {
	return &v
}

func PtrFloat64(v float32) *float32 {
	return &v
}

func PtrString(v string) *string {
	return &v
}

func PtrBool(v bool) *bool {
	return &v
}
