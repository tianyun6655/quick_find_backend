package quick_find_services_friend

import (
	"context"
	"quick_find_backend/friend/dao"
)

type FriendServices struct {
	friendDao *dao.Dao
}

func NewFriendServices(dao *dao.Dao)*FriendServices{
	return &FriendServices{
		friendDao:dao,
	}
}
func(fs *FriendServices)AddFriend(context context.Context,resq *AddFriendRequest,resp *FriendResponse)error{
    return fs.friendDao.Insert(resq.Uid,resq.FriendId)
}


func(fs *FriendServices)List(context context.Context,res *ListRequest,resp *ListResponse)error{
	friend,err:=fs.friendDao.JoinUser(res.Uid)
	if err!=nil{
		return err
	}
	*resp=ListResponse{
		List:make([]*FriendJoinUser,len(friend)),
		Uid:res.Uid,
	}
	for i:=0;i<len(friend);i++{
		single:=&FriendJoinUser{
			Friend:friend[i].Friend,
			Name:friend[i].UserName,
		}
		resp.List[i]=single
	}
	return nil
}