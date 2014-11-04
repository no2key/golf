package list

type SLList interface {
	IsEmpty() bool
	Clear()
	Size() int
	Add(v interface{})
	Remove() interface{}
	Peek() interface{}
	Contains(v interface{}) bool
	Travel(fn func(interface{}) bool)
	ReverseTravel(fn func(interface{}) bool)
}
