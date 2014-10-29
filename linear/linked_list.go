// linked_list
package linear

type LinkedList interface {
	AddTail(v interface{})
	AddHead(v interface{})

	RemoveTail() interface{}
	RemoveHead() interface{}

	Clear()
	Size() int
	IsEmpty() bool

	Contains(v interface{}) bool

	IndexOf(v interface{}) int
	LastIndexOf(v interface{}) int

	Sub(start, end int) (LinkedList, error)

	ToSlice(s []interface{}) []interface{}
}
