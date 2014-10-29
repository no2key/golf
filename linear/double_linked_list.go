// linked_list
package linear

type DLList interface {
	AddTail(v interface{})
	AddHead(v interface{})

	PeekTail() interface{}
	PeekHead() interface{}

	RemoveTail() interface{}
	RemoveHead() interface{}

	Clear()
	Size() int
	IsEmpty() bool

	Contains(v interface{}) bool

	Iterator(fn func(interface{}) bool)
	ReverseIterator(fn func(interface{}) bool)
}
