package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortler/db/dao"
)

func RedirectLinkHandler(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(404, gin.H{"error": "token_not_found"})
		return
	}

	//clientIP := c.Request.Header.Get("X-Forwarded-For")

	linkDAO := dao.LinkDAO{}
	link := linkDAO.GetByShort(token)

	if link == nil {
		c.JSON(404, gin.H{"error": "token_not_found"})
		return
	}

	c.Redirect(http.StatusFound, link.Link)
}
