// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeDataCreateRequestJSON(
	req OptData,
	r *http.Request,
) error {
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	if req.Set {
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
