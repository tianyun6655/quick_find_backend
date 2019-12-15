package quick_find_services_im

import (
	"context"
	"errors"
	"quick_find_backend/im/ai"
	"quick_find_backend/im/dao"
)

type IMServices struct {
	dao *dao.Dao
}

func NewImServices(dao *dao.Dao) *IMServices {
	return &IMServices{
		dao: dao,
	}
}

func (im *IMServices) SendMessage(context context.Context, res *SendMessageRequest, resp *SendMessageResponse) error {
	//translate, err := ai.Translate(res.Message)
	//if err != nil {
	//	return errors.New("translate error")
	//}

	//
	//if err := utils.SendMessage(
	//	strconv.Itoa(int(res.From)),
	//	strconv.Itoa(int(res.To)),
	//	translate, "100"); err != nil {
	//		common.Logger.Error("[s.im] SendMessage(%#v) err:",res,err.Error())
	//	return errors.New("send message failed")
	//}

	if err:=im.dao.Insert(res.From,res.To,res.Message);err!=nil{
		return errors.New("send message failed")
	}

	return nil
}

func  (im *IMServices) Translate(context context.Context, res *TranslateRequest, resp *TranslateResponse) error{
	translate, err := ai.Translate(res.Text)
	if err != nil {
		return errors.New("translate error")
	}
	resp.Text=translate
	return nil
}
