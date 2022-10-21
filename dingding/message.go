package dingding

// Message 钉钉消息.
// 目前只支持发送text,markdown消息类型.
// 若想支持其它消息类型,需要自己添加.
// 	其它消息类型请参考钉钉消息类型格式:
// 	https://developers.dingtalk.com/document/robots/message-types-and-data-format.
type Message struct {
	Msgtype  string    `json:"msgtype"`
	Text     *Text     `json:"text"`
	Markdown *Markdown `json:"markdown"`
	At       *At       `json:"at"`
}

// Text 文本.
type Text struct {
	Content string `json:"content"`
}

// Markdown markdown.
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// At 群消息是否@其他人.
type At struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}
