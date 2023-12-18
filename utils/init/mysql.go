package init

import (
	"github.com/CloudStriver/go-pkg/utils/gormlogger"
	"gorm.io/driver/mysql"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitMysql(DataSource string) *gorm.DB {
	GormDB, err := gorm.Open(mysql.Open(DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表
		},
		Logger:                                   gormlogger.New(gormlogger.Config{LogLevel: logger.Info}),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logx.Errorf("gorm.Open Error: %v", err)
	}
	Db, err := GormDB.DB()
	if err != nil {
		logx.Infof("GormDB.DB() Error: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Db.SetMaxIdleConns(64)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(64)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Db.SetConnMaxIdleTime(time.Minute)
	// SetConnMaxLifetime 设置了连接存活的最大时间。
	Db.SetConnMaxLifetime(5 * time.Minute)

	return GormDB
}
