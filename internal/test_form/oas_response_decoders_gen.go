// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

func decodeTestFormURLEncodedResponse(resp *http.Response) (res TestFormURLEncodedOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return TestFormURLEncodedOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeTestMultipartResponse(resp *http.Response) (res TestMultipartOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return TestMultipartOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeTestMultipartUploadResponse(resp *http.Response) (res TestMultipartUploadOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}

			d := jx.DecodeBytes(b)
			var response TestMultipartUploadOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeTestShareFormSchemaResponse(resp *http.Response) (res TestShareFormSchemaOK, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return TestShareFormSchemaOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
