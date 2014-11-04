// hash_type
package hash

// int
type IntHash int

func (this IntHash) HashCode() uint32 {
	return HashInt(5, int(this))
}

func (this IntHash) Equals(h Hash) bool {
	if v, ok := h.(IntHash); ok {
		return this == v
	}
	return false
}

func Int(n int) Hash {
	return IntHash(n)
}

// uint
type UIntHash uint

func (this UIntHash) HashCode() uint32 {
	return HashUInt(5, uint(this))
}

func (this UIntHash) Equals(h Hash) bool {
	if v, ok := h.(UIntHash); ok {
		return this == v
	}
	return false
}

func Uint(n uint) Hash {
	return UIntHash(n)
}

// uintptr
type UIntPtrHash uintptr

func (this UIntPtrHash) HashCode() uint32 {
	return HashUIntPtr(5, uintptr(this))
}

func (this UIntPtrHash) Equals(h Hash) bool {
	if v, ok := h.(UIntPtrHash); ok {
		return this == v
	}
	return false
}

func Uintptr(p uintptr) Hash {
	return UIntPtrHash(p)
}

// bool
type BoolHash bool

func (this BoolHash) HashCode() uint32 {
	return HashBool(5, bool(this))
}

func (this BoolHash) Equals(h Hash) bool {
	if v, ok := h.(BoolHash); ok {
		return this == v
	}
	return false
}

func Bool(b bool) Hash {
	return BoolHash(b)
}

// int8
type Int8Hash int8

func (this Int8Hash) HashCode() uint32 {
	return HashInt8(5, int8(this))
}

func (this Int8Hash) Equals(h Hash) bool {
	if v, ok := h.(Int8Hash); ok {
		return this == v
	}
	return false
}

func Int8(i int8) Hash {
	return Int8Hash(i)
}

// uint8
type UInt8Hash uint8

func (this UInt8Hash) HashCode() uint32 {
	return HashUInt8(5, uint8(this))
}

func (this UInt8Hash) Equals(h Hash) bool {
	if v, ok := h.(UInt8Hash); ok {
		return this == v
	}
	return false
}

func Uint8(n uint8) Hash {
	return UInt8Hash(n)
}

// int16
type Int16Hash int16

func (this Int16Hash) HashCode() uint32 {
	return HashInt16(5, int16(this))
}

func (this Int16Hash) Equals(h Hash) bool {
	if v, ok := h.(Int16Hash); ok {
		return this == v
	}
	return false
}

func Int16(n int16) Hash {
	return Int16Hash(n)
}

// uint16
type UInt16Hash uint16

func (this UInt16Hash) HashCode() uint32 {
	return HashUInt16(5, uint16(this))
}

func (this UInt16Hash) Equals(h Hash) bool {
	if v, ok := h.(UInt16Hash); ok {
		return this == v
	}
	return false
}

func Uint16(n uint16) Hash {
	return UInt16Hash(n)
}

// int32
type Int32Hash int32

func (this Int32Hash) HashCode() uint32 {
	return HashInt32(5, int32(this))
}

func (this Int32Hash) Equals(h Hash) bool {
	if v, ok := h.(Int32Hash); ok {
		return this == v
	}
	return false
}

func Int32(n int32) Hash {
	return Int32Hash(n)
}

// uint32
type UInt32Hash uint32

func (this UInt32Hash) HashCode() uint32 {
	return HashUInt32(5, uint32(this))
}

func (this UInt32Hash) Equals(h Hash) bool {
	if v, ok := h.(UInt32Hash); ok {
		return this == v
	}
	return false
}

func Uint32(n uint32) Hash {
	return UInt32Hash(n)
}

// int64
type Int64Hash int64

func (this Int64Hash) HashCode() uint32 {
	return HashInt64(5, int64(this))
}

func (this Int64Hash) Equals(h Hash) bool {
	if v, ok := h.(Int64Hash); ok {
		return this == v
	}
	return false
}

func Int64(n int64) Hash {
	return Int64Hash(n)
}

// uint64
type UInt64Hash uint64

func (this UInt64Hash) HashCode() uint32 {
	return HashUInt64(5, uint64(this))
}

func (this UInt64Hash) Equals(h Hash) bool {
	if v, ok := h.(UInt64Hash); ok {
		return this == v
	}
	return false
}

func Uint64(n uint64) Hash {
	return UInt64Hash(n)
}

// float32
type Float32Hash float32

func (this Float32Hash) HashCode() uint32 {
	return HashFloat32(5, float32(this))
}

func (this Float32Hash) Equals(h Hash) bool {
	if v, ok := h.(Float32Hash); ok {
		return this-v < _MIN_FLOAT_VALUE_ && this-v > -_MIN_FLOAT_VALUE_
	}
	return false
}

func Float32(n float32) Hash {
	return Float32Hash(n)
}

// float64
type Float64Hash float64

func (this Float64Hash) HashCode() uint32 {
	return HashFloat64(5, float64(this))
}

func (this Float64Hash) Equals(h Hash) bool {
	if v, ok := h.(Float64Hash); ok {
		return this-v < _MIN_FLOAT_VALUE_ && this-v > -_MIN_FLOAT_VALUE_
	}
	return false
}

func Float64(n float64) Hash {
	return Float64Hash(n)
}
