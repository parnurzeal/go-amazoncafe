package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

type ProvinceResponse struct {
	Message string
	Meta    string
	Result  []Province
}
type Province struct {
	ProvinceLanguageTitle string
	LanguageID            int
	CountBranchs          int
	ZoneLanguageTitle     string
	ProvinceID            int
	ZoneID                int
	IsActivated           bool
	IsDeleted             bool
}

type BranchResponse struct {
	Message string
	Meta    string
	Result  []Branch
}

type Branch struct {
	BranchLanguageTitle   string
	BranchLanguageAddress string
	ZoneTitle             string
	ProvinceTitle         string
	DealerFullName        string
	Telephone             string `json:"Telephone"`
	Fax                   string
}

func getZone() {
	url := "http://www.cafe-amazon.com/api/ZoneApi.ashx"
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Type("form").
		Send(`{"command":"get_zones"}`).
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

func getProvince(zoneId int) []Province {
	url := "http://www.cafe-amazon.com/api/provinceapi.ashx"
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Type("form").
		Send(`{"command":"get_provinces"}`).
		Send(`{"zoneid":"` + strconv.Itoa(zoneId) + `"}`).
		Send(`{"languageid":"1"}`).
		Send(`{"orderby":"provincelanguagetitle asc"}`).
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	allProvinces := []Province{}
	response := &ProvinceResponse{}
	json.Unmarshal([]byte(body), &response)
	for _, province := range response.Result {
		allProvinces = append(allProvinces, province)
	}
	return allProvinces
}

func getBranchById(branchId int) []Branch {
	url := "http://www.cafe-amazon.com/api/BranchApi.ashx"
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Type("form").
		Send(`{"command":"get_branchs"}`).
		Send(`{"BranchID":"` + strconv.Itoa(branchId) + `"}`).
		Send(`{"languageid":"1"}`).
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	allBranchs := []Branch{}
	json.Unmarshal([]byte(body), allBranchs)
	return allBranchs
}

func getBranchByProvinceId(provinceId int) []Branch {
	url := "http://www.cafe-amazon.com/api/BranchApi.ashx"
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Type("form").
		Send(`{"command":"get_branchs"}`).
		Send(`{"ProvinceID":"` + strconv.Itoa(provinceId) + `"}`).
		Send(`{"languageid":"1"}`).
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	allBranchs := []Branch{}
	response := &BranchResponse{}
	json.Unmarshal([]byte(body), &response)
	for _, branch := range response.Result {
		allBranchs = append(allBranchs, branch)
	}
	return allBranchs
}

func main() {
	provinceList := getProvince(1)
	for _, province := range provinceList {
		branchList := getBranchByProvinceId(province.ProvinceID)
		fmt.Println(province.ProvinceLanguageTitle, len(branchList))
		for _, branch := range branchList {
			fmt.Println(branch.BranchLanguageTitle, "//", branch.BranchLanguageAddress, "//", branch.DealerFullName, "//", branch.Telephone)
		}
	}
}
