package helper

import "reflect"

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func Contains(s []reflect.Type, e reflect.Type) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
