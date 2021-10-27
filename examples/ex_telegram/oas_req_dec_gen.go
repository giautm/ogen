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
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
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
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
)

func decodeAnswerCallbackQueryPostRequest(r *http.Request) (req AnswerCallbackQueryPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerCallbackQueryPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerCallbackQueryPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerCallbackQueryPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeAnswerPreCheckoutQueryPostRequest(r *http.Request) (req AnswerPreCheckoutQueryPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerPreCheckoutQueryPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerPreCheckoutQueryPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerPreCheckoutQueryPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeAnswerShippingQueryPostRequest(r *http.Request) (req AnswerShippingQueryPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerShippingQueryPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerShippingQueryPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerShippingQueryPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeDeleteStickerFromSetPostRequest(r *http.Request) (req DeleteStickerFromSetPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request DeleteStickerFromSetPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request DeleteStickerFromSetPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request DeleteStickerFromSetPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeDeleteWebhookPostRequest(r *http.Request) (req DeleteWebhookPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request DeleteWebhookPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request DeleteWebhookPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request DeleteWebhookPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetFilePostRequest(r *http.Request) (req GetFilePostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetFilePostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetFilePostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetFilePostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetGameHighScoresPostRequest(r *http.Request) (req GetGameHighScoresPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetGameHighScoresPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetGameHighScoresPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetGameHighScoresPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetStickerSetPostRequest(r *http.Request) (req GetStickerSetPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetStickerSetPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetStickerSetPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetStickerSetPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetUpdatesPostRequest(r *http.Request) (req GetUpdatesPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetUpdatesPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetUpdatesPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetUpdatesPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetUserProfilePhotosPostRequest(r *http.Request) (req GetUserProfilePhotosPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetUserProfilePhotosPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetUserProfilePhotosPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetUserProfilePhotosPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSendGamePostRequest(r *http.Request) (req SendGamePostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SendGamePostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SendGamePostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SendGamePostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSendInvoicePostRequest(r *http.Request) (req SendInvoicePostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SendInvoicePostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SendInvoicePostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SendInvoicePostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetMyCommandsPostRequest(r *http.Request) (req SetMyCommandsPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SetMyCommandsPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		if err := func() error {
			if err := request.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SetMyCommandsPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SetMyCommandsPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetStickerPositionInSetPostRequest(r *http.Request) (req SetStickerPositionInSetPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SetStickerPositionInSetPostReqApplicationJSON
		i := json.GetIterator()
		defer json.PutIterator(i)
		i.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.ReadJSON(i); err != nil {
				return err
			}
			return i.Error
		}(); err != nil {
			return req, err
		}
		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SetStickerPositionInSetPostReqApplicationXWwwFormUrlencoded
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SetStickerPositionInSetPostReqMultipartFormData
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetWebhookPostRequest(r *http.Request) (req SetWebhookPostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "multipart/form-data":
		var request SetWebhookPostReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeUploadStickerFilePostRequest(r *http.Request) (req UploadStickerFilePostReq, err error) {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)
	if _, err := io.Copy(buf, r.Body); err != nil {
		return req, err
	}

	switch r.Header.Get("Content-Type") {
	case "multipart/form-data":
		var request UploadStickerFilePostReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
