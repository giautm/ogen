// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeAddStickerToSetRequestJSON(
	req AddStickerToSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeAnswerCallbackQueryRequestJSON(
	req AnswerCallbackQuery,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeAnswerInlineQueryRequestJSON(
	req AnswerInlineQuery,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeAnswerPreCheckoutQueryRequestJSON(
	req AnswerPreCheckoutQuery,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeAnswerShippingQueryRequestJSON(
	req AnswerShippingQuery,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeApproveChatJoinRequestRequestJSON(
	req ApproveChatJoinRequest,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeBanChatMemberRequestJSON(
	req BanChatMember,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeBanChatSenderChatRequestJSON(
	req BanChatSenderChat,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeCopyMessageRequestJSON(
	req CopyMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeCreateChatInviteLinkRequestJSON(
	req CreateChatInviteLink,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeCreateNewStickerSetRequestJSON(
	req CreateNewStickerSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeclineChatJoinRequestRequestJSON(
	req DeclineChatJoinRequest,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteChatPhotoRequestJSON(
	req DeleteChatPhoto,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteChatStickerSetRequestJSON(
	req DeleteChatStickerSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteMessageRequestJSON(
	req DeleteMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteMyCommandsRequestJSON(
	req OptDeleteMyCommands,
	r *http.Request,
) error {
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	if req.Set {
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteStickerFromSetRequestJSON(
	req DeleteStickerFromSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeDeleteWebhookRequestJSON(
	req OptDeleteWebhook,
	r *http.Request,
) error {
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	if req.Set {
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditChatInviteLinkRequestJSON(
	req EditChatInviteLink,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditMessageCaptionRequestJSON(
	req EditMessageCaption,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditMessageLiveLocationRequestJSON(
	req EditMessageLiveLocation,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditMessageMediaRequestJSON(
	req EditMessageMedia,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditMessageReplyMarkupRequestJSON(
	req EditMessageReplyMarkup,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeEditMessageTextRequestJSON(
	req EditMessageText,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeExportChatInviteLinkRequestJSON(
	req ExportChatInviteLink,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeForwardMessageRequestJSON(
	req ForwardMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetChatRequestJSON(
	req GetChat,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetChatAdministratorsRequestJSON(
	req GetChatAdministrators,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetChatMemberRequestJSON(
	req GetChatMember,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetChatMemberCountRequestJSON(
	req GetChatMemberCount,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetFileRequestJSON(
	req GetFile,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetGameHighScoresRequestJSON(
	req GetGameHighScores,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetMyCommandsRequestJSON(
	req OptGetMyCommands,
	r *http.Request,
) error {
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	if req.Set {
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetStickerSetRequestJSON(
	req GetStickerSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetUpdatesRequestJSON(
	req OptGetUpdates,
	r *http.Request,
) error {
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	if req.Set {
		req.Encode(e)
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeGetUserProfilePhotosRequestJSON(
	req GetUserProfilePhotos,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeLeaveChatRequestJSON(
	req LeaveChat,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodePinChatMessageRequestJSON(
	req PinChatMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodePromoteChatMemberRequestJSON(
	req PromoteChatMember,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeRestrictChatMemberRequestJSON(
	req RestrictChatMember,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeRevokeChatInviteLinkRequestJSON(
	req RevokeChatInviteLink,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendAnimationRequestJSON(
	req SendAnimation,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendAudioRequestJSON(
	req SendAudio,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendChatActionRequestJSON(
	req SendChatAction,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendContactRequestJSON(
	req SendContact,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendDiceRequestJSON(
	req SendDice,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendDocumentRequestJSON(
	req SendDocument,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendGameRequestJSON(
	req SendGame,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendInvoiceRequestJSON(
	req SendInvoice,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendLocationRequestJSON(
	req SendLocation,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendMediaGroupRequestJSON(
	req SendMediaGroup,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendMessageRequestJSON(
	req SendMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendPhotoRequestJSON(
	req SendPhoto,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendPollRequestJSON(
	req SendPoll,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendStickerRequestJSON(
	req SendSticker,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendVenueRequestJSON(
	req SendVenue,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendVideoRequestJSON(
	req SendVideo,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendVideoNoteRequestJSON(
	req SendVideoNote,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSendVoiceRequestJSON(
	req SendVoice,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatAdministratorCustomTitleRequestJSON(
	req SetChatAdministratorCustomTitle,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatDescriptionRequestJSON(
	req SetChatDescription,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatPermissionsRequestJSON(
	req SetChatPermissions,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatPhotoRequestJSON(
	req SetChatPhoto,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatStickerSetRequestJSON(
	req SetChatStickerSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetChatTitleRequestJSON(
	req SetChatTitle,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetGameScoreRequestJSON(
	req SetGameScore,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetMyCommandsRequestJSON(
	req SetMyCommands,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetPassportDataErrorsRequestJSON(
	req SetPassportDataErrors,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetStickerPositionInSetRequestJSON(
	req SetStickerPositionInSet,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetStickerSetThumbRequestJSON(
	req SetStickerSetThumb,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeSetWebhookRequestJSON(
	req SetWebhook,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeStopMessageLiveLocationRequestJSON(
	req StopMessageLiveLocation,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeStopPollRequestJSON(
	req StopPoll,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeUnbanChatMemberRequestJSON(
	req UnbanChatMember,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeUnbanChatSenderChatRequestJSON(
	req UnbanChatSenderChat,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeUnpinAllChatMessagesRequestJSON(
	req UnpinAllChatMessages,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeUnpinChatMessageRequestJSON(
	req UnpinChatMessage,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
func encodeUploadStickerFileRequestJSON(
	req UploadStickerFile,
	r *http.Request,
) error {
	e := jx.GetEncoder()

	req.Encode(e)
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), "application/json")
	return nil
}
