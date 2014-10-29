package stack

type Stack interface {
	IsEmpty() bool
	Size() int
	Push(v interface{})
	Pop() interface{}
	Iterator(fn func(interface{}))
	ReverseIterator(fn func(interface{}))
}

func NewStack() Stack {
	return new(stack)
}
