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

	IndexOf(v interface{}) int
	LastIndexOf(v interface{}) int

	Sub(start, end int) (DLList, error)

	ToSlice(s []interface{}) []interface{}
}
