// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

func (s Pets) Validate() error {
	if s == nil {
		return errors.New("nil is invalid value")
	}
	return nil // return 1
}
