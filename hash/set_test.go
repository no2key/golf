// set_test
package hash

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewHashSet()

	if !set.IsEmpty() {
		t.Fatal("Set IsEmpty 方法测试失败")
	}

	if set.Size() != 0 {
		t.Fatal("Set Size 方法测试失败")
	}

	if !set.Add(Int(100)) {
		t.Fatal("Set Add 方法测试失败")
	}

	if set.IsEmpty() {
		t.Fatal("Set IsEmpty 方法测试失败")
	}

	if set.Size() != 1 {
		t.Fatal("Set Size 方法测试失败")
	}

	if set.Add(Int(100)) {
		t.Fatal("Set Add 方法测试失败")
	}

	if set.IsEmpty() {
		t.Fatal("Set IsEmpty 方法测试失败")
	}

	if set.Size() != 1 {
		t.Fatal("Set Size 方法测试失败")
	}
}
