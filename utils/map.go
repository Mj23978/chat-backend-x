package utils

import "github.com/mitchellh/mapstructure"

func AppendMapIfExists(condition bool, des map[string]interface{}, key string, value interface{}) {
	if condition {
		des[key] = value
	}
}

func MapToStruct(to interface{}, tag string, from interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: tag, Result: to, WeaklyTypedInput: true})
	if err != nil {
		return err
	}
	if err := decoder.Decode(from); err != nil {
		return err
	}
	return nil
}