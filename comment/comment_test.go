package comment_test

import (
	"github.com/joho/godotenv"
	"gitlab/comment"
	"io/ioutil"
	"log"
	"net/url"
	"testing"
)

func TestGitIssueComment(t *testing.T) {
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
	comment.GitIssueComment(bodyData)
}
