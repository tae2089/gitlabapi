package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetProjectList() {
	url := "https://gitlab.com/api/v4/projects"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("TOKEN"))
	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 결과 출력
	body, _ := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Panic("JSON parse error: ", error)
	}
	log.Println("CSP Violation:", prettyJSON.String())
}
