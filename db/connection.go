package db

import (
	"cawall-be/config"
	"fmt"

	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)
	fmt.Println(dataSourceName)

	//db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	//if err != nil {
	//	panic(err.Error())
	//}

	//userSeeder := seeders.NewUserSeeder(db)
	//userSeeder.SetUsers()

	//return db
	return &gorm.DB{}
}
