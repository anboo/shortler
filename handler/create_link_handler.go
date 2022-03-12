package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "github.com/teris-io/shortid"
    "log"
    "net/url"
    "shortler/db/dao"
)

func CreateLinkHandler(c *gin.Context) {
    type createLinkRequest struct {
        Token     *string `json:"token" validate:"required,min=4,max=15"`
        ExpiresAt *int    `json:"expiresAt" validate:"required"`
        Link      *string `json:"link" validate:"required,url"`
    }
    
    req := createLinkRequest{}
    err := c.BindJSON(&req)
    if err != nil {
        log.Print("bind json error " + err.Error())
        c.JSON(400, CreateErrorResponse("bad_json"))
        return
    }
    
    v := validator.New()
    validateError := v.Struct(req)
    
    validationErrorsResp := GetValidationErrorResponse(validateError)
    if validationErrorsResp != nil {
        c.JSON(400, validationErrorsResp)
        return
    }
    
    u, errParseUrl := url.Parse(*req.Link)
    if errParseUrl != nil {
        log.Print("cannot parse url " + *req.Link + ": " + err.Error())
        c.JSON(400, CreateErrorResponse("parse_link_error"))
        return
    }
    
    if u.Scheme == "" && u.Host == "" {
        log.Print("empty scheme and host")
        c.JSON(400, CreateErrorResponse("bad_link"))
        return
    }
    
    linkDAO := dao.LinkDAO{}
    
    alreadyExistsByLink := linkDAO.GetByLink(*req.Link)
    if alreadyExistsByLink != nil {
        c.JSON(201, gin.H{"token": alreadyExistsByLink.Token})
        return
    }
    
    if req.Token == nil || *req.Token == "" {
        sid, err := shortid.New(1, shortid.DefaultABC, 2342)
        if err != nil {
            log.Print("error create sid " + err.Error())
            c.JSON(500, CreateErrorResponse("internal_server_error"))
            return
        }
        
        *req.Token, err = sid.Generate()
        
        if err != nil {
            log.Print("error generate short id " + err.Error())
            c.JSON(500, CreateErrorResponse("internal_server_error"))
            return
        }
    } else {
        alreadyExists := linkDAO.GetByShort(*req.Token)
        if alreadyExists != nil {
            c.JSON(400, CreateErrorResponse("token_already_exists"))
            
            return
        }
    }
    
    link := dao.Link{}
    link.Link = *req.Link
    link.Token = *req.Token
    link.CreatedById = 0
    
    err = linkDAO.Create(link)
    if err != nil {
        c.JSON(400, CreateErrorResponse("internal_server_error_2"))
        return
    }
    
    c.JSON(201, gin.H{"token": link.Token})
}
