package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"quick_find_backend/common"
	"strconv"
	"strings"
	"time"
)

const(
   appSecret="ce0a63dd86d2"
   appKey="3a2c15109bc3b854708ba18f1d7e39ba"
   registerUrl = "https://api.netease.im/nimserver/user/create.action"

)
type YunXinResponse struct {
	Code int `json:"code"`
	Info struct {
		Token string `json:"token"`
	} `json:"info"`
}

type ApplyYunchatModel struct {
	MessageType int    `json:"messageType"`
	Content     string `json:"content"`
	ManagerType int    `json:"manager_type"`
	ApplyYunchatData struct {
		Apid     int    `json:"apid"`
		Nickname string `json:"nickname"`
	} `json:"business_data"`
}

func Resiter(yunxinId int)(string,error){
	var (
		response        []byte
		header          map[string]string
		currentTime     = time.Now().Unix()
		formData        []byte
		yunxingResponce YunXinResponse
	)
	nonce, err := common.GenerateUUID()
	if err!=nil{
		return "",err
	}
	formData = []byte("accid=" + strconv.Itoa(yunxinId))
	header = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "utf-8",
		"AppKey":       appKey,
		"Nonce":        nonce,
		"CurTime":      strconv.Itoa(int(currentTime)),
		"CheckSum":     common.Sha1(appSecret + nonce + strconv.Itoa(int(currentTime))),
	}
	if response, err = sendOriginHttpPost(registerUrl, formData, header); err != nil {
		return "",err
	}

	if err = json.Unmarshal(response, &yunxingResponce); err != nil {
		return "",err
	}

	if yunxingResponce.Code != 200 {
		return "", errors.New("header is not 200")
	}
	return yunxingResponce.Info.Token, nil
}

func sendOriginHttpPost(url string, params []byte, heaer map[string]string) (result []byte, err error) {
	client := http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(params)))
	if err != nil {
		return nil, err
	}

	for key, val := range heaer {
		req.Header.Set(key, val)
	}

	resp, err := client.Do(req)

	if err != nil || resp == nil {
		return nil, err
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}