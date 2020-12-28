package main

import (
	"net/http"
	"strings"
)

var (
	HuaweiRequestBody = `{
		"template": "healthyon",
		"w3Account": "lWX748836",
		"fillingMethod": "pc",
		"beginTime": "2020-02-16 10:34:39",
		"endTime": "2020-02-16 10:35:18",
		"t1": "否",
		"t2": "否",
		"t3": "否",
		"t4": "否",
		"t5": "否",
		"t6": "",
		"t7": "否",
		"t8": "否",
		"t11": "内蒙古自治区/Nei Mongol",
		"t12": "乌兰察布市/Ulanqab",
		"t13": "",
		"t14": "",
		"t15": "",
		"clientId": "0.0.0.0"
		}`
	HuaweiUriForHealth    = "https://app.huawei.com/engine/epidemicSurvey/api/survey"
	SoftStoneUriForHealth = "http://i.isoftstone.com/cht/cht.html#/"
)

func main() {

}

func HitCardForHealth() error {
	// req, err := http.NewRequest("POST", HuaweiUriForHealth, strings.NewReader(HuaweiRequestBody))
	resp, err := http.Post(HuaweiUriForHealth, "", strings.NewReader(HuaweiRequestBody))
	if err != nil {
		return err
	}
	return nil
}
