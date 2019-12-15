package dao

import (
	"flag"
	"os"
	"quick_find/common"
	"quick_find/user/conf"
	"testing"
)

var (
	d *Dao
)

func TestMain(m *testing.M) {
	//init log
	common.InitLogger()
	//init mysql
	flag.Set("conf", "../cmd/friend-test.toml")
	flag.Parse()

	err := conf.Init()
	if err != nil {
		panic(err)
	}
	d, err = New(conf.Conf)
	if err != nil {
		panic(err)
	}
	m.Run()
	os.Exit(0)
}
