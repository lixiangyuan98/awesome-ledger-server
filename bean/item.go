package bean

import (
    "time"
)

// Item 账单条目的实体类
type Item struct {
    ID int64
    Date time.Time
    ItemType string
    ItemKind string
    Address string
    Money float32
    Comment string
    UpdateTime time.Time
}

