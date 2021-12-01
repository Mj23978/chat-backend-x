package viper

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/spf13/viper"

	"github.com/mj23978/chat-backend-x/utils"
)

// UnmarshalKey unmarshals a viper key into the destination struct. The destination struct
// must be JSON-compatible (i.e. have `json` struct tags)
func UnmarshalKey(key string, destination interface{}) error {
	value := viper.Get(key)
	if value == `null` || value == "" || value == nil {
		value = make(map[string]interface{})
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(utils.ToJSONMap(viper.Get(key))); err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(json.NewDecoder(&b).Decode(destination))
}
