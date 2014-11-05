// set_impl
package hash

type hashSet struct {
	internalMap Map
}

func newSizedHashSet(capacity int) *hashSet {
	set := new(hashSet)
	set.internalMap = NewSizedHashMap(capacity)
	return set
}

func NewSizedHashSet(capacity int) Set {
	return newSizedHashSet(capacity)
}

func NewHashSet() Set {
	return newSizedHashSet(16)
}

func (this *hashSet) Size() int {
	return this.internalMap.Size()
}

func (this *hashSet) IsEmpty() bool {
	return this.internalMap.IsEmpty()
}

func (this *hashSet) Add(h Hash) bool {
	return this.internalMap.Put(h, this) == nil
}

func (this *hashSet) Remove(h Hash) bool {
	return this.internalMap.Remove(h) != nil
}

func (this *hashSet) Contains(h Hash) bool {
	return this.Contains(h)
}

func (this *hashSet) Clear() {
	this.internalMap.Clear()
}

func (this *hashSet) Travel(fn func(Hash) bool) {
	this.internalMap.Travel(func(e MapEntry) bool {
		return fn(e.Key())
	})
}
