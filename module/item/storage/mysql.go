package storage

import (
	"errors"
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/samber/do"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConfig struct {
	mysqlCfg.Config
}

func (c *mysqlConfig) Validate() error {
	if c.User == "" {
		return errors.New("mysql: user is empty")
	}
	if c.Passwd == "" {
		return errors.New("mysql: password is empty")
	}
	if c.Net == "" {
		return errors.New("mysql: network type is empty")
	}
	if c.Addr == "" {
		return errors.New("mysql: address is empty")
	}
	if c.DBName == "" {
		return errors.New("mysql: database name is empty")
	}
	return nil
}

func NewMySQL(i *do.Injector) (*gorm.DB, error) {
	var cfg = &mysqlConfig{
		Config: mysqlCfg.Config{
			User:   os.Getenv("MYSQL_USER"),
			Passwd: os.Getenv("MYSQL_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("MYSQL_ADDR"),
			DBName: os.Getenv("MYSQL_DB"),
		},
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	dsn := cfg.FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database", db)
	return db, nil
}
