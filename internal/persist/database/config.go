package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	User        string       `json:"user"`
	Password    string       `json:"password"`
	Host        string       `json:"host"`
	Port        string       `json:"port"`
	Database    string       `json:"database"`
	AutoMigrate bool         `json:"auto_migrate"`
	Logging     bool         `json:"logging"`
	Provider    ProviderType `json:"provider"`
}

type ProviderType string

const (
	MysqlProvider    ProviderType = "mysql"
	SqliteProvider   ProviderType = "sqlite"
	PostgresProvider ProviderType = "postgres"
)

func DbConn(conf *Config) gorm.Dialector {
	switch conf.Provider {
	case MysqlProvider:
		return mysqlConn(conf)
	case PostgresProvider:
		return postgresConn(conf)
	case SqliteProvider:
		fallthrough
	default:
		return sqliteConn(conf)
	}
}

func mysqlConn(conf *Config) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Password, conf.Database)
	return mysql.Open(dsn)
}

func postgresConn(conf *Config) gorm.Dialector {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port)
	return postgres.Open(dsn)
}

func sqliteConn(conf *Config) gorm.Dialector {
	dsn := conf.Database + ".db"
	return sqlite.Open(dsn)
}

func DbConfig(conf *Config) *gorm.Config {
	gormConfig := &gorm.Config{}
	if conf.Logging {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	return gormConfig
}
