package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/yanming-open/wechat/mp"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Token          = "7kcI9CZfN7lrlWXzWYpFPQtXLc0h"
	AppId          = "wx2c5b27e19c98d3c0"
	AppSecret      = "8c3042dbb692206e90a6d035fd3459a7"
	EncodingAesKey = "A015JQ13miwlzoNlHu5jt4xV8DwB7N1zBZmj1"
)

func main() {
	MP := mp.NewMp(Token, EncodingAesKey, AppId, AppSecret)

	route := gin.Default()
	route.GET("/api/mp", func(c *gin.Context) {
		signature := c.Query("signature")
		timestamp := c.Query("timestamp")
		nonce := c.Query("nonce")
		echostr := c.Query("echostr")
		ok := mp.CheckSignature(signature, timestamp, nonce, Token)
		if !ok {
			return
		}
		_, _ = c.Writer.WriteString(echostr)
	})

	route.POST("/api/mp", func(c *gin.Context) {
		//signature := c.Query("signature")
		//timestamp := c.Query("timestamp")
		//nonce := c.Query("nonce")
		//openid := c.Query("openid")
		//encrypt_type := c.Query("encrypt_type")
		//msg_signature := c.Query("msg_signature")
		//encryptMsg := mp.EncryptMsg{}
		body, _ := ioutil.ReadAll(c.Request.Body)
		//xml.Unmarshal(body, &encryptMsg)
		var baseMessage = mp.BizMessage{}
		xml.Unmarshal(body, &baseMessage)
		log.Println(baseMessage)
		switch baseMessage.MsgType.Text {
		case mp.Text:
			break
		}
		//_, _ = c.Writer.WriteString("success")
	})

	route.GET("/api/callback/login", func(c *gin.Context) {
		code := c.Query("code")
		//state := c.Query("state")
		resp, err := MP.GetSnsUserAccessToken(code)
		user, _ := MP.GetSnsUserInfo(resp.AccessToken, resp.OpenId)
		userMp, err := MP.GetUserInfo(resp.OpenId)
		c.JSON(http.StatusOK, gin.H{"user1": user, "user2": userMp, "err": err})
	})
	route.GET("/", func(c *gin.Context) {
		//resp ,_ := MP.BatchGetMaterial(mp.News,0,10)
		resp := MP.GetJsTicketSignature("m.d.yanming.info")
		c.JSON(http.StatusOK, gin.H{"data": resp})
	})
	route.Run(":9081")
}
