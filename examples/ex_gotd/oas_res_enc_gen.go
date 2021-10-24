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

func encodeAddStickerToSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerCallbackQueryResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerInlineQueryResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerPreCheckoutQueryResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeAnswerShippingQueryResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeBanChatMemberResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCopyMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCreateChatInviteLinkResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeCreateNewStickerSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteChatPhotoResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteChatStickerSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteMyCommandsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteStickerFromSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeDeleteWebhookResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditChatInviteLinkResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageCaptionResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageLiveLocationResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageMediaResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageReplyMarkupResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeEditMessageTextResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeExportChatInviteLinkResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeForwardMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatAdministratorsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatMemberResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetChatMemberCountResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetFileResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetGameHighScoresResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetMeResponse(response User, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetMyCommandsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetStickerSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetUpdatesResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeGetUserProfilePhotosResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeLeaveChatResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodePinChatMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodePromoteChatMemberResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeRestrictChatMemberResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeRevokeChatInviteLinkResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendAnimationResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendAudioResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendChatActionResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendContactResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendDiceResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendDocumentResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendGameResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendInvoiceResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendLocationResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendMediaGroupResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendPhotoResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendPollResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendStickerResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVenueResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVideoResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVideoNoteResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSendVoiceResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatAdministratorCustomTitleResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatDescriptionResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatPermissionsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatPhotoResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatStickerSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetChatTitleResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetGameScoreResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetMyCommandsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetPassportDataErrorsResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetStickerPositionInSetResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetStickerSetThumbResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeSetWebhookResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeStopMessageLiveLocationResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeStopPollResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnbanChatMemberResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnpinAllChatMessagesResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUnpinChatMessageResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}

func encodeUploadStickerFileResponse(response Result, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := response.WriteJSONTo(w); err != nil {
		return err
	}
	return nil
}
