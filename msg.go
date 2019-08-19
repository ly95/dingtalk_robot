package robot

// MsgText 文本消息
type MsgText struct {
	Msgtype string         `json:"msgtype"`
	Text    MsgTextPayload `json:"text"`
	At      MsgAt          `json:"at"`
}

// NewMsgText create MsgText instance
func NewMsgText(content string) *MsgText {
	return &MsgText{
		Msgtype: "text",
		Text: MsgTextPayload{
			Content: content,
		},
	}
}

// MsgAt 消息@
type MsgAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// MsgTextPayload 本文消息内容
type MsgTextPayload struct {
	Content string `json:"content"`
}

// MsgMarkdown markdown消息
type MsgMarkdown struct {
	Msgtype  string             `json:"msgtype"`
	Markdown MsgMarkdownPayload `json:"markdown"`
	At       MsgAt              `json:"at"`
}

// MsgMarkdownPayload markdown消息内容
type MsgMarkdownPayload struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// NewMsgMarkdown create MsgMarkdown instance
func NewMsgMarkdown(title, text string) *MsgMarkdown {
	return &MsgMarkdown{
		Msgtype: "markdown",
		Markdown: MsgMarkdownPayload{
			Title: title,
			Text:  text,
		},
	}
}
