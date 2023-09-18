package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
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

func buildMessage() string {

	doc, _ := htmlquery.LoadURL("https://kids.yahoo.co.jp/today/")

	titleElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dt/span")
	fmt.Printf("タイトル: %s\n", htmlquery.InnerText(titleElement))

	descriptionElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dd")
	fmt.Printf("本文: %s\n", htmlquery.InnerText(descriptionElement))

	people := parsePeople()
	fmt.Printf("今日が誕生日の有名人: %s\n", people)

	message := fmt.Sprintf("✨✨✨ 今日は何の日？ ✨✨✨\n\n%s\n\n%s\n\n🎂🎂🎂 今日が誕生日の有名人 🎂🎂🎂\n\n%s\n", htmlquery.InnerText(titleElement), htmlquery.InnerText(descriptionElement), people)
	fmt.Println(message)

	return message
}

func parsePeople() string {
	doc, _ := htmlquery.LoadURL("https://kids.yahoo.co.jp/today/")

	personAndYear, _ := htmlquery.QueryAll(doc, "//*[@id=\"birthdayDtl\"]/li")

	var list []string
	for _, n := range personAndYear {
		nameElement := htmlquery.FindOne(n, "b")
		yearElement := htmlquery.FindOne(n, "span")

		text := fmt.Sprintf("%s %s", strings.TrimSpace(htmlquery.InnerText(nameElement)), strings.TrimSpace(htmlquery.InnerText(yearElement)))
		list = append(list, text)
	}

	return strings.Join(list, "\n")
}

func main() {
	lambda.Start(Handler)
}
