package dao

import (
	"github.com/smartystreets/goconvey/convey"
	"quick_find/user/model"
	"testing"
)

func TestDao_Insert(t *testing.T) {
	convey.Convey("insert user",t, func(ctx convey.C) {
		user:=&model.User{}
		err:=d.Insert(user)
		ctx.Convey("err should be nil", func(ctx convey.C) {
			convey.So(err,convey.ShouldBeNil)
		})
	})
}

func TestDao_GetOneByPhone(t *testing.T) {
	convey.Convey("get user by phone",t, func(ctx convey.C) {
		user,err:=d.GetOneByPhone("123")
		ctx.Convey("user should not be nill,err should be nil", func(ctx convey.C) {
			convey.So(err,convey.ShouldBeNil)
			convey.So(user,convey.ShouldNotBeNil)
		})
	})
}