package controller

import (
    "github.com/gin-gonic/gin"
    "../bean"
    "../service"
)

// Route 对所有URL进行路由
// @param server: 服务器实例
func Route(server *gin.Engine) {
    server.GET("/sync", sync)
    server.POST("/insert", insert)
    server.POST("/update", update)
}

// sync 处理同步数据的请求
// 比对数据的差异, 返回要求客户端更新/新增的条目的信息和服务端需要更新/新增的条目的UUID
// Request: [{"uuid": "uuid1", "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null}]
func sync(c *gin.Context) {
    localItems := []*bean.Item{}
    if err := c.ShouldBindJSON(&localItems); err != nil {
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    localInsert, localUpdate, localDelete, remoteUpdate, remoteInsert := service.Sync(localItems)
    c.JSON(200, gin.H{
        "localInsert": localInsert,
        "localUpdate": localUpdate,
        "localDelete": localDelete,
        "remoteUpdate": remoteUpdate,
        "remoteInsert": remoteInsert,
    })
}
 
// insert 将条目新增到服务端
// Request: [{"uuid": "uuid1", "createdAt": "2019-06-01T04:12:01+08:00", 
//           "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null
//           "date": "2019-06-01T04:12:01+08:00", "itemType": "outgoing",
//           "itemKind": "FOOD", "address", "", "money": 100.2, "comment": ""}]
func insert(c *gin.Context) {
    localItems := []*bean.Item{}
    if err := c.ShouldBindJSON(&localItems); err != nil {
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    service.Insert(localItems)
    c.JSON(200, nil)
}

// update 更新服务端的条目
// Request: [{"uuid": "uuid1", "createdAt": "2019-06-01T04:12:01+08:00", 
//           "updatedAt": "2019-06-01T04:12:01+08:00", "deletedAt": null
//           "date": "2019-06-01T04:12:01+08:00", "itemType": "outgoing",
//           "itemKind": "FOOD", "address", "", "money": 100.2, "comment": ""}]
func update(c *gin.Context) {
    localItems := []*bean.Item{}
    if err := c.ShouldBindJSON(&localItems); err != nil {
        c.JSON(400, gin.H{"message": "Invalid Arguments"})
        return
    }
    service.Update(localItems)
    c.JSON(200, nil)
}
