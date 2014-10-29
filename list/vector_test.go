// vector_test.go
package list

import (
	"testing"
)

func TestNewVector(t *testing.T) {
	if v := NewVector(); v == nil {
		t.Fatal("NewVector 函数测试失败")
	}

	if v := NewSizedVector(100); v == nil {
		t.Fatal("NewVector 函数测试失败")
	}
}

func TestAddGet(t *testing.T) {
	v := NewSizedVector(100)
	s := make([]int, 100)

	for i := 0; i < 100; i++ {
		v.Add(i + 1)
		s[i] = i + 1
	}

	for i := 0; i < 100; i++ {
		if v.Get(i).(int) != s[i] {
			t.Fatal("Add Get 方法测试失败")
		}
	}
}

func TestAddTo(t *testing.T) {
	v := NewSizedVector(100)
	s := make([]int, 100)

	for i := 0; i < 100; i++ {
		if i == 49 {
			continue
		}
		v.Add(i + 1)
		s[i] = i + 1
	}

	v.AddTo(49, 50)
	s[49] = 50

	for i := 0; i < 100; i++ {
		if v.Get(i).(int) != s[i] {
			t.Fatal("AddTo Get 方法测试失败")
		}
	}
}

func TestIsEmptyClear(t *testing.T) {
	v := NewSizedVector(100)
	if !v.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}
	v.Add(100)
	if v.IsEmpty() {
		t.Fatal("IsEmpty 方法测试失败")
	}
	v.Clear()
	if !v.IsEmpty() {
		t.Fatal("Clear IsEmpty 方法测试失败")
	}
}
