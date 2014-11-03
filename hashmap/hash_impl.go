// hash_impl
package hashmap

import (
	"errors"
	"fmt"
)

var (
	paramIsNilError = errors.New("the param is nil")
)

const (
	_MINIMUM_CAPACITY = 4
	_MAXIMUM_CAPACITY = 1 << 30
)

type hashMapEntry struct {
	key   Hash
	value interface{}
	hash  uint32
	next  *hashMapEntry
}

func newHashMapEntry(key Hash, value interface{},
	hash uint32, next *hashMapEntry) *hashMapEntry {
	rs := new(hashMapEntry)
	rs.key = key
	rs.value = value
	rs.hash = hash
	rs.next = next
	return rs
}

func (this *hashMapEntry) Key() Hash {
	return this.key
}

func (this *hashMapEntry) Value() interface{} {
	return this.value
}

func (this *hashMapEntry) String() string {
	return fmt.Sprintf("%v: %v", this.key, this.value)
}

type hashMap struct {
	table     []*hashMapEntry
	size      int
	threshold int
}

func (this *hashMap) Put(h Hash, v interface{}) (interface{}, error) {
	if h == nil {
		return nil, paramIsNilError
	}
	hash := secondaryHash(h)
	index := hash & (uint32(len(this.table)) - 1)
	// 查找是否已存在值
	for entry := this.table[index]; entry != nil; entry = entry.next {
		if entry.hash == hash && h.Equals(entry.key) {
			rs := entry.value
			entry.value = v
			return rs, nil
		}
	}
	// 新值
	if this.size > this.threshold {
		this.size++
		tab := this.doubleCapacity()
		index = hash & (uint32(len(tab) - 1))
	}
	this.addNewEntry(h, v, hash, index)
	return nil, nil
}

func (this *hashMap) Get(h Hash) (interface{}, error) {
	if h == nil {
		return nil, paramIsNilError
	}
	hash := secondaryHash(h)
	tab := this.table
	for e := this.table[int(hash)&(len(tab)-1)]; e != nil; e = e.next {
		key := e.key
		if key == h || (e.hash == hash && h.Equals(key)) {
			return e.value, nil
		}
	}
	return nil, nil
}

func (this *hashMap) Remove(h Hash) interface{} {
	return nil
}

func (this *hashMap) Contains(h Hash) bool {
	return false
}

func (this *hashMap) Size() int {
	return this.size
}

func (this *hashMap) IsEmpty() bool {
	return this.size == 0
}

func (this *hashMap) Clear() {
}

func (this *hashMap) Travel(func(MapEntry) bool) {
}

func (this *hashMap) addNewEntry(h Hash, v interface{}, hash uint32, index uint32) {
	this.table[index] = newHashMapEntry(h, v, hash, this.table[index])
}

func (this *hashMap) doubleCapacity() []*hashMapEntry {
	oldTable := this.table
	oldCapacity := len(oldTable)
	if oldCapacity == _MAXIMUM_CAPACITY {
		return oldTable
	}
	newCapacity := oldCapacity * 2
	newTable := this.makeTable(newCapacity)
	if this.size == 0 {
		return newTable
	}

	// 拷贝并重新计算就 table 中的元素
	for j, e := range oldTable {
		if e == nil {
			continue
		}
		highBit := int(e.hash) & oldCapacity
		newTable[j|highBit] = e
		var broken *hashMapEntry
		for n := e.next; n != nil; n = n.next {
			nextHighBit := int(n.hash) & oldCapacity
			if nextHighBit != highBit {
				if broken == nil {
					newTable[j|nextHighBit] = n
				} else {
					broken.next = n
				}
				broken = e
				highBit = nextHighBit
			}
		}
		if broken != nil {
			broken.next = nil
		}
	}
	return newTable
}

func (this *hashMap) makeTable(newCapacity int) []*hashMapEntry {
	newTable := make([]*hashMapEntry, newCapacity)
	this.table = newTable
	this.threshold = (newCapacity >> 1) + (newCapacity >> 2) // 3/4 capacity
	return newTable
}

func secondaryHash(h Hash) uint32 {
	rs := h.HashCode()
	rs ^= (rs >> 20) ^ (rs >> 12)
	rs ^= (rs >> 7) ^ (rs >> 4)
	return rs
}

func NewMap() Map {
	return new(hashMap)
}
