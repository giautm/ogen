// Code generated by ogen, DO NOT EDIT.

package api

// APICaptcha2chcaptchaShowGetNotFound is response for APICaptcha2chcaptchaShowGet operation.
type APICaptcha2chcaptchaShowGetNotFound struct{}

func (*APICaptcha2chcaptchaShowGetNotFound) aPICaptcha2chcaptchaShowGetRes() {}

// APICaptcha2chcaptchaShowGetOK is response for APICaptcha2chcaptchaShowGet operation.
type APICaptcha2chcaptchaShowGetOK struct{}

func (*APICaptcha2chcaptchaShowGetOK) aPICaptcha2chcaptchaShowGetRes() {}

// APICaptchaInvisibleRecaptchaMobileGetOK is response for APICaptchaInvisibleRecaptchaMobileGet operation.
type APICaptchaInvisibleRecaptchaMobileGetOK struct{}

// APICaptchaRecaptchaMobileGetOK is response for APICaptchaRecaptchaMobileGet operation.
type APICaptchaRecaptchaMobileGetOK struct{}

// Модель доски.
// Ref: #/components/schemas/Board
type Board struct {
	BumpLimit        int      "json:\"bump_limit\""
	Category         string   "json:\"category\""
	DefaultName      string   "json:\"default_name\""
	EnableDices      bool     "json:\"enable_dices\""
	EnableFlags      bool     "json:\"enable_flags\""
	EnableIcons      bool     "json:\"enable_icons\""
	EnableLikes      bool     "json:\"enable_likes\""
	EnableNames      bool     "json:\"enable_names\""
	EnableOekaki     bool     "json:\"enable_oekaki\""
	EnablePosting    bool     "json:\"enable_posting\""
	EnableSage       bool     "json:\"enable_sage\""
	EnableShield     bool     "json:\"enable_shield\""
	EnableSubject    bool     "json:\"enable_subject\""
	EnableThreadTags bool     "json:\"enable_thread_tags\""
	EnableTrips      bool     "json:\"enable_trips\""
	FileTypes        []string "json:\"file_types\""
	// Массив иконок, которые доступны на этой доске, если
	// они включены.
	Icons []BoardIconsItem "json:\"icons\""
	ID    string           "json:\"id\""
	// Информация о доске.
	Info string "json:\"info\""
	// Информация о доске для главной.
	InfoOuter    string "json:\"info_outer\""
	MaxComment   int    "json:\"max_comment\""
	MaxFilesSize int    "json:\"max_files_size\""
	MaxPages     int    "json:\"max_pages\""
	Name         string "json:\"name\""
	// Массив тегов, которые доступны на этой доске, если они
	// включены.
	Tags           []string "json:\"tags\""
	ThreadsPerPage int      "json:\"threads_per_page\""
}

type BoardIconsItem struct {
	Name OptString "json:\"name\""
	Num  OptInt    "json:\"num\""
	URL  OptString "json:\"url\""
}

type Boards []Board

// Ref: #/components/schemas/Captcha
type Captcha struct {
	Error OptError "json:\"error\""
	// Время в секундах после которого id перестанет
	// действовать.
	Expires OptInt "json:\"expires\""
	ID      string "json:\"id\""
	// Тип текста, изображённого на картинке капчи.
	// Возможные варианты:
	// * numeric - только цифры. (0123456789)
	// * english - цифры и английские буквы. (0123456789abcdefghijklmnopqrstuvwxyz)
	// * russian - цифры и русские буквы.
	// (0123456789абвгдеёжзийклмнопрстуфхцчшщъыьэюя)
	// * all - цифры, русские и английские буквы.
	// (0123456789abcdefghijklmnopqrstuvwxyzабвгдеёжзийклмнопрстуфхцчшщъыьэюя).
	Input  OptString   "json:\"input\""
	Result int         "json:\"result\""
	Type   CaptchaType "json:\"type\""
}

// Каждый тип капчи так же требует дополнительные
// параметры для её валидации:
// * recaptcha: g-recaptcha-response
// * invisible_recaptcha: g-recaptcha-response
// * recaptcha3: g-recaptcha-response
// * 2chcaptcha: Два.ч капча
// 2chcaptcha_id - идентификатор Два.ч капчи.
// 2chcaptcha_value - строка, которую пользователь увидел на
// картинке.
// * appid: app_response_id и app_response
// app_response_id - результат запроса к этому методу с
// публичным ключём приложения: /api/captcha/app/id/{public_key}
// app_response - sha256(app_response_id + '|' + private_key)
// * passcode: cookie passcode_auth
// * nocaptcha: капча не требуется, никакие дополнительные
// параметры тоже.
// Ref: #/components/schemas/CaptchaType
type CaptchaType string

const (
	CaptchaTypeRecaptcha          CaptchaType = "recaptcha"
	CaptchaTypeInvisibleRecaptcha CaptchaType = "invisible_recaptcha"
	CaptchaTypeRecaptcha3         CaptchaType = "recaptcha3"
	CaptchaType2chcaptcha         CaptchaType = "2chcaptcha"
	CaptchaTypeAppid              CaptchaType = "appid"
	CaptchaTypePasscode           CaptchaType = "passcode"
	CaptchaTypeNocaptcha          CaptchaType = "nocaptcha"
)

// Ошибка запроса.
// Ref: #/components/schemas/Error
type Error struct {
	Code OptErrorCode "json:\"code\""
	// Описание ошибки на русском языке.
	Message OptString "json:\"message\""
}

// * 0 NoError, ошибки нет.
// * 403 ErrorForbidden, ошибка доступа.
// * 666 ErrorInternal, внутренняя ошибка.
// * 667 ErrorNotFound, используется для совместимости, если
// запрос не существует.
// * -2 ErrorNoBoard, доска не существует.
// * -3 ErrorNoParent, тред не существует.
// * -31 ErrorNoPost, пост не существует.
// * -4 ErrorNoAccess, контент существует, но у вас нет доступа.
// * -41 ErrorBoardClosed, доска закрыта.
// * -42 ErrorBoardOnlyVIP, доступ к доске возможен только с
// пасскодом.
// * -5 ErrorCaptchaNotValid, капча не валидна.
// * -6 ErrorBanned, вы были забанены. Сообщение содержит
// причину и номер бана.
// * -7 ErrorThreadClosed, тред закрыт.
// * -8 ErrorPostingToFast, вы постите слишком быстро ИЛИ
// установлен лимит на создание тредов на доске.
// * -9 ErrorFieldTooBig, поле слишком большое. Например,
// комментарий превысил лимит.
// * -10 ErrorFileSimilar, похожий файл уже был загружен.
// * -11 ErrorFileNotSupported, файл не поддерживается.
// * -12 ErrorFileTooBig, слишком большой файл.
// * -13 ErrorFilesTooMuch, вы загрузили больше файлов, чем
// разрешено на доске.
// * -14 ErrorTripBanned, трипкод был забанен.
// * -15 ErrorWordBanned, в комментарии недопустимое выражение.
// * -16 ErrorSpamList, в комментарии выражение из спамлиста.
// * -19 ErrorEmptyOp, при создании треда необходимо загрузить
// файл.
// * -20 ErrorEmptyPost, пост не может быть пустым, необходим
// комментарий/файл/etc.
// * -21 ErrorPasscodeNotExist, пасскод не существует.
// * -22 ErrorLimitReached, достигнут лимит запросов, попробуйте
// позже.
// * -23 ErrorFieldTooSmall, слишком короткое сообщение.
// (используется в поиске).
// * -50 ErrorReportTooManyPostsm, слишком много постов для жалобы.
// * -51 ErrorReportEmpty, вы ничего не написали в жалобе.
// * -52 ErrorReportExist, вы уже отправляли жалобу.
// * -300 ErrorAppNotExist, приложение не существует или было
// отключено.
// * -301 ErrorAppIDWrong, некорректный идентификатор приложения.
// * -302 ErrorAppIDExpired, идентификатор приложения истёк.
// * -303 ErrorAppIDSignature, неверная подпись поста с помощью
// идентификатора.
// * -304 ErrorAppIDUsed, указанный идентификатор уже был
// использован.
// * -24 ErrorWrongStickerID, некорректный идентификатор стикера.
// * -25 ErrorStickerNotFound, стикер не найден.
// Ref: #/components/schemas/ErrorCode
type ErrorCode int

const (
	ErrorCode0        ErrorCode = 0
	ErrorCode403      ErrorCode = 403
	ErrorCode666      ErrorCode = 666
	ErrorCode667      ErrorCode = 667
	ErrorCodeMinus2   ErrorCode = -2
	ErrorCodeMinus3   ErrorCode = -3
	ErrorCodeMinus31  ErrorCode = -31
	ErrorCodeMinus4   ErrorCode = -4
	ErrorCodeMinus41  ErrorCode = -41
	ErrorCodeMinus42  ErrorCode = -42
	ErrorCodeMinus5   ErrorCode = -5
	ErrorCodeMinus6   ErrorCode = -6
	ErrorCodeMinus7   ErrorCode = -7
	ErrorCodeMinus8   ErrorCode = -8
	ErrorCodeMinus9   ErrorCode = -9
	ErrorCodeMinus10  ErrorCode = -10
	ErrorCodeMinus11  ErrorCode = -11
	ErrorCodeMinus12  ErrorCode = -12
	ErrorCodeMinus13  ErrorCode = -13
	ErrorCodeMinus14  ErrorCode = -14
	ErrorCodeMinus15  ErrorCode = -15
	ErrorCodeMinus16  ErrorCode = -16
	ErrorCodeMinus19  ErrorCode = -19
	ErrorCodeMinus20  ErrorCode = -20
	ErrorCodeMinus21  ErrorCode = -21
	ErrorCodeMinus22  ErrorCode = -22
	ErrorCodeMinus23  ErrorCode = -23
	ErrorCodeMinus300 ErrorCode = -300
	ErrorCodeMinus301 ErrorCode = -301
	ErrorCodeMinus302 ErrorCode = -302
	ErrorCodeMinus303 ErrorCode = -303
	ErrorCodeMinus304 ErrorCode = -304
	ErrorCodeMinus24  ErrorCode = -24
	ErrorCodeMinus25  ErrorCode = -25
)

// Модель файла.
// Ref: #/components/schemas/File
type File struct {
	Displayname string "json:\"displayname\""
	// В случае видео/аудио файла, содержит
	// продолжительность в формате XX:XX:XX.
	Duration OptString "json:\"duration\""
	// В случае видео/аудио файла, содержит
	// продолжительность в секундах.
	DurationSecs OptInt "json:\"duration_secs\""
	Fullname     string "json:\"fullname\""
	Height       int    "json:\"height\""
	// В случае стикера, содержит ссылку на установку.
	Install OptString "json:\"install\""
	MD5     OptString "json:\"md5\""
	Name    string    "json:\"name\""
	// Если >= 0, файл содержит NSFW контент, в данный момент
	// реализовано не на всех досках.
	Nsfw OptInt "json:\"nsfw\""
	// В случае стикера, содержит ID стикер пака.
	Pack OptString "json:\"pack\""
	Path string    "json:\"path\""
	// Размер файла, в КБ.
	Size int "json:\"size\""
	// В случае стикера, содержит ID стикера.
	Sticker   OptString "json:\"sticker\""
	Thumbnail string    "json:\"thumbnail\""
	TnHeight  int       "json:\"tn_height\""
	TnWidth   int       "json:\"tn_width\""
	Type      FileType  "json:\"type\""
	Width     int       "json:\"width\""
}

// * 0 FileTypeNone
// * 1 FileTypeJpg
// * 2 FileTypePng
// * 3 FileTypeAPng
// * 4 FileTypeGif
// * 5 FileTypeBmp
// * 6 FileTypeWebm
// * 7 FileTypeMp3, не используется в данный момент.
// * 8 FileTypeOgg, не используется в данный момент.
// * 10 FileTypeMp4
// * 100 FileTypeSticker.
// Ref: #/components/schemas/FileType
type FileType int

const (
	FileType0   FileType = 0
	FileType1   FileType = 1
	FileType2   FileType = 2
	FileType3   FileType = 3
	FileType4   FileType = 4
	FileType5   FileType = 5
	FileType6   FileType = 6
	FileType7   FileType = 7
	FileType8   FileType = 8
	FileType10  FileType = 10
	FileType100 FileType = 100
)

// Ref: #/components/schemas/Like
type Like struct {
	Error  OptError "json:\"error\""
	Result OptInt   "json:\"result\""
}

// Ref: #/components/schemas/MobilePost
type MobilePost struct {
	Error  OptError "json:\"error\""
	Post   OptPost  "json:\"post\""
	Result OptInt   "json:\"result\""
}

// Ref: #/components/schemas/MobileThreadLastInfo
type MobileThreadLastInfo struct {
	Error  OptError                      "json:\"error\""
	Result OptInt                        "json:\"result\""
	Thread OptMobileThreadLastInfoThread "json:\"thread\""
}

type MobileThreadLastInfoThread struct {
	Num       OptInt "json:\"num\""
	Posts     OptInt "json:\"posts\""
	Timestamp OptInt "json:\"timestamp\""
}

// Ref: #/components/schemas/MobileThreadPostsAfter
type MobileThreadPostsAfter struct {
	Error         OptError "json:\"error\""
	Posts         []Post   "json:\"posts\""
	Result        OptInt   "json:\"result\""
	UniquePosters OptInt   "json:\"unique_posters\""
}

// NewOptError returns new OptError with value set to v.
func NewOptError(v Error) OptError {
	return OptError{
		Value: v,
		Set:   true,
	}
}

// OptError is optional Error.
type OptError struct {
	Value Error
	Set   bool
}

// IsSet returns true if OptError was set.
func (o OptError) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptError) Reset() {
	var v Error
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptError) SetTo(v Error) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptError) Get() (v Error, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptError) Or(d Error) Error {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptErrorCode returns new OptErrorCode with value set to v.
func NewOptErrorCode(v ErrorCode) OptErrorCode {
	return OptErrorCode{
		Value: v,
		Set:   true,
	}
}

// OptErrorCode is optional ErrorCode.
type OptErrorCode struct {
	Value ErrorCode
	Set   bool
}

// IsSet returns true if OptErrorCode was set.
func (o OptErrorCode) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptErrorCode) Reset() {
	var v ErrorCode
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptErrorCode) SetTo(v ErrorCode) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptErrorCode) Get() (v ErrorCode, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptErrorCode) Or(d ErrorCode) ErrorCode {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMobileThreadLastInfoThread returns new OptMobileThreadLastInfoThread with value set to v.
func NewOptMobileThreadLastInfoThread(v MobileThreadLastInfoThread) OptMobileThreadLastInfoThread {
	return OptMobileThreadLastInfoThread{
		Value: v,
		Set:   true,
	}
}

// OptMobileThreadLastInfoThread is optional MobileThreadLastInfoThread.
type OptMobileThreadLastInfoThread struct {
	Value MobileThreadLastInfoThread
	Set   bool
}

// IsSet returns true if OptMobileThreadLastInfoThread was set.
func (o OptMobileThreadLastInfoThread) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMobileThreadLastInfoThread) Reset() {
	var v MobileThreadLastInfoThread
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMobileThreadLastInfoThread) SetTo(v MobileThreadLastInfoThread) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMobileThreadLastInfoThread) Get() (v MobileThreadLastInfoThread, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMobileThreadLastInfoThread) Or(d MobileThreadLastInfoThread) MobileThreadLastInfoThread {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptPost returns new OptPost with value set to v.
func NewOptPost(v Post) OptPost {
	return OptPost{
		Value: v,
		Set:   true,
	}
}

// OptPost is optional Post.
type OptPost struct {
	Value Post
	Set   bool
}

// IsSet returns true if OptPost was set.
func (o OptPost) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptPost) Reset() {
	var v Post
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptPost) SetTo(v Post) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptPost) Get() (v Post, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptPost) Or(d Post) Post {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Модель поста.
// Ref: #/components/schemas/Post
type Post struct {
	Banned    int       "json:\"banned\""
	Board     string    "json:\"board\""
	Closed    int       "json:\"closed\""
	Comment   string    "json:\"comment\""
	Date      string    "json:\"date\""
	Dislikes  OptInt    "json:\"dislikes\""
	Email     OptString "json:\"email\""
	Endless   int       "json:\"endless\""
	Files     []File    "json:\"files\""
	Icon      OptString "json:\"icon\""
	Lasthit   int       "json:\"lasthit\""
	Likes     OptInt    "json:\"likes\""
	Name      OptString "json:\"name\""
	Num       int       "json:\"num\""
	Op        int       "json:\"op\""
	Parent    int       "json:\"parent\""
	Sticky    int       "json:\"sticky\""
	Subject   OptString "json:\"subject\""
	Tags      OptString "json:\"tags\""
	Timestamp int       "json:\"timestamp\""
	Trip      OptString "json:\"trip\""
	TripStyle OptString "json:\"trip_style\""
	Views     int       "json:\"views\""
}
