package ai

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	result,err:=Translate("你好世界")
	fmt.Println(err)
	fmt.Println(result)
}
