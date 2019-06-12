package helper

import (
    "sync"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/Unknwon/goconfig"
    "../bean"
)

// ItemHelper Item的DB对象
var ItemHelper *gorm.DB
var once sync.Once

// GetItemHelper 获取ItemDao对象的指针
func GetItemHelper() *gorm.DB {
    once.Do(func() {
        config, err := goconfig.LoadConfigFile("config.ini")
        if err != nil {
            panic("Load Config Error")
        }
        argString := ""
        args, _ := config.GetSection("PostgreSQL")
        for argName, argValue := range args {
            argString += argName + "=" + argValue + " "
        }
        db, err := gorm.Open("postgres", argString)
        if err != nil {
            panic(err)
        }
        ItemHelper = db
        ItemHelper.AutoMigrate(&bean.Item{})
    })
    return ItemHelper
}

