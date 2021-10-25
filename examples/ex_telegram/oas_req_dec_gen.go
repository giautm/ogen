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

func decodeAnswerCallbackQueryPostRequest(r *http.Request) (req AnswerCallbackQueryPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerCallbackQueryPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerCallbackQueryPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerCallbackQueryPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeAnswerPreCheckoutQueryPostRequest(r *http.Request) (req AnswerPreCheckoutQueryPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerPreCheckoutQueryPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerPreCheckoutQueryPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerPreCheckoutQueryPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeAnswerShippingQueryPostRequest(r *http.Request) (req AnswerShippingQueryPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request AnswerShippingQueryPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request AnswerShippingQueryPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request AnswerShippingQueryPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeDeleteStickerFromSetPostRequest(r *http.Request) (req DeleteStickerFromSetPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request DeleteStickerFromSetPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request DeleteStickerFromSetPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request DeleteStickerFromSetPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeDeleteWebhookPostRequest(r *http.Request) (req DeleteWebhookPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request DeleteWebhookPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request DeleteWebhookPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request DeleteWebhookPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetFilePostRequest(r *http.Request) (req GetFilePostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetFilePostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetFilePostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetFilePostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetGameHighScoresPostRequest(r *http.Request) (req GetGameHighScoresPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetGameHighScoresPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetGameHighScoresPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetGameHighScoresPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetStickerSetPostRequest(r *http.Request) (req GetStickerSetPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetStickerSetPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetStickerSetPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetStickerSetPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetUpdatesPostRequest(r *http.Request) (req GetUpdatesPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetUpdatesPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetUpdatesPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetUpdatesPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeGetUserProfilePhotosPostRequest(r *http.Request) (req GetUserProfilePhotosPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request GetUserProfilePhotosPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request GetUserProfilePhotosPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request GetUserProfilePhotosPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSendGamePostRequest(r *http.Request) (req SendGamePostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SendGamePostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SendGamePostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SendGamePostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSendInvoicePostRequest(r *http.Request) (req SendInvoicePostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SendInvoicePostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SendInvoicePostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SendInvoicePostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetMyCommandsPostRequest(r *http.Request) (req SetMyCommandsPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SetMyCommandsPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}
		if err := request.Validate(); err != nil {
			return req, fmt.Errorf("validate: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SetMyCommandsPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SetMyCommandsPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetStickerPositionInSetPostRequest(r *http.Request) (req SetStickerPositionInSetPostReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		var request SetStickerPositionInSetPostApplicationJSONReq
		if err := request.ReadJSONFrom(r.Body); err != nil {
			return req, fmt.Errorf("json: %w", err)
		}

		return &request, nil
	case "application/x-www-form-urlencoded":
		var request SetStickerPositionInSetPostApplicationXWwwFormUrlencodedReq
		_ = request
		return req, fmt.Errorf("application/x-www-form-urlencoded decoder not implemented")
	case "multipart/form-data":
		var request SetStickerPositionInSetPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeSetWebhookPostRequest(r *http.Request) (req SetWebhookPostMultipartFormDataReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "multipart/form-data":
		var request SetWebhookPostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}

func decodeUploadStickerFilePostRequest(r *http.Request) (req UploadStickerFilePostMultipartFormDataReq, err error) {
	switch r.Header.Get("Content-Type") {
	case "multipart/form-data":
		var request UploadStickerFilePostMultipartFormDataReq
		_ = request
		return req, fmt.Errorf("multipart/form-data decoder not implemented")
	default:
		return req, fmt.Errorf("unexpected content-type: %s", r.Header.Get("Content-Type"))
	}
}
