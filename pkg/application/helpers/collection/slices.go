package collection

import (
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"
)

func SortSlice[K constraints.Ordered](slice []K) []K {
    sorted := make([]K, len(slice))
    copy(sorted, slice)
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i] < sorted[j]
    })
    return sorted
}

func SliceIsEqual[T comparable](aa, bb []T) bool {
    eqCtr := 0
    for _, a := range aa {
        for _, b := range bb {
            if reflect.DeepEqual(a, b) {
                eqCtr++
            }
        }
    }
    if eqCtr != len(bb) || len(aa) != len(bb) {
        return false
    }
    return true
}