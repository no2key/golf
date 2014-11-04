// hash
package hash

import "errors"

type Hash interface {
	HashCode() uint32
	Equals(Hash) bool
}

var (
	paramIsNilError = errors.New("the param is nil")
)

const (
	_MINIMUM_CAPACITY = 4
	_MAXIMUM_CAPACITY = 1 << 30
	_MIN_FLOAT_VALUE_ = 0.000000001
)
