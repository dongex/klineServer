package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetToken() {
	fmt.Println("进来了")
	//token := models.Token
	url := Url
	/*p := map[string]string{
		"userName": "testadmin",
		"passWord": "dex20220505",
	}*/
	p := make(map[string]string)
	p["userName"] = "testadmin"
	p["passWord"] = "dex20220505"

	paramsJson, err := json.Marshal(p)
	if err != nil {
		return
	}
	paramsByte := bytes.NewReader(paramsJson)
	resq, err := http.NewRequest("POST", url, paramsByte)
	if err != nil {
		return
	}
	resq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	res := map[string]interface{}{}
	err1 := json.Unmarshal(body, &res)
	if err1 != nil {
		return
	}
	fmt.Println("aaaaaaaaaaaaa===。", res)
}
