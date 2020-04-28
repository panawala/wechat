package message

type CustomText struct {
	Content string `json:"content"`
}

type CustomMedia struct {
	MediaID string `json:"media_id"`
}

type CustomVideo struct {
	CustomMedia
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type CustomMiniProgramPage struct {
	Title        string `json:"title"`
	AppId        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

//客服 图片消息
type CustomMessage struct {
	CommonCustom
	Text            *CustomText            `json:"text,omitempty"`
	Image           *CustomMedia           `json:"image,omitempty"`
	Voice           *CustomMedia           `json:"voice,omitempty"`
	Video           *CustomVideo           `json:"video,omitempty"`
	MiniProgramPage *CustomMiniProgramPage `json:"miniprogrampage,omitempty"`
}

//NewCustomText 回复图片消息
func NewCustomText(content string) *CustomMessage {
	text := new(CustomMessage)
	text.SetMsgType(MsgTypeText)
	customText := new(CustomText)
	customText.Content = content
	text.Text = customText
	return text
}

//NewCustomImage 回复图片消息
func NewCustomImage(mediaID string) *CustomMessage {
	image := new(CustomMessage)
	image.SetMsgType(MsgTypeImage)
	customImage := new(CustomMedia)
	customImage.MediaID = mediaID
	image.Image = customImage
	return image
}

//NewCustomVoice 回复声音消息
func NewCustomVoice(mediaID string) *CustomMessage {
	voice := new(CustomMessage)
	voice.SetMsgType(MsgTypeVoice)
	voice.Voice.MediaID = mediaID

	customVoice := new(CustomMedia)
	customVoice.MediaID = mediaID
	voice.Voice = customVoice
	return voice
}

//NewCustom 回复视频消息
func NewCustomVideo(mediaID, thumbMediaId, title, description string) *CustomMessage {
	video := new(CustomMessage)
	video.SetMsgType(MsgTypeVideo)
	customVideo := new(CustomVideo)
	customVideo.MediaID = mediaID
	customVideo.ThumbMediaId = thumbMediaId
	customVideo.Title = title
	customVideo.Description = description
	video.Video = customVideo
	return video
}

//NewCustom 回复视频消息
func NewCustomMiniProgramPage(appId, thumbMediaId, title, pagePath string) *CustomMessage {
	miniProgramPage := new(CustomMessage)
	miniProgramPage.SetMsgType(MsgTypeMiniProgramPage)
	customMiniProgramPage := new(CustomMiniProgramPage)
	customMiniProgramPage.AppId = appId
	customMiniProgramPage.ThumbMediaId = thumbMediaId
	customMiniProgramPage.Title = title
	customMiniProgramPage.PagePath = pagePath
	return miniProgramPage
}
