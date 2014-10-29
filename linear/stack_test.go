package list

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	if stack := NewStack(); stack == nil {
		t.Fatal("NewStack 函数测试失败")
	}
}

func TestEmpty(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}

	stack.Push(100)
	if stack.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}
}

func TestSize(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 100; i++ {
		if stack.Size() != i {
			t.Fatal("Size 方法测试失败")
		}
		stack.Push(i + 1)
	}

	for i := 100; i > 0; i-- {
		if stack.Size() != i {
			t.Fatal("Size 方法测试失败")
		}
		stack.Pop()
	}
}

func TestPushPop(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Push(i + 1)
		slice[100-1-i] = i + 1
	}

	for _, i := range slice {
		if i != stack.Pop().(int) {
			t.Fatal("Push Pop 方法测试失败")
		}
	}
}

func TestIterator(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Push(i + 1)
		slice[100-1-i] = i + 1
	}

	i := 0
	stack.Iterator(func(v interface{}) {
		if v.(int) != slice[i] {
			t.Fatal("Iterator 方法测试失败")
		}
		i++
	})
}

func TestReverseInterator(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Push(i + 1)
		slice[i] = i + 1
	}

	i := 0
	stack.ReverseIterator(func(v interface{}) {
		if v.(int) != slice[i] {
			t.Fatal("ReverseIterator 方法测试失败")
		}
		i++
	})
}
