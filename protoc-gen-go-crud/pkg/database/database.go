package database

import (
	secretscfg "bitbucket.org/antalphadev/earn/configs/common_secrets"
	"bitbucket.org/antalphadev/earn/pkg/dictionary"
	"bitbucket.org/antalphadev/earn/pkg/logger"
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"strings"
	"time"
)

var DB *gorm.DB

type MysqlCfg struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func NewDatabase(env string) {
	var masterDsn string

	for name, c := range secretscfg.Databases.Mysql {
		if strings.Contains(name, "master") {
			masterDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
		}
	}
	zLogger := zapLogger(env)

	db, err := gorm.Open(mysql.Open(masterDsn), &gorm.Config{Logger: zLogger})
	if err != nil {
		panic(fmt.Sprintf("gorm.Open error: %v\n", err))
	}

	// TODO: add slave db

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("db.DB error: %v\n", err))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	DB = db.Debug()
}

func zapLogger(env string) gormlogger.Interface {
	zaplogger := logger.Get(context.Background()).Named("gorm")
	zLogger := zapgorm2.Logger{
		ZapLogger:                 zaplogger,
		LogLevel:                  gormlogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
		Context: func(ctx context.Context) (fields []zapcore.Field) {
			return fields
		},
	}

	if env != dictionary.EnvProduction {
		zLogger.LogLevel = gormlogger.Info
	}
	zLogger.SetAsDefault()
	return zLogger
}
