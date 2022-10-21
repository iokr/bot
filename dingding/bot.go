package dingding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/imroc/req"
)

const (
	// defaultUrl 默认钉钉url,如果在使用过程中想修改,请使用WithUrl传入新的url
	defaultUrl = "https://oapi.dingtalk.com/robot/send"
)

// Bot 钉钉群聊机器人.
type Bot struct {
	url         string
	accessToken string
	secret      string
	mobiles     []string
	atUserIds   []string
	isAtAll     bool
	message     *Message
}

// NewBot new a ding ding bot.
// 目前只支持发送text,markdown消息类型.
// 若想支持其它消息类型,需要自己添加.
// 	其它消息类型请参考钉钉消息类型格式:
// 	https://developers.dingtalk.com/document/robots/message-types-and-data-format.
func NewBot(accessToken string, opts ...BotOption) *Bot {
	bot := &Bot{
		url:         defaultUrl,
		accessToken: accessToken,
		mobiles:     []string{},
		atUserIds:   []string{},
		isAtAll:     false,
	}

	for _, opt := range opts {
		opt(bot)
	}
	return bot
}

// SendText send text message.
func (b *Bot) SendText(content string) (resp string, err error) {
	b.message = &Message{
		Msgtype: "text",
		Text: &Text{
			Content: fmt.Sprintf("%s%s", b.atMobiles(), content),
		},
		At: b.messageAt(),
	}
	return b.SendMessage()
}

// SendMarkdown send markdown message.
func (b *Bot) SendMarkdown(title, content string) (resp string, err error) {
	b.message = &Message{
		Msgtype: "markdown",
		Markdown: &Markdown{
			Title: title,
			Text:  fmt.Sprintf("%s%s", b.atMobiles(), content),
		},
		At: b.messageAt(),
	}
	return b.SendMessage()
}

// SendMessage 发送message.
// message 不能为空.
func (b *Bot) SendMessage() (resp string, err error) {
	header := req.Header{
		"Accept": "application/json;charset=utf-8",
	}

	queryParam := req.QueryParam{
		"access_token": b.accessToken,
	}

	if b.secret != "" {
		timestamp := time.Now().UnixNano() / 1e6
		queryParam["timestamp"] = timestamp
		queryParam["sign"] = sign(timestamp, b.secret)
	}

	response, err := req.Post(b.url, header, queryParam, req.BodyJSON(b.message))
	if err != nil {
		return "", err
	}
	return response.ToString()
}

func (b *Bot) messageAt() *At {
	return &At{
		AtMobiles: b.mobiles,
		AtUserIds: b.atUserIds,
		IsAtAll:   b.isAtAll,
	}
}

func (b *Bot) atMobiles() string {
	if len(b.mobiles) == 0 {
		return ""
	}

	buf := new(bytes.Buffer)
	for idx := range b.mobiles {
		buf.WriteByte('@')
		buf.WriteString(b.mobiles[idx])
	}
	buf.WriteByte('\n')
	return buf.String()
}

func sign(timestamp int64, secret string) string {
	strToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToSign))
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}
