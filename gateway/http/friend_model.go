package http


type AddFriendRequest struct {
	Uid int32  `json:"uid"`
	FriendId int32  `json:"friend_id"`
}



type FriendListRequest struct {
	Uid int32 `json:"id"`
}