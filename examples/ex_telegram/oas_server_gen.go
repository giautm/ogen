// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddStickerToSet implements addStickerToSet operation.
	//
	// POST /addStickerToSet
	AddStickerToSet(ctx context.Context, req AddStickerToSet) (Result, error)
	// AnswerCallbackQuery implements answerCallbackQuery operation.
	//
	// POST /answerCallbackQuery
	AnswerCallbackQuery(ctx context.Context, req AnswerCallbackQuery) (Result, error)
	// AnswerInlineQuery implements answerInlineQuery operation.
	//
	// POST /answerInlineQuery
	AnswerInlineQuery(ctx context.Context, req AnswerInlineQuery) (Result, error)
	// AnswerPreCheckoutQuery implements answerPreCheckoutQuery operation.
	//
	// POST /answerPreCheckoutQuery
	AnswerPreCheckoutQuery(ctx context.Context, req AnswerPreCheckoutQuery) (Result, error)
	// AnswerShippingQuery implements answerShippingQuery operation.
	//
	// POST /answerShippingQuery
	AnswerShippingQuery(ctx context.Context, req AnswerShippingQuery) (Result, error)
	// ApproveChatJoinRequest implements approveChatJoinRequest operation.
	//
	// POST /approveChatJoinRequest
	ApproveChatJoinRequest(ctx context.Context, req ApproveChatJoinRequest) (Result, error)
	// BanChatMember implements banChatMember operation.
	//
	// POST /banChatMember
	BanChatMember(ctx context.Context, req BanChatMember) (Result, error)
	// BanChatSenderChat implements banChatSenderChat operation.
	//
	// POST /banChatSenderChat
	BanChatSenderChat(ctx context.Context, req BanChatSenderChat) (Result, error)
	// Close implements close operation.
	//
	// POST /close
	Close(ctx context.Context) (Result, error)
	// CopyMessage implements copyMessage operation.
	//
	// POST /copyMessage
	CopyMessage(ctx context.Context, req CopyMessage) (ResultMessageId, error)
	// CreateChatInviteLink implements createChatInviteLink operation.
	//
	// POST /createChatInviteLink
	CreateChatInviteLink(ctx context.Context, req CreateChatInviteLink) (ResultChatInviteLink, error)
	// CreateNewStickerSet implements createNewStickerSet operation.
	//
	// POST /createNewStickerSet
	CreateNewStickerSet(ctx context.Context, req CreateNewStickerSet) (Result, error)
	// DeclineChatJoinRequest implements declineChatJoinRequest operation.
	//
	// POST /declineChatJoinRequest
	DeclineChatJoinRequest(ctx context.Context, req DeclineChatJoinRequest) (Result, error)
	// DeleteChatPhoto implements deleteChatPhoto operation.
	//
	// POST /deleteChatPhoto
	DeleteChatPhoto(ctx context.Context, req DeleteChatPhoto) (Result, error)
	// DeleteChatStickerSet implements deleteChatStickerSet operation.
	//
	// POST /deleteChatStickerSet
	DeleteChatStickerSet(ctx context.Context, req DeleteChatStickerSet) (Result, error)
	// DeleteMessage implements deleteMessage operation.
	//
	// POST /deleteMessage
	DeleteMessage(ctx context.Context, req DeleteMessage) (Result, error)
	// DeleteMyCommands implements deleteMyCommands operation.
	//
	// POST /deleteMyCommands
	DeleteMyCommands(ctx context.Context, req OptDeleteMyCommands) (Result, error)
	// DeleteStickerFromSet implements deleteStickerFromSet operation.
	//
	// POST /deleteStickerFromSet
	DeleteStickerFromSet(ctx context.Context, req DeleteStickerFromSet) (Result, error)
	// DeleteWebhook implements deleteWebhook operation.
	//
	// POST /deleteWebhook
	DeleteWebhook(ctx context.Context, req OptDeleteWebhook) (Result, error)
	// EditChatInviteLink implements editChatInviteLink operation.
	//
	// POST /editChatInviteLink
	EditChatInviteLink(ctx context.Context, req EditChatInviteLink) (ResultChatInviteLink, error)
	// EditMessageCaption implements editMessageCaption operation.
	//
	// POST /editMessageCaption
	EditMessageCaption(ctx context.Context, req EditMessageCaption) (Result, error)
	// EditMessageLiveLocation implements editMessageLiveLocation operation.
	//
	// POST /editMessageLiveLocation
	EditMessageLiveLocation(ctx context.Context, req EditMessageLiveLocation) (Result, error)
	// EditMessageMedia implements editMessageMedia operation.
	//
	// POST /editMessageMedia
	EditMessageMedia(ctx context.Context, req EditMessageMedia) (Result, error)
	// EditMessageReplyMarkup implements editMessageReplyMarkup operation.
	//
	// POST /editMessageReplyMarkup
	EditMessageReplyMarkup(ctx context.Context, req EditMessageReplyMarkup) (Result, error)
	// EditMessageText implements editMessageText operation.
	//
	// POST /editMessageText
	EditMessageText(ctx context.Context, req EditMessageText) (Result, error)
	// ExportChatInviteLink implements exportChatInviteLink operation.
	//
	// POST /exportChatInviteLink
	ExportChatInviteLink(ctx context.Context, req ExportChatInviteLink) (ResultString, error)
	// ForwardMessage implements forwardMessage operation.
	//
	// POST /forwardMessage
	ForwardMessage(ctx context.Context, req ForwardMessage) (ResultMessage, error)
	// GetChat implements getChat operation.
	//
	// POST /getChat
	GetChat(ctx context.Context, req GetChat) (ResultChat, error)
	// GetChatAdministrators implements getChatAdministrators operation.
	//
	// POST /getChatAdministrators
	GetChatAdministrators(ctx context.Context, req GetChatAdministrators) (ResultArrayOfChatMember, error)
	// GetChatMember implements getChatMember operation.
	//
	// POST /getChatMember
	GetChatMember(ctx context.Context, req GetChatMember) (ResultChatMember, error)
	// GetChatMemberCount implements getChatMemberCount operation.
	//
	// POST /getChatMemberCount
	GetChatMemberCount(ctx context.Context, req GetChatMemberCount) (ResultInt, error)
	// GetFile implements getFile operation.
	//
	// POST /getFile
	GetFile(ctx context.Context, req GetFile) (Result, error)
	// GetGameHighScores implements getGameHighScores operation.
	//
	// POST /getGameHighScores
	GetGameHighScores(ctx context.Context, req GetGameHighScores) (ResultArrayOfGameHighScore, error)
	// GetMe implements getMe operation.
	//
	// POST /getMe
	GetMe(ctx context.Context) (ResultUser, error)
	// GetMyCommands implements getMyCommands operation.
	//
	// POST /getMyCommands
	GetMyCommands(ctx context.Context, req OptGetMyCommands) (ResultArrayOfBotCommand, error)
	// GetStickerSet implements getStickerSet operation.
	//
	// POST /getStickerSet
	GetStickerSet(ctx context.Context, req GetStickerSet) (Result, error)
	// GetUpdates implements getUpdates operation.
	//
	// POST /getUpdates
	GetUpdates(ctx context.Context, req OptGetUpdates) (ResultArrayOfUpdate, error)
	// GetUserProfilePhotos implements getUserProfilePhotos operation.
	//
	// POST /getUserProfilePhotos
	GetUserProfilePhotos(ctx context.Context, req GetUserProfilePhotos) (ResultUserProfilePhotos, error)
	// GetWebhookInfo implements getWebhookInfo operation.
	//
	// POST /getWebhookInfo
	GetWebhookInfo(ctx context.Context) (ResultWebhookInfo, error)
	// LeaveChat implements leaveChat operation.
	//
	// POST /leaveChat
	LeaveChat(ctx context.Context, req LeaveChat) (Result, error)
	// LogOut implements logOut operation.
	//
	// POST /logOut
	LogOut(ctx context.Context) (Result, error)
	// PinChatMessage implements pinChatMessage operation.
	//
	// POST /pinChatMessage
	PinChatMessage(ctx context.Context, req PinChatMessage) (Result, error)
	// PromoteChatMember implements promoteChatMember operation.
	//
	// POST /promoteChatMember
	PromoteChatMember(ctx context.Context, req PromoteChatMember) (Result, error)
	// RestrictChatMember implements restrictChatMember operation.
	//
	// POST /restrictChatMember
	RestrictChatMember(ctx context.Context, req RestrictChatMember) (Result, error)
	// RevokeChatInviteLink implements revokeChatInviteLink operation.
	//
	// POST /revokeChatInviteLink
	RevokeChatInviteLink(ctx context.Context, req RevokeChatInviteLink) (ResultChatInviteLink, error)
	// SendAnimation implements sendAnimation operation.
	//
	// POST /sendAnimation
	SendAnimation(ctx context.Context, req SendAnimation) (ResultMessage, error)
	// SendAudio implements sendAudio operation.
	//
	// POST /sendAudio
	SendAudio(ctx context.Context, req SendAudio) (ResultMessage, error)
	// SendChatAction implements sendChatAction operation.
	//
	// POST /sendChatAction
	SendChatAction(ctx context.Context, req SendChatAction) (Result, error)
	// SendContact implements sendContact operation.
	//
	// POST /sendContact
	SendContact(ctx context.Context, req SendContact) (ResultMessage, error)
	// SendDice implements sendDice operation.
	//
	// POST /sendDice
	SendDice(ctx context.Context, req SendDice) (ResultMessage, error)
	// SendDocument implements sendDocument operation.
	//
	// POST /sendDocument
	SendDocument(ctx context.Context, req SendDocument) (ResultMessage, error)
	// SendGame implements sendGame operation.
	//
	// POST /sendGame
	SendGame(ctx context.Context, req SendGame) (ResultMessage, error)
	// SendInvoice implements sendInvoice operation.
	//
	// POST /sendInvoice
	SendInvoice(ctx context.Context, req SendInvoice) (ResultMessage, error)
	// SendLocation implements sendLocation operation.
	//
	// POST /sendLocation
	SendLocation(ctx context.Context, req SendLocation) (ResultMessage, error)
	// SendMediaGroup implements sendMediaGroup operation.
	//
	// POST /sendMediaGroup
	SendMediaGroup(ctx context.Context, req SendMediaGroup) (ResultArrayOfMessage, error)
	// SendMessage implements sendMessage operation.
	//
	// POST /sendMessage
	SendMessage(ctx context.Context, req SendMessage) (ResultMessage, error)
	// SendPhoto implements sendPhoto operation.
	//
	// POST /sendPhoto
	SendPhoto(ctx context.Context, req SendPhoto) (ResultMessage, error)
	// SendPoll implements sendPoll operation.
	//
	// POST /sendPoll
	SendPoll(ctx context.Context, req SendPoll) (ResultMessage, error)
	// SendSticker implements sendSticker operation.
	//
	// POST /sendSticker
	SendSticker(ctx context.Context, req SendSticker) (ResultMessage, error)
	// SendVenue implements sendVenue operation.
	//
	// POST /sendVenue
	SendVenue(ctx context.Context, req SendVenue) (ResultMessage, error)
	// SendVideo implements sendVideo operation.
	//
	// POST /sendVideo
	SendVideo(ctx context.Context, req SendVideo) (ResultMessage, error)
	// SendVideoNote implements sendVideoNote operation.
	//
	// POST /sendVideoNote
	SendVideoNote(ctx context.Context, req SendVideoNote) (ResultMessage, error)
	// SendVoice implements sendVoice operation.
	//
	// POST /sendVoice
	SendVoice(ctx context.Context, req SendVoice) (ResultMessage, error)
	// SetChatAdministratorCustomTitle implements setChatAdministratorCustomTitle operation.
	//
	// POST /setChatAdministratorCustomTitle
	SetChatAdministratorCustomTitle(ctx context.Context, req SetChatAdministratorCustomTitle) (Result, error)
	// SetChatDescription implements setChatDescription operation.
	//
	// POST /setChatDescription
	SetChatDescription(ctx context.Context, req SetChatDescription) (Result, error)
	// SetChatPermissions implements setChatPermissions operation.
	//
	// POST /setChatPermissions
	SetChatPermissions(ctx context.Context, req SetChatPermissions) (Result, error)
	// SetChatPhoto implements setChatPhoto operation.
	//
	// POST /setChatPhoto
	SetChatPhoto(ctx context.Context, req SetChatPhoto) (Result, error)
	// SetChatStickerSet implements setChatStickerSet operation.
	//
	// POST /setChatStickerSet
	SetChatStickerSet(ctx context.Context, req SetChatStickerSet) (Result, error)
	// SetChatTitle implements setChatTitle operation.
	//
	// POST /setChatTitle
	SetChatTitle(ctx context.Context, req SetChatTitle) (Result, error)
	// SetGameScore implements setGameScore operation.
	//
	// POST /setGameScore
	SetGameScore(ctx context.Context, req SetGameScore) (Result, error)
	// SetMyCommands implements setMyCommands operation.
	//
	// POST /setMyCommands
	SetMyCommands(ctx context.Context, req SetMyCommands) (Result, error)
	// SetPassportDataErrors implements setPassportDataErrors operation.
	//
	// POST /setPassportDataErrors
	SetPassportDataErrors(ctx context.Context, req SetPassportDataErrors) (Result, error)
	// SetStickerPositionInSet implements setStickerPositionInSet operation.
	//
	// POST /setStickerPositionInSet
	SetStickerPositionInSet(ctx context.Context, req SetStickerPositionInSet) (Result, error)
	// SetStickerSetThumb implements setStickerSetThumb operation.
	//
	// POST /setStickerSetThumb
	SetStickerSetThumb(ctx context.Context, req SetStickerSetThumb) (Result, error)
	// SetWebhook implements setWebhook operation.
	//
	// POST /setWebhook
	SetWebhook(ctx context.Context, req SetWebhook) (Result, error)
	// StopMessageLiveLocation implements stopMessageLiveLocation operation.
	//
	// POST /stopMessageLiveLocation
	StopMessageLiveLocation(ctx context.Context, req StopMessageLiveLocation) (Result, error)
	// StopPoll implements stopPoll operation.
	//
	// POST /stopPoll
	StopPoll(ctx context.Context, req StopPoll) (ResultPoll, error)
	// UnbanChatMember implements unbanChatMember operation.
	//
	// POST /unbanChatMember
	UnbanChatMember(ctx context.Context, req UnbanChatMember) (Result, error)
	// UnbanChatSenderChat implements unbanChatSenderChat operation.
	//
	// POST /unbanChatSenderChat
	UnbanChatSenderChat(ctx context.Context, req UnbanChatSenderChat) (Result, error)
	// UnpinAllChatMessages implements unpinAllChatMessages operation.
	//
	// POST /unpinAllChatMessages
	UnpinAllChatMessages(ctx context.Context, req UnpinAllChatMessages) (Result, error)
	// UnpinChatMessage implements unpinChatMessage operation.
	//
	// POST /unpinChatMessage
	UnpinChatMessage(ctx context.Context, req UnpinChatMessage) (Result, error)
	// UploadStickerFile implements uploadStickerFile operation.
	//
	// POST /uploadStickerFile
	UploadStickerFile(ctx context.Context, req UploadStickerFile) (ResultFile, error)
	// NewError creates ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s, err := newConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
