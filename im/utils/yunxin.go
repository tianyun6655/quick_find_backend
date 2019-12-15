package utils

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"quick_find_backend/common"
	"strconv"
	"time"
)

const (
	appSecret   = "ce0a63dd86d2"
	appKey      = "3a2c15109bc3b854708ba18f1d7e39ba"
	sendMessage = "https://api.netease.im/nimserver/msg/sendMsg.action"
)

type YunXinResponse struct {
	Code int `json:"code"`
	Info struct {
		Token string `json:"token"`
	} `json:"info"`
}

func SendMessage(from, to string, content string, types string) (err error) {
	var (
		response []byte
		code     int
		formData = new(fasthttp.Args)
	)
	nonce, _ := common.GenerateUUID()
	currentTime := time.Now().Unix()
	formData.Add("from", from)
	formData.Add("to", to)
	formData.Add("ope", "0")
	formData.Add("type", types)
	formData.Add("body", content)

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "utf-8",
		"AppKey":       appKey,
		"Nonce":        nonce,
		"CurTime":      strconv.Itoa(int(currentTime)),
		"CheckSum":     common.Sha1(appSecret + nonce + strconv.Itoa(int(currentTime))),
	}

	if response, code, err = common.SendHTTPReuqestWithHeadr(common.POST,
		sendMessage,
		formData,
		header); err != nil {
		return err
	} else if code != 200 {
		return errors.New("head code is not 200")
	}
	var yunxinResponse YunXinResponse

	if err = json.Unmarshal(response, &yunxinResponse); err != nil {
		return err
	}
	if yunxinResponse.Code != 200 {
		return errors.New("update yunxin user failed: err: " + string(response))
	}
	return nil
}

