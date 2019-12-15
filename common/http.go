package common

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

const (
	POST = iota
	GET
)

func SendHTTPReuqestWithHeadr(method int, url string, paramsRequest interface{}, header map[string]string) (result []byte, headerCode int, err error) {

	var (
		params      []byte
		bytesBuffer = bytes.NewBuffer(make([]byte, 0))
		resp        = &fasthttp.Response{}
		req         = &fasthttp.Request{}
	)

	switch param := paramsRequest.(type) {

	case *fasthttp.Args:
		param.WriteTo(bytesBuffer)
		params = bytesBuffer.Bytes()
		break
	case string:
		params = []byte(param)
		break
	case map[string]string:
		stringArr := make([]string, 0)
		for key, value := range param {
			stringArr = append(stringArr, key+"="+value)
		}
		params = []byte(strings.Join(stringArr, "&"))
		break
	case []byte:
		params = param
	}

	switch method {

	case POST:
		req.Header.SetMethod("POST")
		fmt.Println(string(params))
		req.BodyWriter().Write(params)

		req.SetRequestURI(url)
	case GET:
		req.Header.SetMethod("GET")
		req.SetRequestURI(url + "?" + string(params))
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	fastClient := fasthttp.Client{
		MaxIdleConnDuration: time.Second * 5,
	}
	if err = fastClient.DoTimeout(req, resp, 5*time.Second); err != nil {
		return
	}
	fmt.Println(req.Header.String())
	fmt.Println(string(req.Body()))
	req.SetConnectionClose()

	return resp.Body(), resp.Header.StatusCode(), nil

}
