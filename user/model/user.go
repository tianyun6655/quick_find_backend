package model

type User struct {
	Uid         int    `json:"uid" xorm:"not null pk autoincr INT(10)"`
	UserName    string `json:"user_name"`
	Age         int    `json:"age"`
	Phone       string `json:"-"`
	Sex         int    `json:"sex"`
	YunxinToken string `json:"yunxin_token"`
}


type Friend struct {
	Uid int32
	Friend int32
}

type FriendJoinUser struct {
	Uid int32  `json:"uid"`
	UserName string `json:"name"`
}