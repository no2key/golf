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

	Travel(fn func(interface{}) bool)
	ReverseTravel(fn func(interface{}) bool)
}
