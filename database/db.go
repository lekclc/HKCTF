package Db

import (
	cfg "ctf/config"
	"ctf/logic"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Get("sql.name"), cfg.Get("sql.password"), cfg.Get("sql.host"), cfg.Get("sql.port"), cfg.Get("sql.dbname"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	cfg.DB = db
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&Level{})
	db.AutoMigrate(&Images{})
	db.AutoMigrate(&Container{})
	db.AutoMigrate(&Ctf{})

	admin_passwd := logic.Passwd_hash(cfg.Get("admin.password").(string))
	var user User
	var admin Admin
	res := db.Where("name = ?", cfg.Get("admin.name")).First(&user)
	if res.Error != nil {
		//找不到admin用户，创建
		db.Create(&User{Name: cfg.Get("admin.name").(string), Passwd: admin_passwd, ContNum: 0})
	} else {
		db.Model(&user).Update("passwd", admin_passwd)
	}
	db.Where("name = ?", cfg.Get("admin.name")).First(&user)
	res = db.Where("name = ?", cfg.Get("admin.name")).First(&admin)
	if res.Error != nil {
		db.Create(&Admin{UserID: user.UserID, Name: cfg.Get("admin.name").(string), Passwd: admin_passwd})
	} else {
		db.Model(&admin).Update("passwd", admin_passwd)
		db.Model(&admin).Update("UserID", user.UserID)
	}
}
