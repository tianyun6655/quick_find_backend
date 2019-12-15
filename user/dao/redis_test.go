package dao

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDao_SetPhoneCode(t *testing.T) {
	convey.Convey("set phone code",t, func(ctx convey.C) {
		err:=d.SetPhoneCode("11111",1111)
		convey.Convey("err should be err", func(ctx convey.C) {
			convey.So(err,convey.ShouldBeNil)
		})
	})
}
