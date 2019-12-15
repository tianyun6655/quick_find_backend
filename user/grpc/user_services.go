package quick_find_services_user

import (
	"context"
	"errors"
	"quick_find_backend/common"
	"quick_find_backend/user/conf"
	"quick_find_backend/user/dao"
	"quick_find_backend/user/model"
	"quick_find_backend/user/utils"
)

type UserService struct {
	userDao *dao.Dao
}

func NewUserService(dao *dao.Dao) *UserService {
	return &UserService{
		userDao: dao,
	}
}

func (us *UserService) Login(ctx context.Context, req *LoginRequest, rsp *User) error {

	user, err := us.userDao.GetOneByPhone(req.Phone)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("please register first")
	}

	*rsp = User{
		Sex:      int32(user.Sex),
		Age:      int32(user.Age),
		UserName: user.UserName,
		Uid:      int32(user.Uid),
		Token:user.YunxinToken,
	}

	return nil
}

func (us *UserService) Update(ctx context.Context, req *User, rsp *UpdateResponse) error {
	user := &model.User{
		Uid:      int(req.Uid),
		UserName: req.UserName,
		Age:      int(req.Age),
		Sex:      int(req.Sex),
	}
	if err := us.userDao.UpdateById(user); err != nil {
		return err
	}
	rsp.Result = true
	return nil
}

func (us *UserService) Register(ctx context.Context, req *RegisterRequest, rsp *User) error {
	user := &model.User{
		Phone: req.Phone,
	}

	if err := us.userDao.Insert(user); err != nil {
		return errors.New("phone is already registered")
	}

	token, err := utils.Resiter(user.Uid)
	if err != nil {
		common.Logger.Error("[s.user] Register(%#v) error:%s", req, err.Error())
		return errors.New("system error")
	}

	if err := us.userDao.UpdateYunxinToken(user.Uid, token); err != nil {
		return errors.New("system error")
	}

	rsp.Uid = int32(user.Uid)

	return nil
}

func (us *UserService) AddFriend(ctx context.Context, req *AddFriendRequest, rsp *AddFriendResponse) error {
	if err := us.userDao.InsertFriend(req.Uid, req.FriendId); err != nil {
		return err
	}
	return nil
}

func Init(conf *conf.Config) (*UserService, error) {
	dao, err := dao.New(conf)
	if err != nil {
		return nil, err
	}
	return &UserService{
		userDao: dao,
	}, nil
}
