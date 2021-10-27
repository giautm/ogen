// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
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
	_ = net.IP{}
)

func encodeFoobarPostRequest(req Pet) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	req.WriteJSON(j)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePetCreateRequest(req PetCreateReq) (data *bytes.Buffer, contentType string, err error) {
	switch req := req.(type) {
	case *Pet:
		buf := json.GetBuffer()
		j := json.GetStream(buf)
		defer json.PutStream(j)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		req.WriteJSON(j)
		if err := j.Flush(); err != nil {
			json.PutBuffer(buf)
			return nil, "", err
		}

		return buf, "application/json", nil
	case *PetCreateReqTextPlain:
		return nil, "", fmt.Errorf("text/plain encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodePetUpdateNameAliasPostRequest(req PetName) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	// Unsupported kind "alias".
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}

func encodePetUpdateNamePostRequest(req string) (data *bytes.Buffer, contentType string, err error) {
	buf := json.GetBuffer()
	j := json.GetStream(buf)
	defer json.PutStream(j)
	more := json.NewMore(j)
	defer more.Reset()
	more.More()
	j.WriteString(req)
	if err := j.Flush(); err != nil {
		json.PutBuffer(buf)
		return nil, "", err
	}

	return buf, "application/json", nil
}
