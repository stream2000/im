/*
@Time : 2020/1/28 00:09
@Author : Minus4
*/
package ecode

import "github.com/bilibili/kratos/pkg/ecode"

func Init() {
	var em = map[int]string{
		120001:                 "请求的小组不存在",
		120002:                 "用户已经在小组里了",
		ecode.ServerErr.Code(): "internal server error",
	}
	ecode.Register(em)
}
