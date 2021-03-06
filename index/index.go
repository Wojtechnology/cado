package index

type Key interface {
	Hashable
	Less(than Key) bool
}

type Value interface {
	Hashable
	Merge(with Value) Value
}

type Index interface {
	Put(Key, Value)
	Get(Key) Value
	RootHash() []byte
}

func keysEqual(k1 Key, k2 Key) bool {
	return !k1.Less(k2) && !k2.Less(k1)
}
