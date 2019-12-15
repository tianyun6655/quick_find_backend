package dao

import (
	"quick_find_backend/common"
	"quick_find_backend/user/model"
)

func (d *Dao) InsertFriend(uid int32, friendId int32) error {
	friend := &model.Friend{
		Uid:    uid,
		Friend: friendId,
	}
	if _, err := d.sql.Insert(friend); err != nil {
		common.Logger.Error("[d.friend]InsertFriend(%d,%d) error:", uid, friendId, err.Error())
		return err
	}
	return nil
}

func (d *Dao) FriendList(uid int32) ([]*model.FriendJoinUser, error) {
	friend := make([]*model.FriendJoinUser, 0)
	if err := d.sql.Table("friend").
		Join("INNER", "user", "user.uid=friend.friendId").
		Find(&friend); err != nil {
		common.Logger.Error("[d.friend]FriendList(%d) error:", uid, err.Error())
		return nil, err
	}
	return friend, nil
}




