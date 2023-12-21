package db

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckDuplicateError(err error) bool {
	logx.Infof("err:[%v]\n", err)
	if mongoError, ok := err.(mongo.WriteException); ok {
		for _, writeError := range mongoError.WriteErrors {
			if writeError.Code == 11000 {
				return true
			}
		}
	}
	return false
}
