package hexago

import (
	"fmt"
	"reflect"
)

func extractRepoField(v reflect.Value) (*reflect.Value, error) {
	defer func() {
		recover()
	}()

	var value reflect.Value

	for j := 0; j < v.NumField(); j++ {
		value = v.Field(j)
		_, ok := value.Interface().(Repo)
		if ok {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("repo field not found")
}

func hasRepoField(v reflect.Value) bool {
	defer func() {
		recover()
	}()

	for j := 0; j < v.NumField(); j++ {
		_, ok := v.Field(j).Interface().(Repo)
		if ok {
			return true
		}
	}

	return false
}
