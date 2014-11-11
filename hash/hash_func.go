// hash_func
package hash

import "unsafe"

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
	p := (*int32)(unsafe.Pointer(&f))
	return HashInt32(code, *p)
}

func HashFloat64(code uint32, f float64) uint32 {
	p := (*int64)(unsafe.Pointer(&f))
	return HashInt64(code, *p)
}

func HashComplex64(code uint32, cpx complex64) uint32 {
	r := real(cpx)
	i := imag(cpx)
	p := (*int32)(unsafe.Pointer(&r))
	code = HashInt32(code, *p)
	p = (*int32)(unsafe.Pointer(&i))
	return HashInt32(code, *p)
}

func HashComplex128(code uint32, cpx complex128) uint32 {
	r := real(cpx)
	i := imag(cpx)
	p := (*int64)(unsafe.Pointer(&r))
	code = HashInt64(code, *p)
	p = (*int64)(unsafe.Pointer(&i))
	return HashInt64(code, *p)
}

func HashString(code uint32, s string) uint32 {
	rs := code
	chs := []rune(s)
	for _, ch := range chs {
		rs = 31*rs + uint32(ch)
	}
	return rs
}
