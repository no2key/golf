// set_impl
package hash

type hashSetEntry struct {
	key  Hash
	hash uint32
	next *hashMapEntry
}

func newHashSetEntry(key Hash, hash uint32, next *hashMapEntry) *hashMapEntry {
	rs := new(hashMapEntry)
	rs.key = key
	rs.hash = hash
	rs.next = next
	return rs
}

type hashSet struct {
	table     []*hashSetEntry
	size      int
	threshold int
}

func (this *hashSet) Size() int {
	return this.size
}

func (this *hashSet) IsEmpty() bool {
	return this.size == 0
}

func (this *hashSet) Add(h Hash) bool {
	return false
}

func (this *hashSet) Remove(h Hash) bool {
	return false
}

func (this *hashSet) Contains(h Hash) bool {
	return false
}

func (this *hashSet) Clear() {
}

func (this *hashSet) Travel(fn func(Hash) bool) {
}
