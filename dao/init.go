package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	ikhnaieDB *gorm.DB
)

func init()  {
	var err error
	ikhnaieDB, err = gorm.Open("mysql", "root:sqlcareful@/ikhnaie?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
