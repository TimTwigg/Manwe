package lists

import (
	"reflect"
)

func Reduce[T any](vals []T, reduction func(T, T) T) T {
	ret := vals[0]
	for i := 1; i < len(vals); i++ {
		ret = reduction(ret, vals[i])
	}
	return ret
}

func Any(arr []bool) bool {
	for _, item := range arr {
		if item {
			return true
		}
	}
	return false
}

func All(arr []bool) bool {
	for _, item := range arr {
		if !item {
			return false
		}
	}
	return true
}

func MapWithError[T any, U any](vals []T, mapfunc func(T) (U, error)) ([]U, error) {
	mappedVals := make([]U, len(vals))
	for i, val := range vals {
		mapped, err := mapfunc(val)
		if err != nil {
			return nil, err
		}
		mappedVals[i] = mapped
	}
	return mappedVals, nil
}

func UnpackArray(s any) []any {
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}
