package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Error loading .env file")
	}
	data, err := ioutil.ReadFile("test.md")
	if err != nil {
		panic(err)
	}
	//
	bodyData := url.Values{
		"body": {string(data)},
	}
	gitlabUrl := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/issues/%s/notes", os.Getenv("PROJECT"), os.Getenv("ISSUE"))
	req, err := http.NewRequest("POST", gitlabUrl, bytes.NewBufferString(bodyData.Encode()))

	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("TOKEN"))

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
}
