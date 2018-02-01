package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func main() {
	content := "This is text message"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(content))
	md5Str := md5Ctx.Sum(nil)
	sign := hex.EncodeToString(md5Str)
	fmt.Println(sign)


	token := "Tfd47jzfdtewFFEwf"
	data := map[string]string{
		"user_id": "82692",
		"order_id": "20180201433442",
		"name": "smith",
		"bank_id": "2040",
	}

	text := ""
	keys := make([]string, 0)
	for key, _ := range data{
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, val := range keys {
		text += val + "=" + data[val] + "&"
	}

	text = strings.TrimRight(text, "&")
	text += token
	fmt.Println(text)
	//must recreate a md5 pointer
	md5Ctx = md5.New()
	md5Ctx.Write([]byte(text))
	md5Str = md5Ctx.Sum(nil)
	sign = hex.EncodeToString(md5Str)
	fmt.Println(sign)
}
