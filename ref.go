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
