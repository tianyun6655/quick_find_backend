// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_service.proto

package quick_find_services_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserServices service

type UserServicesService interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*User, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*UpdateResponse, error)
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...client.CallOption) (*AddFriendResponse, error)
}

type userServicesService struct {
	c    client.Client
	name string
}

func NewUserServicesService(name string, c client.Client) UserServicesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "quick.find.services.user"
	}
	return &userServicesService{
		c:    c,
		name: name,
	}
}

func (c *userServicesService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserServices.Login", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserServices.Register", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesService) Update(ctx context.Context, in *User, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "UserServices.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicesService) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...client.CallOption) (*AddFriendResponse, error) {
	req := c.c.NewRequest(c.name, "UserServices.AddFriend", in)
	out := new(AddFriendResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserServices service

type UserServicesHandler interface {
	Login(context.Context, *LoginRequest, *User) error
	Register(context.Context, *RegisterRequest, *User) error
	Update(context.Context, *User, *UpdateResponse) error
	AddFriend(context.Context, *AddFriendRequest, *AddFriendResponse) error
}

func RegisterUserServicesHandler(s server.Server, hdlr UserServicesHandler, opts ...server.HandlerOption) error {
	type userServices interface {
		Login(ctx context.Context, in *LoginRequest, out *User) error
		Register(ctx context.Context, in *RegisterRequest, out *User) error
		Update(ctx context.Context, in *User, out *UpdateResponse) error
		AddFriend(ctx context.Context, in *AddFriendRequest, out *AddFriendResponse) error
	}
	type UserServices struct {
		userServices
	}
	h := &userServicesHandler{hdlr}
	return s.Handle(s.NewHandler(&UserServices{h}, opts...))
}

type userServicesHandler struct {
	UserServicesHandler
}

func (h *userServicesHandler) Login(ctx context.Context, in *LoginRequest, out *User) error {
	return h.UserServicesHandler.Login(ctx, in, out)
}

func (h *userServicesHandler) Register(ctx context.Context, in *RegisterRequest, out *User) error {
	return h.UserServicesHandler.Register(ctx, in, out)
}

func (h *userServicesHandler) Update(ctx context.Context, in *User, out *UpdateResponse) error {
	return h.UserServicesHandler.Update(ctx, in, out)
}

func (h *userServicesHandler) AddFriend(ctx context.Context, in *AddFriendRequest, out *AddFriendResponse) error {
	return h.UserServicesHandler.AddFriend(ctx, in, out)
}
