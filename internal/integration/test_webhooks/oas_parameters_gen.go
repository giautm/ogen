// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// UpdateWebhookParams is parameters of updateWebhook operation.
type UpdateWebhookParams struct {
	EventType     string
	XWebhookToken OptString
}

func unpackUpdateWebhookParams(packed map[string]any) (params UpdateWebhookParams) {
	params.EventType = packed["event_type"].(string)
	if v, ok := packed["X-Webhook-Token"]; ok {
		params.XWebhookToken = v.(OptString)
	}
	return params
}

func decodeUpdateWebhookParams(args [0]string, r *http.Request) (params UpdateWebhookParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	h := uri.NewHeaderDecoder(r.Header)
	// Decode query: event_type.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "event_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.EventType = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: event_type: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode header: X-Webhook-Token.
	{
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Webhook-Token",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotXWebhookTokenVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotXWebhookTokenVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.XWebhookToken.SetTo(paramsDotXWebhookTokenVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "header: X-Webhook-Token: parse")
			}
		}
	}
	return params, nil
}
