package unmarshal

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func Unmarshal(s string, v any) error {
	vType := reflect.TypeOf(v)
	if vType == nil || vType.Kind() != reflect.Pointer {
		return errors.New("pointer is needed")
	}

	output, err := parse(s)
	if err != nil {
		return err
	}

	vVal := reflect.ValueOf(v)
	oType := reflect.TypeOf(output)
	oVal := reflect.ValueOf(output)

	switch vType.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Map:
		if vType.Elem().Kind() != oType.Kind() {
			return errors.New("dismatched kind")
		}
		vVal.Elem().Set(oVal)
	case reflect.Struct:
		if oType.Kind() != reflect.Map {
			return nil
		}

		d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "gson", Result: v})
		if err != nil {
			panic(err)
		}

		d.Decode(output)
		// 这里可以decode("", output, vVal)
	default:
		return errors.New("dismatched kind")
	}

	return nil
}
