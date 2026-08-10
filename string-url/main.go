package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	user1, err := ReturnUser("https://spb.hh.ru/applicant/profile/me?user=user&date=123&time=333")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)

	user2, err := ReturnUser("https://spb.hh.ru/applicant/profile/me?user=yuriy&date=123&time=333")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user2)
	user3, err := ReturnUser("https://spb.hh.ru/applicant/profile/me")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user3)
	user4, err := ReturnUser("https://spb.hh.ru/applicant/profile/me?date=123&time=333")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user4)
}

func ReturnUser(url string) (string, error) {
	urlQuestionIndex := strings.Index(url, "?")
	if urlQuestionIndex == -1 {
		return "", errors.New("no ? in url")
	}
	partAfterQuestion := url[urlQuestionIndex+1:]

	urlParts := strings.Split(partAfterQuestion, "&")

	for i := 0; i < len(urlParts); i++ {
		equalSignIndex := strings.Index(urlParts[1], "=")
		if equalSignIndex != -1 {
			strBefore := urlParts[i][:equalSignIndex]
			strAfter := urlParts[i][equalSignIndex+1:]

			if strBefore == "user" {
				return strAfter, nil
			}
		}
	}

	return "", errors.New("no user")
}
