package generics

import (
	"fmt"
	"reflect"
)

func CastSlice[IN, OUT any](in []IN) ([]OUT, error) {
	out := make([]OUT, len(in))
	outType := reflect.TypeOf(out).Elem()
	inType := reflect.TypeOf(in).Elem()
	if !inType.ConvertibleTo(outType) {
		return nil, fmt.Errorf("cannot convert %s to %s", inType.String(), outType.String())
	}

	outVal := reflect.ValueOf(out)

	for i, v := range in {
		outVal.Index(i).Set(reflect.ValueOf(v).Convert(outType))
	}

	return out, nil
}
