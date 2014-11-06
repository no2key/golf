// weak_ref
package weak

import (
	"errors"
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

var notPointerError = errors.New("Not a pointer")

type Ref struct {
	t uintptr // Type
	d uintptr // Data
}

func New(v interface{}) *Ref {
	// not include reflect.UnsafePointer Kind
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		panic(notPointerError)
	}
	i := (*[2]uintptr)(unsafe.Pointer(&v))
	w := &Ref{^i[0], ^i[1]}
	runtime.SetFinalizer((*uintptr)(unsafe.Pointer(i[1])), func(_ *uintptr) {
		atomic.StoreUintptr(&w.d, 0)
		w.t = 0
	})
	return w
}

func (w *Ref) Get() (v interface{}) {
	t := w.t
	if d := atomic.LoadUintptr(&w.d); d != 0 {
		i := (*[2]uintptr)(unsafe.Pointer(&v))
		i[0] = ^t
		i[1] = ^d
	}
	return
}
