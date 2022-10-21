package dingding

// BotOption bot option.
type BotOption func(*Bot)

// WithUrl set url.
func WithUrl(url string) BotOption {
	return func(bot *Bot) {
		bot.url = url
	}
}

// WithSecret set secret.
func WithSecret(secret string) BotOption {
	return func(bot *Bot) {
		bot.secret = secret
	}
}

// WithMobiles set mobiles.
// default is empty.
func WithMobiles(mobiles []string) BotOption {
	return func(bot *Bot) {
		bot.mobiles = mobiles
	}
}

// WithAtAll set at all.
// default is false.
func WithAtAll(isAtAll bool) BotOption {
	return func(bot *Bot) {
		bot.isAtAll = isAtAll
	}
}

// WithAtUserIds set at user ids.
// default is empty.
func WithAtUserIds(atUserIds []string) BotOption {
	return func(bot *Bot) {
		bot.atUserIds = atUserIds
	}
}

// SendText send text to ding ding bot.
func SendText(accessToken, content string, opts ...BotOption) error {
	bot := NewBot(accessToken, opts...)
	_, err := bot.SendText(content)
	return err
}

// SendMarkdown send markdown to ding ding bot.
func SendMarkdown(accessToken, title, content string, opts ...BotOption) error {
	bot := NewBot(accessToken, opts...)
	_, err := bot.SendMarkdown(title, content)
	return err
}
