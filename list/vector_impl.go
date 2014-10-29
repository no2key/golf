package list

type vector struct {
	base []interface{}
	pt   int
}

func newVector() *vector {
	vector := new(vector)
	vector.base = make([]interface{}, 16)
	vector.pt = 0
	return vector
}

func (this *vector) Add(v interface{}) {
	size := len(this.base)
	if this.pt == size {
		newBase := make([]interface{}, size*2)
		copy(newBase, this.base)
		this.base = newBase
	}
	this.base[this.pt] = v
	this.pt++
}

func (this *vector) AddTo(i int, v interface{}) error {
	size := this.pt
	if i < 0 || i > size {
		return IndexOutOfBounds
	}

	if i == size {
		this.Add(v)
		return nil
	}

	// 0 <= i < size
	size = len(this.base)
	if this.pt == size {
		newBase := make([]interface{}, size*2)
		copy(newBase, this.base[:i]) // copy [0..i]
		newBase[i] = v
		copy(newBase[i+1:], this.base[i:]) // copy [i..]
		this.base = newBase
	} else {
		copy(this.base[i+1:], this.base[i:])
		this.base[i] = v
	}
	this.pt++

	return nil
}

func (this *vector) Set(i int, v interface{}) interface{} {
	size := this.pt
	if i < 0 || i >= size {
		return IndexOutOfBounds
	}

	old := this.base[i]
	this.base[i] = v
	return old
}

func (this *vector) Get(i int) interface{} {
	size := this.pt
	if i < 0 || i >= size {
		return IndexOutOfBounds
	}

	return this.base[i]
}

func (this *vector) Contains(v interface{}) bool {
	for i, ele := range this.base {
		if i == this.pt {
			break
		}
		if ele == v {
			return true
		}
	}

	return false
}

func (this *vector) Remove(i int) interface{} {
	size := this.pt
	if i < 0 || i >= size {
		return IndexOutOfBounds
	}

	old := this.base[i]
	if i != size-1 {
		copy(this.base[i:], this.base[i+1:size])
	}
	this.pt--
	this.base[this.pt] = nil // 避免内存泄露
	return old
}

func (this *vector) RemoveObject(v interface{}) bool {
	index := -1
	for i, ele := range this.base {
		if i == this.pt {
			break
		}
		if ele == v {
			index = i
			break
		}
	}

	if index != -1 {
		this.Remove(index)
		return true
	} else {
		return false
	}
}

func (this *vector) Clear() {
	for i, _ := range this.base {
		this.base[i] = nil
	}
	this.pt = 0
}

func (this *vector) Size() int {
	return this.pt
}

func (this *vector) IsEmpty() bool {
	return this.pt == 0
}

func (this *vector) IndexOf(v interface{}) int {
	index := -1
	for i, ele := range this.base {
		if i == this.pt {
			break
		}
		if ele == v {
			index = i
			break
		}
	}
	return index
}

func (this *vector) LastIndexOf(v interface{}) int {
	index := -1
	for i := this.pt - 1; i >= 0; i-- {
		if this.base[i] == v {
			index = i
			break
		}
	}
	return index
}

func (this *vector) Sub(start, end int) (List, error) {
	if start < 0 || start > end || end > this.pt {
		return nil, IndexOutOfBounds
	}
	rs := newVector()
	rs.pt = end - start
	if rs.pt != 0 {
		size := len(rs.base)
		for rs.pt > size {
			size *= 2
		}
		rs.base = make([]interface{}, size)
		copy(rs.base, this.base[start:end])
	}
	return rs, nil
}

func (this *vector) ToSlice(s []interface{}) []interface{} {
	if this.pt == 0 {
		return nil
	}

	rs := make([]interface{}, this.pt)
	for i := 0; i < this.pt; i++ {
		rs[i] = this.base[i]
	}
	return rs
}
