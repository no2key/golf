package linear

import (
	"bytes"
	"fmt"
)

func NewStack() SLList {
	return new(stack)
}

type stackNode struct {
	prev *stackNode
	elem interface{}
}

type stack struct {
	pt   *stackNode
	size int
}

func (this *stack) IsEmpty() bool {
	return this.size == 0
}

func (this *stack) Size() int {
	return this.size
}

func (this *stack) Clear() {
	this.pt = nil
	this.size = 0
}

func (this *stack) Add(v interface{}) {
	if this.pt == nil {
		this.pt = new(stackNode)
		this.pt.elem = v
	} else {
		newNode := new(stackNode)
		newNode.prev = this.pt
		newNode.elem = v
		this.pt = newNode
	}

	this.size++
}

func (this *stack) Remove() interface{} {
	if this.size == 0 {
		return nil
	}

	rs := this.pt
	this.pt = this.pt.prev
	this.size--
	return rs.elem
}

func (this *stack) Peek() interface{} {
	if this.size == 0 {
		return nil
	}

	return this.pt.elem
}

func (this *stack) Iterator(fn func(interface{}) bool) {
	if this.size > 0 {
		for node := this.pt; node != nil; node = node.prev {
			if !fn(node.elem) {
				break
			}
		}
	}
}

func (this *stack) ReverseIterator(fn func(interface{}) bool) {
	if this.size > 0 {
		nodeSlice := make([]interface{}, this.size)
		for node, i := this.pt, this.size-1; node != nil && i >= 0; node, i = node.prev, i-1 {
			nodeSlice[i] = node.elem
		}

		for _, v := range nodeSlice {
			if !fn(v) {
				break
			}
		}
	}
}

func (this *stack) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("stack [")
	for node, i := this.pt, 0; node != nil; node, i = node.prev, i+1 {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%v", node.elem))
	}
	buf.WriteString("]")
	return buf.String()
}
