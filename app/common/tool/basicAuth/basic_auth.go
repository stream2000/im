/*
@Time : 2020/1/22 16:03
@Author : Minus4
*/
package basicAuth

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type basicAuthError struct {
	internalError error
}

func newBasicAuthError(e error) basicAuthError {
	return basicAuthError{internalError: e}
}

func (e basicAuthError) Error() string {
	return fmt.Sprintf("basic auth failed with error %s", e.internalError.Error())
}

func ParseBasicHeader(header string) (email, password string, err error) {
	if header == "" {
		err = newBasicAuthError(fmt.Errorf("wrong basic header: %s", header))
		return
	}
	s := strings.SplitN(header, " ", 2)
	if len(s) != 2 {
		err = newBasicAuthError(fmt.Errorf("wrong basic header format: %s", header))
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		err = newBasicAuthError(fmt.Errorf("failed parse based64 code: %s", header))
		return
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 || pair[0] == "" || pair[1] == "" {
		err = newBasicAuthError(fmt.Errorf("wrong encoded infomation: %s", header))
		return
	}
	email = pair[0]
	password = pair[1]
	return
}
