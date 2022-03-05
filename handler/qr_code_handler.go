package handler

import (
    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
    "github.com/gin-gonic/gin"
    "image/png"
    "log"
    "shortler/db/dao"
    "strconv"
)

type BufferedWriter struct {
    img []byte
}

func (w *BufferedWriter) Write(p []byte) (n int, err error) {
    w.img = append(w.img, p...)

    return len(w.img), nil
}

func QrCodeHandler(c *gin.Context) {
    linkDao := dao.LinkDAO{}

    token := c.Param("token")
    width, _ := strconv.Atoi(c.Param("width"))

    if width <= 0 {
        width = 300
    } else if width >= 1000 {
        width = 1000
    }

    link := linkDao.GetByShort(token)
    if link == nil {
        c.JSON(404, gin.H{"error": "not_found"})
        return
    }

    log.Println("generate qr code for link " + link.Token + " " + link.Link)

    qrCode, _ := qr.Encode(link.Link, qr.H, qr.Unicode)
    qrCode, _ = barcode.Scale(qrCode, width, width)

    writer := &BufferedWriter{}
    png.Encode(writer, qrCode)

    c.Header("Content-Type", "image/png")
    c.String(200, string(writer.img))
}
