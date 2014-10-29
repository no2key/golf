package list

type List interface {
	Add(v interface{})
	AddTo(i int, v interface{}) error

	Set(i int, v interface{}) interface{}

	Get(i int) interface{}
	Contains(v interface{}) bool

	Remove(i int) interface{}
	RemoveObject(v interface{}) bool

	Clear()
	Size() int
	IsEmpty() bool

	IndexOf(v interface{}) int
	LastIndexOf(v interface{}) int

	Sub(start, end int) (List, error)

	ToSlice(s []interface{}) []interface{}
}
