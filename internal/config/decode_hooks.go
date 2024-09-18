package config

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func composeDecodeHook() mapstructure.DecodeHookFunc {
	return mapstructure.ComposeDecodeHookFunc(
		decodeStringToURLHook(),
	)
}

func decodeStringToURLHook() mapstructure.DecodeHookFuncType {
	return func(f reflect.Type, t reflect.Type, data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(url.URL{}) {
			return data, nil
		}

		input := data.(string)

		if input == "" {
			return nil, nil
		}

		output, err := url.Parse(input)
		if err != nil {
			return nil, fmt.Errorf("error parsing %s: %w", input, err)
		}

		return output, nil
	}
}
