package masked

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func MapStructureDecodeHook() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(MaskedText{}) {
			return data, nil
		}

		if n, ok := t.(masker); ok {
			return NewMaskedText(data.(string), n.DefaultMaskFunc()), nil
		}

		return NewMaskedText(data.(string), FixedMaskFunc), nil
	}
}
