// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

func decodeListPetsParams(args [0]string, r *http.Request) (params ListPetsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: limit.
	{
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int32
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(s)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return params, errors.Wrap(err, "query: limit: parse")
			}
		}
	}
	return params, nil
}

func decodeShowPetByIdParams(args [1]string, r *http.Request) (params ShowPetByIdParams, _ error) {
	// Decode path: petId.
	{
		param := args[0]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "petId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New("path: petId: not specified")
		}
	}
	return params, nil
}
