package data_type_utils

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

type MapLike[K comparable, V any] interface {
	Get(K) V
	Set(K, V)
	Has(K) bool
	Iter() []KeyValuePair[K, V]
	Delete(K)
	Keys() []K
	Values() []V
}
