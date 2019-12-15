package http

type RegisterRequest struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

type RegisterResponse struct {
	Token string `json:"token"`
	Uid   int    `json:"uid"`
}

type LoginResponse struct {

}

type UpdateRequest struct {
	Uid      int    `json:"uid"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}
