// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/uri"
)

// ListPetsParams is parameters of listPets operation.
type ListPetsParams struct {
	// How many items to return at one time (max 100).
	Limit OptInt32
}

func unpackListPetsParams(packed map[string]any) (params ListPetsParams) {
	if v, ok := packed["limit"]; ok {
		params.Limit = v.(OptInt32)
	}
	return params
}

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
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
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

// ShowPetByIdParams is parameters of showPetById operation.
type ShowPetByIdParams struct {
	// The id of the pet to retrieve.
	PetId string
}

func unpackShowPetByIdParams(packed map[string]any) (params ShowPetByIdParams) {
	params.PetId = packed["petId"].(string)
	return params
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
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.PetId = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, "path: petId: parse")
			}
		} else {
			return params, errors.New("path: petId: not specified")
		}
	}
	return params, nil
}
