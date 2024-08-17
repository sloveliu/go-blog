package model

import (
	"blog/global"
	"blog/pkg/setting"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	Id         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDbEngine(database *setting.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local", database.Username, database.Password, database.Host, database.Port, database.Name, database.Charset, database.ParseTime)
	config := &gorm.Config{}
	if global.Server.Mode == "debug" {
		config.Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(database.MaxOpenConns)
	return db, nil
}
