// queue_impl
package linear

func NewQueue() SLList {
	return newQueue()
}

type queueNode struct {
	value interface{}
	next  *queueNode
}

type queue struct {
	head *queueNode
	tail *queueNode
	size int
}

func newQueue() *queue {
	rs := new(queue)
	return rs
}

func (this *queue) IsEmpty() bool {
	return this.size == 0
}

func (this *queue) Size() int {
	return this.size
}

func (this *queue) Clear() {
	if this.size != 0 {
		this.head = nil
		this.tail = nil
		this.size = 0
	}
}

func (this *queue) Add(v interface{}) {
	if this.size == 0 {
		node := new(queueNode)
		node.value = v
		this.head = node
		this.tail = node
	} else {
		this.tail.next = new(queueNode)
		this.tail.next.value = v
		this.tail = this.tail.next
	}
	this.size++
}

func (this *queue) Remove() interface{} {
	if this.size == 0 {
		return nil
	} else if this.size == 1 {
		old := this.head.value
		this.Clear()
		return old
	} else {
		old := this.head.value
		this.head = this.head.next
		this.size--
		return old
	}
}

func (this *queue) Peek() interface{} {
	if this.size == 0 {
		return nil
	} else {
		return this.head.value
	}
}

func (this *queue) Contains(v interface{}) bool {
	rs := false
	for node := this.head; node != nil; node = node.next {
		if node.value == v {
			rs = true
			break
		}
	}
	return rs
}

func (this *queue) Iterator(fn func(interface{}) bool) {
	if this.size == 0 {
		return
	}

	for node := this.head; node != nil; node = node.next {
		if !fn(node.value) {
			break
		}
	}
}

func (this *queue) ReverseIterator(fn func(interface{}) bool) {
	if this.size == 0 {
		return
	}

	s := make([]interface{}, this.size)
	for node, i := this.head, this.size-1; node != nil; node, i = node.next, i+1 {
		s[i] = node.value
	}

	for _, v := range s {
		if !fn(v) {
			break
		}
	}
}
