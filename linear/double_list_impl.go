// double_list_impl
package linear

type dLListNode struct {
	next  *dLListNode
	prev  *dLListNode
	value interface{}
}

type dLList struct {
	head *dLListNode
	tail *dLListNode
	size int
}

func newDLList() *dLList {
	return new(dLList)
}

func (this *dLList) AddTail(v interface{}) {
	if this.size == 0 {
		newNode := new(dLListNode)
		newNode.value = v
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.next = new(dLListNode)
		this.tail.next.prev = this.tail
		this.tail.next.value = v
		this.tail = this.tail.next
	}
	this.size++
}

func (this *dLList) AddHead(v interface{}) {
	if this.size == 0 {
		newNode := new(dLListNode)
		newNode.value = v
		this.head = newNode
		this.tail = newNode
	} else {
		this.head.prev = new(dLListNode)
		this.head.prev.next = this.head
		this.head.prev.value = v
		this.head = this.head.prev
	}
	this.size++
}

func (this *dLList) PeekTail() interface{} {
	if this.size == 0 {
		return nil
	}
	return this.tail.value
}

func (this *dLList) PeekHead() interface{} {
	if this.size == 0 {
		return nil
	}
	return this.head.value
}

func (this *dLList) RemoveTail() interface{} {
	if this.size == 0 {
		return nil
	} else if this.size == 1 {
		old := this.tail.value
		this.Clear()
		return old
	} else {
		old := this.tail.value
		this.tail.prev.next = nil
		this.tail.value = nil
		this.tail = this.tail.prev
		this.size--
		return old
	}
}

func (this *dLList) RemoveHead() interface{} {
	if this.size == 0 {
		return nil
	} else if this.size == 1 {
		old := this.head.value
		this.Clear()
		return old
	} else {
		old := this.head.value
		this.head.next.prev = nil
		this.head.value = nil
		this.head = this.head.next
		this.size--
		return old
	}
}

func (this *dLList) Clear() {
	this.head = nil
	this.tail = nil
	this.size = 0
}

func (this *dLList) Size() int {
	return this.size
}

func (this *dLList) IsEmpty() bool {
	return this.size == 0
}

func (this *dLList) Contains(v interface{}) bool {
	rs := false
	for node := this.head; node != nil; node = node.next {
		if node.value == v {
			rs = true
			break
		}
	}
	return rs
}

func (this *dLList) Iterator(fn func(interface{}) bool) {
	for node := this.head; node != nil; node = node.next {
		if !fn(node.value) {
			break
		}
	}
}

func (this *dLList) ReverseIterator(fn func(interface{}) bool) {
	for node := this.tail; node != nil; node = node.prev {
		if !fn(node.value) {
			break
		}
	}
}
