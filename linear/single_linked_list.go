package linear

type SLList interface {
	IsEmpty() bool
	Clear()
	Size() int
	Add(v interface{})
	Remove() interface{}
	Peek() interface{}
	Contains(v interface{}) bool
	Iterator(fn func(interface{}) bool)
	ReverseIterator(fn func(interface{}) bool)
}
