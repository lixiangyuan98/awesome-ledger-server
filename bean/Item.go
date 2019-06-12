package bean

import (
	"github.com/jinzhu/gorm"
    "time"
)

// Item 账单条目的实体类
type Item struct {
    gorm.Model   
    UUID string      `gorm:"type:uuid;not null"`
    Date time.Time   `gorm:"type:timestamp;not null"`
    ItemType string  `gorm:"type:varchar(32);not null"`
    ItemKind string  `gorm:"type:varchar(32);not null"`
    Address string   `gorm:"type:varchar(128)"`
    Money float32    `gorm:"type:money;not null"`
    Comment string   `gorm:"type:varchar(256)"`
}

// TableName sets the table's name
func (Item) TableName() string {
    return "ledger_t"
}

