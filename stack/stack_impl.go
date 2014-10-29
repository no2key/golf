package stack

type node struct {
	prev *node
	elem interface{}
}

type stack struct {
	pt   *node
	size int
}

func (this *stack) IsEmpty() bool {
	return this.size == 0
}

func (this *stack) Size() int {
	return this.size
}

func (this *stack) Push(v interface{}) {
	if this.pt == nil {
		this.pt = new(node)
		this.pt.elem = v
	} else {
		newNode := new(node)
		newNode.prev = this.pt
		newNode.elem = v
		this.pt = newNode
	}

	this.size++
}

func (this *stack) Pop() interface{} {
	if this.size == 0 {
		return nil
	}

	rs := this.pt
	this.pt = this.pt.prev
	this.size--
	return rs.elem
}

func (this *stack) Iterator(fn func(interface{})) {
	if this.size > 0 {
		for node := this.pt; node != nil; node = node.prev {
			fn(node.elem)
		}
	}
}

func (this *stack) ReverseIterator(fn func(interface{})) {
	if this.size > 0 {
		nodeSlice := make([]interface{}, this.size)
		for node, i := this.pt, this.size-1; node != nil && i >= 0; node, i = node.prev, i-1 {
			nodeSlice[i] = node.elem
		}

		for _, elem := range nodeSlice {
			fn(elem)
		}
	}
}
