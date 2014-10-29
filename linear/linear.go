package linear

type Linear interface {
	IsEmpty() bool
	Clear()
	Size() int
	Add(v interface{})
	Remove() interface{}
	Peek() interface{}
	Iterator(fn func(interface{}))
	ReverseIterator(fn func(interface{}))
}
