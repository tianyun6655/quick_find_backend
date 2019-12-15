package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"quick_find_backend/common"
	"quick_find_backend/friend/conf"
)

type Dao struct {
	config *conf.Config
	sql    *xorm.Engine
	redis  *redis.Client

}

func New(config *conf.Config) (*Dao, error) {

	//mysql
	mysqlInform := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4",
		config.Mysql.UserName,
		config.Mysql.Password,
		config.Mysql.Address,
		config.Mysql.DBName)
	engin, err := xorm.NewEngine("mysql", mysqlInform)
	if err != nil {
		common.Logger.Error("[init.mysql.New()] error(%v)", err)
		return nil, err
	}


	engin.SetMaxIdleConns(config.Mysql.MaxIde)
	engin.SetMaxOpenConns(config.Mysql.MaxOpen)

	//redis
	return &Dao{
		config: config,
		sql:    engin,
		redis: redis.NewClient(&redis.Options{
			Addr:     config.Redis.Address,
			PoolSize: config.Redis.PoolSize,
			DB:       config.Redis.DB,
		}),
	}, nil

}
