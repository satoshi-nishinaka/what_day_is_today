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

	titleElement := htmlquery.FindOne(doc, "//*[@id=\"__next\"]/div/main/div[3]/div[1]/div[2]/div/h2")
	fmt.Printf("タイトル: %s\n", htmlquery.InnerText(titleElement))

	descriptionElement := htmlquery.FindOne(doc, "//*[@id=\"__next\"]/div/main/div[3]/div[1]/div[2]/p")
	fmt.Printf("本文: %s\n", htmlquery.InnerText(descriptionElement))

	people := parsePeople()
	fmt.Printf("今日が誕生日の有名人: %s\n", people)

	message := fmt.Sprintf("✨✨✨ 今日は何の日？ ✨✨✨\n\n%s\n\n%s\n\n🎂🎂🎂 今日が誕生日の有名人 🎂🎂🎂\n\n%s\n", htmlquery.InnerText(titleElement), htmlquery.InnerText(descriptionElement), people)
	fmt.Println(message)

	return message
}

func parsePeople() string {
	doc, _ := htmlquery.LoadURL("https://kids.yahoo.co.jp/today/")

	name, _ := htmlquery.QueryAll(doc, "//*[@id=\"mod_birthdays\"]/dl/dt")
	birthdayAndJob, _ := htmlquery.QueryAll(doc, "//*[@id=\"mod_birthdays\"]/dl/dd")

	var list []string
	for i, n := range name {
		nameElement := htmlquery.InnerText(n)
		birthdayElement := htmlquery.InnerText(birthdayAndJob[i*2])
		jobElement := htmlquery.InnerText(birthdayAndJob[i*2+1])
		fmt.Println(birthdayElement, jobElement)

		text := fmt.Sprintf("%s (%s %s)", nameElement, birthdayElement, jobElement)

		list = append(list, text)
	}

	return strings.Join(list, "\n")
}

func main() {
	lambda.Start(Handler)
	// fmt.Println(buildMessage())
}
