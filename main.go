package main

import (
	"fmt"
	"os"

	"github.com/parnurzeal/gorequest"
)

func main() {
	//url := "http://requestb.in/vkfyqovk"
	url := "http://www.cafe-amazon.com/api/BranchApi.ashx"
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Type("form").
		Send(`{"command":"get_branchs"}`).
		Send(`{"BranchID":"1"}`).
		Send(`{"languageid":"1"}`).
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	fmt.Println(resp, body)
}
