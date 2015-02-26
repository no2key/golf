// conv
package conv

import (
	"unsafe"
)

func Convi(i uint) int {
	return *((*int)(unsafe.Pointer(&i)))
}

func Convi8(i uint8) int8 {
	return *((*int8)(unsafe.Pointer(&i)))
}

func Convi16(i uint16) int16 {
	return *((*int16)(unsafe.Pointer(&i)))
}

func Convi32(i uint32) int32 {
	return *((*int32)(unsafe.Pointer(&i)))
}

func Convi64(i uint64) int64 {
	return *((*int64)(unsafe.Pointer(&i)))
}
