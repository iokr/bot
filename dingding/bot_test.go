package dingding

import (
	"testing"
)

const (
	// TestAccessToken 创建钉钉机器人时，Webhook后面access_token的值.
	// https://oapi.dingtalk.com/robot/send?access_token=xxxx.
	TestAccessToken = "access token value xxxxxx"

	// TestSecret 创建钉钉机器人时，安全设置中加签密钥.
	TestSecret = "secret value xxxxxx"
)

func TestNewBotByKeyWord(t *testing.T) {
	bot := NewBot(TestAccessToken)

	_, err := bot.SendText("这是文本内容")
	if err != nil {
		t.Fatal(err)
	}

	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	_, err = bot.SendMarkdown("标题", content)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewBotBySecret(t *testing.T) {
	bot := NewBot(TestAccessToken, WithSecret(TestSecret))

	_, err := bot.SendText("这是文本内容")
	if err != nil {
		t.Fatal(err)
	}

	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	_, err = bot.SendMarkdown("标题", content)
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendTextKeyWork 测试关键词.
func TestSendTextKeyWork(t *testing.T) {
	err := SendText(TestAccessToken, "测试: 这是文本内容")
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendTextKeyWorkAndMobiles 测试关键词@某人.
func TestSendTextKeyWorkAndMobiles(t *testing.T) {
	err := SendText(TestAccessToken, "测试: 这是文本内容", WithMobiles([]string{"10010", "10086"}))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendTextKeyWork 测试关键词@所有人.
func TestSendTextKeyWorkAndAtAll(t *testing.T) {
	err := SendText(TestAccessToken, "测试: 这是文本内容", WithAtAll(true))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendTestBySecret 测试使用秘钥.
func TestSendTestBySecret(t *testing.T) {
	err := SendText(TestAccessToken, "这是文本内容", WithSecret(TestSecret))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendTestBySecretAndMobiles(t *testing.T) {
	err := SendText(TestAccessToken, "这是文本内容",
		WithSecret(TestSecret),
		WithMobiles([]string{"10010", "10086"}))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendTestBySecretAndAtAll 测试秘钥@所有人.
func TestSendTestBySecretAndAtAll(t *testing.T) {
	err := SendText(TestAccessToken, "这是文本内容",
		WithSecret(TestSecret),
		WithAtAll(true))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMarkdown 测试同text消息.
func TestSendMarkdown(t *testing.T) {
	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	err := SendMarkdown(TestAccessToken, "标题", content)
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMarkdownBySecret 测试同text消息.
func TestSendMarkdownBySecret(t *testing.T) {
	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	err := SendMarkdown(TestAccessToken, "杭州天气标题", content, WithSecret(TestSecret))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMarkdownBySecretAndMobiles 测试秘钥@手机号.
func TestSendMarkdownBySecretAndMobiles(t *testing.T) {
	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	err := SendMarkdown(TestAccessToken, "杭州天气标题", content,
		WithSecret(TestSecret),
		WithMobiles([]string{"10010", "10086"}))
	if err != nil {
		t.Fatal(err)
	}
}

// TestSendMarkdownBySecretAndAtAll 测试秘钥@所有人.
func TestSendMarkdownBySecretAndAtAll(t *testing.T) {
	content := "#### 杭州天气 \n" +
		"> 9度，西北风1级，空气良89，相对温度73% \n" +
		"> ![screenshot](https://t7.baidu.com/it/u=2604797219,1573897854&fm=193&f=GIF)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingalk.com) \n"
	err := SendMarkdown(TestAccessToken, "杭州天气标题", content,
		WithSecret(TestSecret),
		WithAtAll(true))
	if err != nil {
		t.Fatal(err)
	}
}
