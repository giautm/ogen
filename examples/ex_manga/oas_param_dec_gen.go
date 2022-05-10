// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func decodeGetBookParams(args [1]string, r *http.Request) (params GetBookParams, _ error) {
	// Decode path: book_id.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "book_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.BookID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: book_id: not specified")
		}
	}
	return params, nil
}

func decodeSearchParams(args [0]string, r *http.Request) (params SearchParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: query.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "query",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Query = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: query: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: page.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	return params, nil
}

func decodeSearchByTagIDParams(args [0]string, r *http.Request) (params SearchByTagIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: tag_id.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "tag_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.TagID = c
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: tag_id: parse")
			}
		} else {
			return params, errors.Wrap(err, "query")
		}
	}
	// Decode query: page.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: page: parse")
			}
		}
	}
	return params, nil
}
