package main

import (
    "github.com/gin-gonic/gin"
    "shortler/handler"
)

func main() {
    r := gin.Default()

    r.POST("/link", handler.CreateLinkHandler)
    r.GET("/s/:token", handler.RedirectLinkHandler)
    r.GET("/qr/:token/:width", handler.QrCodeHandler)

    r.Run()
}
