package dao

import (
	"errors"
	"quick_find/common"
	"time"
)

const (
	//Golang 常量
	PhoneCode = "phone:code:"
)


func (d *Dao) SetPhoneCode(phone string, code int) (err error) {
	err = d.redis.Set(cominKey(PhoneCode, phone), code, time.Second*60).Err()
	if err != nil {
		common.Logger.Error("d.SetPhoneCode(%v,%v) error(%v)", phone, code, err)
	}
	return
}



func (d *Dao) CheckPhoneCode(phone string, code int) (error) {
	orgCode, err := d.redis.Get(cominKey(PhoneCode, phone)).Int()
	if err != nil {
		common.Logger.Error("d.SetPhoneCode(%v,%v) error(%v)", phone, code, err)
		return  err
	}
	if orgCode != code {
		return  errors.New("code is incorrect")
	}


	return nil

}


func cominKey(pre string, content string) string {
	return pre + content
}
