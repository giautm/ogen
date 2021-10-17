// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
)

func encodeFoobarGetResponse(response FoobarGetResponse, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if err := response.WriteJSONTo(w); err != nil {
			return err
		}
		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	default:
		return fmt.Errorf("/foobar: unexpected response type: %T", response)
	}
}

func encodeFoobarPutResponse(response FoobarPutDefault, w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

func encodeFoobarPostResponse(response FoobarPostResponse, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if err := response.WriteJSONTo(w); err != nil {
			return err
		}
		return nil
	case *NotFound:
		w.WriteHeader(404)
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		if err := response.Response.WriteJSONTo(w); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/foobar: unexpected response type: %T", response)
	}
}

func encodePetGetResponse(response PetGetResponse, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if err := response.WriteJSONTo(w); err != nil {
			return err
		}
		return nil
	case *PetGetDefaultStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		if err := response.Response.WriteJSONTo(w); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/pet: unexpected response type: %T", response)
	}
}

func encodePetCreateResponse(response Pet, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodePetGetByNameResponse(response Pet, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}
