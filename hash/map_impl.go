// hash_impl
package hash

import (
	"bytes"
	"fmt"
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
	// key = nil 时的值
	nilVal interface{}
}

func (this *hashMap) Put(h Hash, v interface{}) interface{} {
	if h == nil {
		prev := this.nilVal
		this.nilVal = v
		return prev
	}
	hash := secondaryHash(h)
	index := hash & (uint32(len(this.table)) - 1)
	// 查找是否已存在值
	for entry := this.table[index]; entry != nil; entry = entry.next {
		if entry.hash == hash && h.Equals(entry.key) {
			rs := entry.value
			entry.value = v
			return rs
		}
	}
	// 新值
	if this.size > this.threshold {
		tab := this.doubleCapacity()
		index = hash & (uint32(len(tab) - 1))
	}
	this.size++
	this.addNewEntry(h, v, hash, index)
	return nil
}

func (this *hashMap) Get(h Hash) interface{} {
	if h == nil {
		return this.nilVal
	}
	hash := secondaryHash(h)
	tab := this.table
	for e := this.table[int(hash)&(len(tab)-1)]; e != nil; e = e.next {
		key := e.key
		if key == h || (e.hash == hash && h.Equals(key)) {
			return e.value
		}
	}
	return nil
}

func (this *hashMap) Remove(h Hash) interface{} {
	if h == nil {
		oldV := this.nilVal
		this.nilVal = nil
		return oldV
	}
	hash := secondaryHash(h)
	tab := this.table
	index := int(hash) & (len(tab) - 1)
	var prev, e *hashMapEntry
	for e, prev = tab[index], nil; e != nil; prev, e = e, e.next {
		if e.hash == hash || h.Equals(e.key) {
			if prev == nil {
				tab[index] = e.next
			} else {
				prev.next = e.next
			}
			this.size--
			return e.value
		}
	}
	return nil
}

func (this *hashMap) Contains(h Hash) bool {
	if h == nil {
		return false
	}
	hash := secondaryHash(h)
	tab := this.table
	for e := tab[int(hash)&(len(tab)-1)]; e != nil; e = e.next {
		key := e.key
		if key == h || (e.hash == hash && h.Equals(e.key)) {
			return true
		}
	}
	return false
}

func (this *hashMap) Size() int {
	return this.size
}

func (this *hashMap) IsEmpty() bool {
	return this.size == 0
}

func (this *hashMap) Clear() {
	if this.size != 0 {
		for i, _ := range this.table {
			this.table[i] = nil
			this.size = 0
		}
	}
}

func (this *hashMap) Travel(fn func(MapEntry) bool) {
	if this.size != 0 {
		tab := this.table
		var rs bool

		for _, entry := range tab {
			if entry == nil {
				continue
			}

			for n := entry; n != nil; n = n.next {
				rs = fn(n)
			}

			if !rs {
				break
			}
		}
	}
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

func roundUpToPowerOfTwo(i int) int {
	i-- // If input is a power of two, shift its high-order bit right.

	// "Smear" the high-order bit all the way to the right.
	i |= i >> 1
	i |= i >> 2
	i |= i >> 4
	i |= i >> 8
	i |= i >> 16

	return i + 1
}

func (this *hashMap) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("map [")
	index := -1
	this.Travel(func(e MapEntry) bool {
		index++
		if index != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%v", e))
		return true
	})
	buf.WriteString("]")
	return buf.String()
}

func newSizedHashMap(capacity int) *hashMap {
	newMap := new(hashMap)
	if capacity < _MINIMUM_CAPACITY {
		capacity = _MINIMUM_CAPACITY
	} else if capacity > _MAXIMUM_CAPACITY {
		capacity = _MAXIMUM_CAPACITY
	} else {
		capacity = roundUpToPowerOfTwo(capacity)
	}
	newMap.makeTable(capacity)
	return newMap
}

func newHashMap() *hashMap {
	return newSizedHashMap(16)
}

func NewSizedHashMap(capacity int) Map {
	return newSizedHashMap(capacity)
}

func NewHashMap() Map {
	return newHashMap()
}
