package main

import (
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
)

func main() {
	log.Println("Start")
	doc, _ := htmlquery.LoadURL("https://kids.yahoo.co.jp/today/")

	titleElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dt/span")
	fmt.Printf("タイトル: %s\n", htmlquery.InnerText(titleElement))

	descriptionElement := htmlquery.FindOne(doc, "//*[@id=\"dateDtl\"]/dd")
	fmt.Printf("本文: %s\n", htmlquery.InnerText(descriptionElement))
}
