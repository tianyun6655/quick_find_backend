package common

import(
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
)

//生成UUID
func GenerateUUID() (token string, err error) {

	var (
		u1 uuid.UUID
	)
	if u1, err = uuid.NewV4(); err != nil {
		return
	}
	token = fmt.Sprintf("%s", u1)

	return
}

func Sha1(plainText string) string {
	s := sha1.New()
	io.WriteString(s, plainText)
	secret := s.Sum(nil)
	secret16 := hex.EncodeToString(secret)
	return secret16
}