/*
@Time : 2020/1/19 15:22
@Author : Minus4
*/
package validate

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

var (
	RegEmailCheck = regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)
)

func init() {
	_ = binding.Validator.RegisterValidation("email", emailCheck)
}

func emailCheck(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return RegEmailCheck.MatchString(email)
}
