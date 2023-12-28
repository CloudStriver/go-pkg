package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

func Newuuid() string {
	uid, err := uuid.NewV4()
	if err != nil {
		logx.Errorf("uuid NewV4异常[%v]\n", err)
	}
	return uid.String()
}
