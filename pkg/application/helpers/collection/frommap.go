package collection

func KeysFromMap[K comparable, V any](m map[K]V) []K {
	var keys []K
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func ValuesFromMap[K comparable, V any](m map[K]V) []V {
	var values []V
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

func MapToSlice[K comparable, V any](m map[K]V) []KeyValuePair[K, V] {
	var slice []KeyValuePair[K, V]
	for key, value := range m {
		slice = append(slice, KeyValuePair[K, V]{Key: key, Value: value})
	}
	return slice
}
