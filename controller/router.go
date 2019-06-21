package controller

import (
    "github.com/gin-gonic/gin"
    "../bean"
    "../service"
    "fmt"
)

// Route 对所有URL进行路由
// @param server: 服务器实例
func Route(server *gin.Engine) {
    server.POST("/sync", sync)
    server.POST("/insert", insert)
    server.POST("/update", update)
}

// 请求的结构体
type itemsRequest struct {
    Data []*bean.Item
}

// sync 处理同步数据的请求
// 比对数据的差异, 返回要求客户端更新/新增的条目的信息和服务端需要更新/新增的条目的UUID
// Request: {"data": [{"uuid": "uuid1", "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null}]}
func sync(c *gin.Context) {
    var localItems itemsRequest
    if err := c.ShouldBindJSON(&localItems); err != nil {
        fmt.Println(err)
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    for _, item := range localItems.Data {
        fmt.Println(item)
    }
    localInsert, localUpdate, localDelete, remoteUpdate, remoteInsert := service.Sync(localItems.Data)
    c.JSON(200, gin.H{
        "localInsert": localInsert,
        "localUpdate": localUpdate,
        "localDelete": localDelete,
        "remoteUpdate": remoteUpdate,
        "remoteInsert": remoteInsert,
    })
}

// insert 将条目新增到服务端
// Request: {"data": [{"uuid": "uuid1", "createdAt": "2019-06-01T04:12:01+08:00", 
//           "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null
//           "date": "2019-06-01T04:12:01+08:00", "itemType": "outgoing",
//           "itemKind": "FOOD", "address": "", "money": 100.2, "comment": ""}]}
func insert(c *gin.Context) {
    var localItems itemsRequest
    if err := c.ShouldBindJSON(&localItems); err != nil {
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    for _, item := range localItems.Data {
        fmt.Println(item)
    }
    localDelete := service.Insert(localItems.Data)
    c.JSON(200, gin.H{
        "localDelete": localDelete,
    })
}

// update 更新服务端的条目
// Request: {"data": [{"uuid": "uuid1", "createdAt": "2019-06-01T04:12:01+08:00", 
//           "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null
//           "date": "2019-06-01T04:12:01+08:00", "itemType": "outgoing",
//           "itemKind": "FOOD", "address": "", "money": 100.2, "comment": ""}]}
func update(c *gin.Context) {
    var localItems itemsRequest
    if err := c.ShouldBindJSON(&localItems); err != nil {
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    for _, item := range localItems.Data {
        fmt.Println(item)
    }
    service.Update(localItems.Data)
    c.JSON(200, nil)
}
