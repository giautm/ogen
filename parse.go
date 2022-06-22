package ogen

import (
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Parse parses JSON/YAML into OpenAPI Spec.
func Parse(data []byte) (s *Spec, err error) {
	s = &Spec{}
	if !jx.Valid(data) {
		if err := yaml.Unmarshal(data, s); err != nil {
			return nil, errors.Wrap(err, "yaml")
		}
	} else {
		if err := json.Unmarshal(data, s); err != nil {
			return nil, errors.Wrap(wrapLineOffset(data, err), "json")
		}
	}
	return s, nil
}
