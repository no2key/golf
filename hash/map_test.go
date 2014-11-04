// map_test
package hash

import (
	"fmt"
	"testing"
)

func TestIntMap(t *testing.T) {
	m1 := make(map[int]string, 100)
	m2 := NewSizedHashMap(100)

	for i := 0; i < 100; i++ {
		m2.Put(Int(i), fmt.Sprintf("%d", i*10))

		if m2.Size() != i+1 {
			t.Fatal("Map Size 函数测试失败")
		}
	}

	m2.Clear()
	if !m2.IsEmpty() {
		t.Fatal("Map IsEmpty Clear  函数失败")
	}

	for i := 0; i < 100; i++ {
		m1[i] = fmt.Sprintf("%d", i*10)
		m2.Put(Int(i), fmt.Sprintf("%d", i*10))

		if m2.Size() != i+1 {
			t.Fatal("Map Size 函数测试失败")
		}
	}

	for i := 0; i < 100; i++ {
		if m1[i] != m2.Get(Int(i)).(string) {
			t.Fatal("Map Get 函数失败")
		}
	}

	for i := 0; i < 100; i++ {
		if !m2.Contains(Int(i)) {
			t.Fatal("Map Contains 函数失败")
		}
	}

	for i := 0; i < 100; i++ {
		if m2.Remove(Int(i)).(string) != fmt.Sprintf("%d", i*10) {
			t.Fatal("Map Remove 函数失败")
		}

		if m2.Size() != 100-i-1 {
			t.Fatal("Map Remove 函数失败")
		}
	}

	if !m2.IsEmpty() {
		t.Fatal("Map IsEmpty 函数失败")
	}
}

func TestStringMap(t *testing.T) {
	m1 := make(map[string]string, 100)
	m2 := NewSizedHashMap(100)

	for i := 0; i < 100; i++ {
		m2.Put(String(fmt.Sprintf("%d", i)), fmt.Sprintf("%d", i*10))

		if m2.Size() != i+1 {
			t.Fatal("Map Size 函数测试失败")
		}
	}

	m2.Clear()
	if !m2.IsEmpty() {
		t.Fatal("Map IsEmpty Clear  函数失败")
	}

	for i := 0; i < 100; i++ {
		m1[fmt.Sprintf("%d", i)] = fmt.Sprintf("%d", i*10)
		m2.Put(String(fmt.Sprintf("%d", i)), fmt.Sprintf("%d", i*10))

		if m2.Size() != i+1 {
			t.Fatal("Map Size 函数测试失败")
		}
	}

	for i := 0; i < 100; i++ {
		if m1[fmt.Sprintf("%d", i)] != m2.Get(String(fmt.Sprintf("%d", i))).(string) {
			t.Fatal("Map Get 函数失败")
		}
	}

	for i := 0; i < 100; i++ {
		if !m2.Contains(String(fmt.Sprintf("%d", i))) {
			t.Fatal("Map Contains 函数失败")
		}
	}

	for i := 0; i < 100; i++ {
		if m2.Remove(String(fmt.Sprintf("%d", i))).(string) != fmt.Sprintf("%d", i*10) {
			t.Fatal("Map Remove 函数失败")
		}

		if m2.Size() != 100-i-1 {
			t.Fatal("Map Remove 函数失败")
		}
	}

	if !m2.IsEmpty() {
		t.Fatal("Map IsEmpty 函数失败")
	}
}
