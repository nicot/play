package main

import (
	"errors"
	"testing"
)

type GetFilter func(key interface{}, next GetFilter) (interface{}, error)
type SetFilter func(key interface{}, value interface{}, next SetFilter) error

var kvs map[interface{}] interface{}

func TestKV(t *testing.T) {
	kvs = make(map[interface{}]interface{})
	err := PreventEvensSetter(2, 2, setUnderlying)
	t.Error(err)
	t.Error(getUnderlying(2, nil))
}

func getUnderlying(key interface{}, _ GetFilter) (interface{}, error) {
	return kvs[key], nil
}

func setUnderlying(key interface{}, value interface{}, _ SetFilter) error {
	kvs[key] = value
	return nil
}

func PreventEvensSetter(key interface{}, value interface{}, next SetFilter) error {
	if n, isInt := key.(int); isInt {
		if n % 2 == 0 {
			return errors.New("NO EVEN NUMBERS")
		}
	}
	return next(key, value, nil)
}
