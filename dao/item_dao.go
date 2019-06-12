package dao

import (
    "sync"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/Unknwon/goconfig"
    "../bean"
)

// ItemDao Item的数据接口
type ItemDao struct {
    db *gorm.DB
}

var itemDao *ItemDao
var once sync.Once

// GetItemDao 获取ItemDao对象的指针
func GetItemDao() *ItemDao {
    once.Do(func() {
        itemDao = new(ItemDao)
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
        itemDao.db = db
    })
    return itemDao
}

// Get 根据指定id获取Item
func (*ItemDao)Get(id int64) *bean.Item {
    // todo
    return nil
}

// Insert 插入指定的Item
func (*ItemDao)Insert(item *bean.Item) int64 {
    // todo
    return 0
}

// Update 更新指定的Item
func (*ItemDao)Update(item *bean.Item) {
    // todo
}

// Delete 删除指定的Item 
func (*ItemDao)Delete(item *bean.Item) {
    // todo
}

