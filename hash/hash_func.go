// hash_func
package hash

import (
	"fmt"
)

func HashInt(code uint32, n int) uint32 {
	return 53*code + uint32(n)
}

func HashUInt(code uint32, n uint) uint32 {
	return 53*code + uint32(n)
}

func HashUIntPtr(code uint32, p uintptr) uint32 {
	return 53*code + uint32(p)
}

func HashBool(code uint32, b bool) uint32 {
	if b {
		return 53*code + 1
	} else {
		return 53*code + 0
	}
}

func HashInt8(code uint32, n int8) uint32 {
	return 53*code + uint32(n)
}

func HashUInt8(code uint32, n uint8) uint32 {
	return 53*code + uint32(n)
}

func HashInt16(code uint32, n int16) uint32 {
	return 53*code + uint32(n)
}

func HashUInt16(code uint32, n uint16) uint32 {
	return 53*code + uint32(n)
}

func HashInt32(code uint32, n int32) uint32 {
	return 53*code + uint32(n)
}

func HashUInt32(code uint32, n uint32) uint32 {
	return 53*code + n
}

func HashInt64(code uint32, n int64) uint32 {
	return 53*code + uint32(uint64(n)^(uint64(n)>>32))
}

func HashUInt64(code uint32, n uint64) uint32 {
	return 53*code + uint32(n^(n>>32))
}

func HashFloat32(code uint32, f float32) uint32 {
	s := fmt.Sprintf("%f", f)
	return HashString(code, s)
}

func HashFloat64(code uint32, f float64) uint32 {
	s := fmt.Sprintf("%f", f)
	return HashString(code, s)
}

func HashComplex64(code uint32, cpx complex64) uint32 {
	s := fmt.Sprintf("%f", real(cpx))
	code = HashString(code, s)
	s = fmt.Sprintf("%f", imag(cpx))
	return HashString(code, s)
}

func HashComplex128(code uint32, cpx complex128) uint32 {
	s := fmt.Sprintf("%f", real(cpx))
	code = HashString(code, s)
	s = fmt.Sprintf("%f", imag(cpx))
	return HashString(code, s)
}

func HashString(code uint32, s string) uint32 {
	rs := code
	chs := []rune(s)
	for _, ch := range chs {
		rs = 31*rs + uint32(ch)
	}
	return rs
}
