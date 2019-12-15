package dao

import (
	"quick_find_backend/common"
	"quick_find_backend/im/model"
)

func(dao *Dao)Insert(from, to int32,message string)error{
	 chat:=&model.Chat{
	 	From:from,
	 	To:to,
	 	Message:message,
	 }
	 _,err:=dao.sql.Insert(chat)
	 if err!=nil{
	 	common.Logger.Error("[d.im] Insert(%d,%d,%s) err: ",from,to,message)
	 	return err
	 }
	 return nil
}