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

func encodeAnswerCallbackQueryPostRequest(req AnswerCallbackQueryPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *AnswerCallbackQueryPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *AnswerCallbackQueryPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *AnswerCallbackQueryPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeAnswerPreCheckoutQueryPostRequest(req AnswerPreCheckoutQueryPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *AnswerPreCheckoutQueryPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *AnswerPreCheckoutQueryPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *AnswerPreCheckoutQueryPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeAnswerShippingQueryPostRequest(req AnswerShippingQueryPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *AnswerShippingQueryPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *AnswerShippingQueryPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *AnswerShippingQueryPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeDeleteStickerFromSetPostRequest(req DeleteStickerFromSetPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *DeleteStickerFromSetPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *DeleteStickerFromSetPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *DeleteStickerFromSetPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeDeleteWebhookPostRequest(req DeleteWebhookPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *DeleteWebhookPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *DeleteWebhookPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *DeleteWebhookPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeGetFilePostRequest(req GetFilePostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *GetFilePostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *GetFilePostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *GetFilePostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeGetGameHighScoresPostRequest(req GetGameHighScoresPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *GetGameHighScoresPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *GetGameHighScoresPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *GetGameHighScoresPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeGetStickerSetPostRequest(req GetStickerSetPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *GetStickerSetPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *GetStickerSetPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *GetStickerSetPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeGetUpdatesPostRequest(req GetUpdatesPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *GetUpdatesPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *GetUpdatesPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *GetUpdatesPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeGetUserProfilePhotosPostRequest(req GetUserProfilePhotosPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *GetUserProfilePhotosPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *GetUserProfilePhotosPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *GetUserProfilePhotosPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeSendGamePostRequest(req SendGamePostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *SendGamePostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *SendGamePostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *SendGamePostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeSendInvoicePostRequest(req SendInvoicePostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *SendInvoicePostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *SendInvoicePostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *SendInvoicePostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeSetMyCommandsPostRequest(req SetMyCommandsPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *SetMyCommandsPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *SetMyCommandsPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *SetMyCommandsPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeSetStickerPositionInSetPostRequest(req SetStickerPositionInSetPostReq) (data []byte, contentType string, err error) {
	switch req := req.(type) {
	case *SetStickerPositionInSetPostApplicationJSONReq:
		return json.Encode(req), "application/json", nil
	case *SetStickerPositionInSetPostApplicationXWwwFormUrlencodedReq:
		return nil, "", fmt.Errorf("application/x-www-form-urlencoded encoder not implemented")
	case *SetStickerPositionInSetPostMultipartFormDataReq:
		return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
	default:
		return nil, "", fmt.Errorf("unexpected request type: %T", req)
	}
}

func encodeSetWebhookPostRequest(req SetWebhookPostMultipartFormDataReq) (data []byte, contentType string, err error) {
	return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
}

func encodeUploadStickerFilePostRequest(req UploadStickerFilePostMultipartFormDataReq) (data []byte, contentType string, err error) {
	return nil, "", fmt.Errorf("multipart/form-data encoder not implemented")
}
