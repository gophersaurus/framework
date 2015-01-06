package mockstar

import (
	"fmt"
	"reflect"
	"time"
)

func (a *Args) Bool(index int) bool {
	val := a.Get(index)
	if reflect.DeepEqual(val, true) {
		return true
	} else if reflect.DeepEqual(val, false) {
		return false
	} else {
		Fatal("cannot cast value to boolean: " + fmt.Sprintf("%v", val))
	}
	return false
}

func (a *Args) Byte(index int) byte {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(byte)
	if !ok {
		Fatal("cannot cast value to byte: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Bytes(index int) []byte {
	val := a.Get(index)
	if val == nil {
		return nil
	}
	out, ok := val.([]byte)
	if !ok {
		Fatal("cannot cast value to []byte: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Complex64(index int) complex64 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(complex64)
	if !ok {
		Fatal("cannot cast value to complex64: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Complex128(index int) complex128 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(complex128)
	if !ok {
		Fatal("cannot cast value to complex128: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Err(index int) error {
	val := a.Get(index)
	if val == nil {
		return nil
	}
	out, ok := val.(error)
	if !ok {
		Fatal("cannot cast value to error: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Float32(index int) float32 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(float32)
	if !ok {
		Fatal("cannot cast value to float32: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Float64(index int) float64 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(float64)
	if !ok {
		Fatal("cannot cast value to float64: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) String(index int) string {
	val := a.Get(index)
	if val == nil {
		return ""
	}
	out, ok := val.(string)
	if !ok {
		Fatal("cannot cast value to string: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Strings(index int) []string {
	val := a.Get(index)
	if val == nil {
		return nil
	}
	out, ok := val.([]string)
	if !ok {
		Fatal("cannot cast value to []string: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Int(index int) int {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(int)
	if !ok {
		Fatal("cannot cast value to int: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Int8(index int) int8 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(int8)
	if !ok {
		Fatal("cannot cast value to int8: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Int16(index int) int16 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(int16)
	if !ok {
		Fatal("cannot cast value to int16: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Int32(index int) int32 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(int32)
	if !ok {
		Fatal("cannot cast value to int32: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Int64(index int) int64 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(int64)
	if !ok {
		Fatal("cannot cast value to int64: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uint(index int) uint {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uint)
	if !ok {
		Fatal("cannot cast value to uint: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uint8(index int) uint8 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uint8)
	if !ok {
		Fatal("cannot cast value to uint8: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uint16(index int) uint16 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uint16)
	if !ok {
		Fatal("cannot cast value to uint16: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uint32(index int) uint32 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uint32)
	if !ok {
		Fatal("cannot cast value to uint32: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uint64(index int) uint64 {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uint64)
	if !ok {
		Fatal("cannot cast value to uint64: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Uintptr(index int) uintptr {
	val := a.Get(index)
	if val == nil {
		return 0
	}
	out, ok := val.(uintptr)
	if !ok {
		Fatal("cannot cast value to uintptr: " + fmt.Sprintf("%v", val))
	}
	return out
}

func (a *Args) Time(index int) *time.Time {
	val := a.Get(index)
	if val == nil {
		return nil
	}
	out, ok := val.(time.Time)
	if !ok {
		Fatal("cannot cast value to Time: " + fmt.Sprintf("%v", val))
	}
	return &out
}
