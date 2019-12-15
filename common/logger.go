package common
import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

var (
	Logger      *logs.BeeLogger
)

func InitLogger() (err error) {
	if Logger, err = createNewLog("info.log"); err != nil {
		return
	}
	return
}

func createNewLog(name string) (log *logs.BeeLogger, err error) {
	config := fmt.Sprintf(`{"filename":"%s","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`, name)
	log = logs.NewLogger(int64(10000))
	err = log.SetLogger("console")
	err = log.SetLogger(logs.AdapterFile, config)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	return
}


