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

//客服 图片消息
type CustomMessage struct {
	CommonCustom
	Text  *CustomText  `json:"text,omitempty"`
	Image *CustomMedia `json:"image,omitempty"`
	Voice *CustomMedia `json:"voice,omitempty"`
	Video *CustomVideo `json:"video,omitempty"`
}

//NewCustomText 回复图片消息
func NewCustomText(toUser, content string) *CustomMessage {
	text := new(CustomMessage)
	text.SetMsgType(MsgTypeText)
	text.SetToUser(toUser)
	customText := new(CustomText)
	customText.Content = content
	text.Text = customText
	return text
}

//NewCustomImage 回复图片消息
func NewCustomImage(toUser, mediaID string) *CustomMessage {
	image := new(CustomMessage)
	image.SetMsgType(MsgTypeImage)
	image.SetToUser(toUser)
	customImage := new(CustomMedia)
	customImage.MediaID = mediaID
	image.Image = customImage
	return image
}

//NewCustomVoice 回复声音消息
func NewCustomVoice(toUser, mediaID string) *CustomMessage {
	voice := new(CustomMessage)
	voice.SetMsgType(MsgTypeVoice)
	voice.SetToUser(toUser)
	voice.Voice.MediaID = mediaID

	customVoice := new(CustomMedia)
	customVoice.MediaID = mediaID
	voice.Voice = customVoice
	return voice
}

//NewCustom 回复视频消息
func NewCustomVideo(toUser, mediaID, thumbMediaId, title, description string) *CustomMessage {
	video := new(CustomMessage)
	video.SetMsgType(MsgTypeVideo)
	video.SetToUser(toUser)
	customVideo := new(CustomVideo)
	customVideo.MediaID = mediaID
	customVideo.ThumbMediaId = thumbMediaId
	customVideo.Title = title
	customVideo.Description = description
	video.Video = customVideo
	return video
}
