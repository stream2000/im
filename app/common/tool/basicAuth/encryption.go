/*
@Time : 2020/1/21 18:18
@Author : Minus4
*/
package basicAuth

import (
	"crypto/md5"
	"fmt"
	"io"
)

var (
	Salt1 string
	Salt2 string
)

func Setup(salt1, salt2 string) {
	Salt1 = salt1
	Salt2 = salt2
}
func EncryptAccount(email, password string) string {
	h := md5.New()
	_, _ = io.WriteString(h, password)

	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))

	_, _ = io.WriteString(h, Salt1)
	_, _ = io.WriteString(h, email)
	_, _ = io.WriteString(h, Salt2)
	_, _ = io.WriteString(h, passwordMd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}
