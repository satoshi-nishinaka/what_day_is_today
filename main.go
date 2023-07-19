package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/slack-go/slack"
)

func main() {
	log.Println("Start")
	token := os.Getenv("SLACK_BOT_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	if token == "" || channel == "" {
		log.Fatalln("環境変数が取得できませんでした")
		return
	}

	doc, _ := htmlquery.LoadURL("https://kids.yahoo.co.jp/today/")

	titleElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dt/span")
	fmt.Printf("タイトル: %s\n", htmlquery.InnerText(titleElement))

	descriptionElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dd")
	fmt.Printf("本文: %s\n", htmlquery.InnerText(descriptionElement))

	message := "✨✨✨ 今日は何の日？ ✨✨✨\n\n**" + htmlquery.InnerText(titleElement) + "**\n\n" + htmlquery.InnerText(descriptionElement)
	client := slack.New(token)

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := client.PostMessage(channel, slack.MsgOptionText(message, true))
	if err != nil {
		panic(err)
	}

	log.Println("Finish")
}
