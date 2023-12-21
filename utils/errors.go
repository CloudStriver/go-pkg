package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

const (
	Normal = iota
	UnNormal
)

func GetCode(err error) *errors.CodeMsg {
	if codeErr, ok := err.(*errors.CodeMsg); ok {
		return codeErr
	} else {
		logx.Errorf("err[%v] not CodeMsg", err)
		return nil
	}
}
