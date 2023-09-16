package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func Handler() {
	log.Println("Start")
	token := os.Getenv("SLACK_BOT_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	if token == "" || channel == "" {
		log.Fatalln("環境変数が取得できませんでした")
		return
	}

	message := buildMessage()
	client := slack.New(token)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := client.PostMessage(channel, slack.MsgOptionText(message, true))
	if err != nil {
		panic(err)
	}

	log.Println("Finish")
}

func main() {
	lambda.Start(Handler)
}
