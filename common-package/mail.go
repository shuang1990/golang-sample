package main

import (
	"net/smtp"
	"log"
	"fmt"
	"strings"
)

const (
	HOST string = "smtp.qq.com"
	PORT int = 25
	USERNAME string = "mail@admin.com"
	PASSWORD string = "fdsd"
	FROM string = "mail@admin.com"
)

func SendEmail(subject, content, mailType string, to []string)  {
	// Set up authentication information.
	auth := smtp.PlainAuth("", USERNAME, PASSWORD, HOST)

	if mailType != "html" {
		mailType = "plain"
	}
	msg := fmt.Sprintf("To: %s\r\n" +
		"From: 九斗鱼<%s>\r\n" +
		"Subject: %s\r\n" +
		"Content-Type: text/%s; charset=UTF-8\r\n\r\n" +
		"%s\r\n", strings.Join(to, ";"), FROM, subject, mailType, content)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", HOST, PORT), auth, FROM, to, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	subject := "这是一封测试邮件"
	content := "errcode: 300001, errmsg: token is not exist"
	mailType := "text"

	to := []string{
		"zha.shng@admin.com",
		"shua94@163.com",
	}

	SendEmail(subject, content, mailType, to)
}
