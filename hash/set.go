// map
package hash

type Set interface {
	Size() int
	IsEmpty() bool
	Contains(Hash) bool

	Add(Hash) bool
	Remove(Hash) bool
	Clear()

	Travel(func(Hash) bool)
}
