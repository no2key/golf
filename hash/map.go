// map
package hash

type MapEntry interface {
	Key() Hash
	Value() interface{}
}

type Map interface {
	Put(h Hash, v interface{}) interface{}

	Get(h Hash) interface{}

	Remove(h Hash) interface{}

	Contains(h Hash) bool

	Size() int
	IsEmpty() bool

	Clear()

	Travel(func(MapEntry) bool)
}
