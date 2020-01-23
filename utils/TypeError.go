package utils

import (
	"fmt"
	"reflect"
)

func TypeError(expected, got interface{}) error {
	return fmt.Errorf("expected %s but got %s", reflect.TypeOf(expected).String(), reflect.TypeOf(got).String())
}
