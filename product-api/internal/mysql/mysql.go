package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlParams struct {
	Url           string
	NameDB        string
	AdminUser     string
	AdminPassword string
}

func Start(params MysqlParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", params.AdminUser, params.AdminPassword, params.Url, params.NameDB)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
