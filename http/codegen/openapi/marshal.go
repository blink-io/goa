package openapi

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func init() {
	// FIXME See the changes in https://github.com/go-yaml/yaml/releases/tag/v2.4.0
	yaml.FutureLineWrap()
}

// MarshalJSON produces the JSON resulting from encoding an object composed of
// the fields in v (which must me a struct) and the keys in extensions.
func MarshalJSON(v interface{}, extensions map[string]interface{}) ([]byte, error) {
	marshaled, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	if len(extensions) == 0 {
		return marshaled, nil
	}
	var unmarshaled interface{}
	if err := json.Unmarshal(marshaled, &unmarshaled); err != nil {
		return nil, err
	}
	asserted := unmarshaled.(map[string]interface{})
	for k, v := range extensions {
		asserted[k] = v
	}
	merged, err := json.Marshal(asserted)
	if err != nil {
		return nil, err
	}
	return merged, nil
}

// MarshalYAML produces the JSON resulting from encoding an object composed of
// the fields in v (which must me a struct) and the keys in extensions.
func MarshalYAML(v interface{}, extensions map[string]interface{}) (interface{}, error) {
	if len(extensions) == 0 {
		return v, nil
	}
	marshaled, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	var unmarshaled map[string]interface{}
	if err := yaml.Unmarshal(marshaled, &unmarshaled); err != nil {
		return nil, err
	}
	for k, v := range extensions {
		unmarshaled[k] = v
	}
	return unmarshaled, nil
}
