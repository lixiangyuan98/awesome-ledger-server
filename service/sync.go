package service

import (
    "../bean"
    "../helper"
)

// Sync 同步服务端数据
// @param localItems: 客户端的条目
// @return localInsert: 要求客户端新增的条目
//		   localUpdate: 要求客户端更新的条目
//         localDelete: 要求客户端删除的条目
//         remoteUpdate: 服务端要更新的条目的UUID
//         remoteInsert: 服务端要新增的条目的UUID
func Sync(localItems []*bean.Item) (localInsert, localUpdate []*bean.Item,
                                    localDelete, remoteUpdate, remoteInsert []*string) {
    itemHelper := helper.GetItemHelper()
    for _, item := range localItems {
        selectedItem := &bean.Item{};
        itemHelper.Where("uuid = ?", item.UUID).Take(selectedItem)
        if itemHelper.NewRecord(selectedItem) {
            if item.DeletedAt == nil {
                // 添加远程记录
                remoteInsert = append(remoteInsert, &item.UUID)
            } else {
                // 删除本地记录
                localDelete = append(localDelete, &item.UUID)
            }
        } else {
            if item.DeletedAt != nil {
                // 彻底删除本地记录
                localDelete = append(localDelete, &item.UUID)
                // 删除远程记录
                itemHelper.Delete(selectedItem)
            } else {
                if item.UpdatedAt.Unix() < selectedItem.UpdatedAt.Unix() {
                    // 更新本地记录
                    localUpdate = append(localUpdate, selectedItem)
                } else if item.UpdatedAt.Unix() > selectedItem.UpdatedAt.Unix() {
                    // 更新远程记录
                    remoteUpdate = append(remoteUpdate, &item.UUID)
                }
            }
        }
    }
    remoteItems := []*bean.Item{}
    itemHelper.Find(&remoteItems)
    // O(n*m)
    // TODO: 二分或hash
    for _, remoteItem := range remoteItems {
        exists := false
        for _, item := range localItems {
            if remoteItem.UUID == item.UUID {
                exists = true
            }
        }
        if !exists {
            // 添加本地记录
            localInsert = append(localInsert, remoteItem)
        }
    }
    return
}

// Insert 将新条目加入服务端
func Insert(localItems []*bean.Item) (deletedItems []*string) {
    itemHelper := helper.GetItemHelper()
    for _, item := range localItems {
        item.ID = 0     // 避免主键冲突
        if err:= itemHelper.Create(item); err != nil {
            deletedItems = append(deletedItems, &item.UUID)
        }
    }
    return
}

// Update 将服务端的条目更新
func Update(localItems []*bean.Item)  {
    itemHelper := helper.GetItemHelper()
    for _, item := range localItems {
        selectedItem := &bean.Item{}
        itemHelper.Where("uuid = ?", item.UUID).Take(selectedItem)
        itemHelper.Model(selectedItem).Updates(map[string]interface{}{
            "date": item.Date,
            "item_type": item.ItemType,
            "item_kind": item.ItemKind,
            "address": item.Address,
            "money": item.Money,
            "comment": item.Comment,
        })
    }
}