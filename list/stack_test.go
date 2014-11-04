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

	stack.Add(100)
	if stack.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}

	stack.Remove()
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
		stack.Add(i + 1)
	}

	for i := 100; i > 0; i-- {
		if stack.Size() != i {
			t.Fatal("Size 方法测试失败")
		}
		stack.Remove()
	}
}

func TestPushPop(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Add(i + 1)
		slice[100-1-i] = i + 1
	}

	for _, i := range slice {
		if i != stack.Remove().(int) {
			t.Fatal("Push Pop 方法测试失败")
		}
	}
}

func TestIterator(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Add(i + 1)
		slice[100-1-i] = i + 1
	}

	i := 0
	stack.Travel(func(v interface{}) bool {
		if v.(int) != slice[i] {
			t.Fatal("Iterator 方法测试失败")
		}
		i++
		return true
	})
}

func TestReverseInterator(t *testing.T) {
	stack := NewStack()
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		stack.Add(i + 1)
		slice[i] = i + 1
	}

	i := 0
	stack.ReverseTravel(func(v interface{}) bool {
		if v.(int) != slice[i] {
			t.Fatal("ReverseIterator 方法测试失败")
		}
		i++
		return true
	})
}
