package dao

import (
	"quick_find_backend/common"
	"quick_find_backend/friend/model"
)

func(d *Dao)Insert(uid int32, friendId int32)error{
	friend := &model.Friend{
		Uid:    uid,
		Friend: friendId,
	}
	if _, err := d.sql.Insert(friend); err != nil {
		common.Logger.Error("[d.friend]InsertFriend(%d,%d) error:", uid, friend, err.Error())
		return err
	}
	return nil
}

func (d *Dao)JoinUser(uid int32)([]*model.FriendJoinUser,error){
	friends:=make([]*model.FriendJoinUser,0)

	err:=d.sql.Table("friend").
		Join("INNER","user","user.uid=friend.friend").
		Where("friend.uid=?",uid).Find(&friends)
	if err!=nil{
		common.Logger.Error("[d.friend]JoinUser(%d) error:", uid, err.Error())
		return nil,err
	}
	return friends,nil
}