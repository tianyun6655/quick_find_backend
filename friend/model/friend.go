package model


type Friend struct {
	Uid int32
	Friend int32
}

type FriendJoinUser struct {
	Uid int32
	Friend int32
	UserName string
}