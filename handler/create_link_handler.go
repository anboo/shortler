package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/teris-io/shortid"
    "log"
    "net/url"
    "shortler/db/dao"
)

func CreateLinkHandler(c *gin.Context) {
    type createLinkRequest struct {
        Token     string `json:"token"`
        ExpiresAt *int   `json:"expiresAt"`
        Link      string `json:"link"`
    }

    req := createLinkRequest{}
    err := c.BindJSON(&req)
    if err != nil {
        log.Print("bind json error " + err.Error())
        c.JSON(400, gin.H{"error": "bad_request"})
        return
    }

    u, errParseUrl := url.Parse(req.Link)
    if errParseUrl != nil {
        log.Print("cannot parse url " + req.Link + ": " + err.Error())
        c.JSON(400, gin.H{"error": "bad_link"})
        return
    }

    if u.Scheme == "" && u.Host == "" {
        log.Print("empty scheme and host")
        c.JSON(400, gin.H{"error": "bad_link"})
        return
    }

    linkDAO := dao.LinkDAO{}

    alreadyExistsByLink := linkDAO.GetByLink(req.Link)
    if alreadyExistsByLink != nil {
        c.JSON(201, gin.H{
            "token": alreadyExistsByLink.Token,
        })
        return
    }

    if req.Token == "" {
        sid, err := shortid.New(1, shortid.DefaultABC, 2342)
        if err != nil {
            log.Print("error create sid " + err.Error())
            c.JSON(500, gin.H{"error": "internal_server_error"})
            return
        }

        req.Token, err = sid.Generate()

        if err != nil {
            log.Print("error generate short id " + err.Error())
            c.JSON(500, gin.H{"error": "internal_server_error"})
            return
        }
    } else {
        alreadyExists := linkDAO.GetByShort(req.Token)
        if alreadyExists != nil {
            c.JSON(400, gin.H{"error": "token_link_already_exists"})

            return
        }
    }

    link := dao.Link{}
    link.Link = req.Link
    link.Token = req.Token
    link.CreatedById = 0

    err = linkDAO.Create(link)
    if err != nil {
        c.JSON(400, gin.H{"error": "error"})
        return
    }

    c.JSON(201, gin.H{
        "token": link.Token,
    })
}
