package ai

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
)

const(
	AccessKeyID="LTAI4FnRPBPHw3AGmLH4XAd7"
	AccessKeySecret="TWIsJgeFFnnPJWxdFv5ppgPeQKpsWU"
)

var(
   client *alimt.Client
)

func GetClient()(*alimt.Client,error){
	var err error
	if client==nil{
		client,err= alimt.NewClientWithAccessKey(
			"cn-hangzhou",
			AccessKeyID,AccessKeySecret)
		if err!=nil{
			return nil,err
		}
	}
	return client,nil
}

func Translate(text string)(string, error){
	client,err:=GetClient()
	if err!=nil{
		return "",err
	}
	request := alimt.CreateTranslateGeneralRequest()
	request.Method = "POST"
	request.FormatType = "text"
	request.SourceLanguage = "zh"
	request.SourceText = text
	request.TargetLanguage = "en"
	response, err := client.TranslateGeneral(request)
	if err != nil {
	   return "",err
	}
	fmt.Println(response)

	return response.Data.Translated,nil

}