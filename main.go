package main

import (
    "github.com/gin-gonic/gin"
    "./controller"
)

func main() {
    server := gin.Default()
    controller.Route(server)
    server.Run()
}

