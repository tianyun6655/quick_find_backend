package conf

import (
	"github.com/BurntSushi/toml"
)

var(
	Conf  =&Config{}
	confPath string
)

type Config struct {
	Mysql *MysqlConf
	Redis *RedisConf
	Env string

}

type MysqlConf struct {
	UserName string
	Password string
	Address  string
	DBName   string
	MaxIde   int
	MaxOpen  int
}

type RedisConf struct {
	Address string
	PoolSize int
	DB int
}

func init(){
	//flag.StringVar(&confPath, "conf", "./friend-test.toml", "config path")
}

func Init()(err error){
	_, err = toml.DecodeFile("./user-test.toml",Conf)
	return
}
