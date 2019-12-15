package dao

import (
	"quick_find_backend/common"
	"quick_find_backend/user/model"
)

//Get user by phone
func (d *Dao) GetOneByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	if result, err := d.sql.Where("phone=?", phone).Get(user); err != nil {
		common.Logger.Error("d.user.GetOneByPhone(%s) error(%v)", phone, err)
		return nil, err
	} else if !result {
		common.Logger.Warn("d.user.GetOneByPhone(%s) error(%s)", phone, "no such user")
		return nil, nil
	}
	return user, nil
}

func (d *Dao) UpdateYunxinToken(uid int, token string) error {
	user := &model.User{
		YunxinToken: token,
	}
	if _, err := d.sql.Where("uid=?", uid).Update(user); err != nil {
		common.Logger.Warn("d.user.UpdateYunxinToken(%d,%s) error(%s)", uid, token, err.Error())
		return err
	}
	return nil
}

//Insert user
func (d *Dao) Insert(user *model.User) error {
	_, err := d.sql.Insert(user)
	if err != nil {
		common.Logger.Warn("d.user.Insert(%s) error(%v)", user.Phone, err)
		return err
	}
	return nil
}

func (d *Dao) UpdateById(use *model.User) error {
	_, err := d.sql.Where("uid=?", use.Uid).Update(use)
	if err != nil {
		common.Logger.Error("d.user.UpdateById(%#T) error(%v)", use, err.Error())
		return err
	}
	return nil
}
