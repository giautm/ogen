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

func decodeAnswerCallbackQueryPostResponse(resp *http.Response) (res AnswerCallbackQueryPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response AnswerCallbackQueryPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeAnswerPreCheckoutQueryPostResponse(resp *http.Response) (res AnswerPreCheckoutQueryPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response AnswerPreCheckoutQueryPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeAnswerShippingQueryPostResponse(resp *http.Response) (res AnswerShippingQueryPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response AnswerShippingQueryPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeClosePostResponse(resp *http.Response) (res ClosePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ClosePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDeleteStickerFromSetPostResponse(resp *http.Response) (res DeleteStickerFromSetPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response DeleteStickerFromSetPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeDeleteWebhookPostResponse(resp *http.Response) (res DeleteWebhookPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response DeleteWebhookPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetFilePostResponse(resp *http.Response) (res GetFilePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetFilePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetGameHighScoresPostResponse(resp *http.Response) (res GetGameHighScoresPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetGameHighScoresPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMePostResponse(resp *http.Response) (res GetMePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetMePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetMyCommandsPostResponse(resp *http.Response) (res GetMyCommandsPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetMyCommandsPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetStickerSetPostResponse(resp *http.Response) (res GetStickerSetPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetStickerSetPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetUpdatesPostResponse(resp *http.Response) (res GetUpdatesPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetUpdatesPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetUserProfilePhotosPostResponse(resp *http.Response) (res GetUserProfilePhotosPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetUserProfilePhotosPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeGetWebhookInfoPostResponse(resp *http.Response) (res GetWebhookInfoPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response GetWebhookInfoPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeLogOutPostResponse(resp *http.Response) (res LogOutPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response LogOutPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSendGamePostResponse(resp *http.Response) (res SendGamePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response SendGamePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSendInvoicePostResponse(resp *http.Response) (res SendInvoicePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response SendInvoicePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSetMyCommandsPostResponse(resp *http.Response) (res SetMyCommandsPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response SetMyCommandsPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSetStickerPositionInSetPostResponse(resp *http.Response) (res SetStickerPositionInSetPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response SetStickerPositionInSetPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeSetWebhookPostResponse(resp *http.Response) (res SetWebhookPostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response SetWebhookPostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}

func decodeUploadStickerFilePostResponse(resp *http.Response) (res UploadStickerFilePostResponse, err error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response UploadStickerFilePostOK
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			if err := response.ReadJSONFrom(resp.Body); err != nil {
				return res, err
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			return res, fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
		}
	}
}
