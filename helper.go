package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func addStandardHeaders(header Header) http.Header {
	var x = http.Header{
		"accept":           {"*/*"},
		"accept-encoding":  {"gzip, deflate, br"},
		"accept-language":  {"en-GB,en-US;q=0.9,en;q=0.8"},
		"content-type":     {"application/x-www-form-urlencoded; charset=UTF-8"},
		"origin":           {"https://www.hottopic.com"},
		"referer":          {"https://www.hottopic.com/product/black-zipper-jogger-pants/13249988.html"},
		"sec-ch-ua":        {"\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\""},
		"sec-ch-ua-mobile": {"?0"},
		"sec-fetch-dest":   {"empty"},
		"sec-fetch-mode":   {"cors"},
		"user-agent":       {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36"},
		"x-requested-with": {"XMLHttpRequest"},

		"Header-Order:": {"accept", "accept-encoding", "accept-language", "content-length", "content-type", "origin", "referer", "sec-ch-ua", "sec-ch-ua-mobile", "sec-fetch-dest", "sec-fetch-mode", "user-agent", "x-requested-with"},
	}
	if header.content != nil {
		x.Set("content-length", fmt.Sprint(header.content.Size()))
	}
	if len(header.cookie) > 0 {
		buildString := ""
		for i := 0; i < len(header.cookie); i++ {
			buildString += header.cookie[i] + "; "
		}
		x.Set("Cookie", buildString+strings.Join(x.Values("Cookie"), "; "))
	}

	return x
}
func getDwCont(resp string) string {
	return getStringInBetweenTwoString(resp, "cart?dwcont=", "\" method")
}
func getStringInBetweenTwoString(str string, startS string, endS string) (result string) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
