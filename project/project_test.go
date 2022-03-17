package project_test

import (
	"github.com/joho/godotenv"
	"gitlab/project"
	"testing"
)

func TestGetProjectList(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	project.GetProjectList()
}
