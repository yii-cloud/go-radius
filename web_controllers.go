package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadControllers(router *gin.Engine) {
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.POST("/manager/info", fetchManagerInfo)
}

func login(c *gin.Context) {
	var manager SysManager
	c.ShouldBindJSON(&manager)
	manager.Password = encrypt(manager.Password)
	ok, _ := engine.Get(&manager)

	if !ok {
		c.JSON(http.StatusOK, newErrorJsonResult("username or password is incorrect"))
		return
	}

	session := GlobalSessionManager.CreateSession(c)
	manager.Password = ""
	session.SetAttr("manager", manager)

	c.JSON(http.StatusOK, newSuccessJsonResult("success", session.SessionId()))
}

func logout(c *gin.Context) {
	GlobalSessionManager.DestroySession(c)
	c.JSON(http.StatusOK, JsonResult{Code: 0, Message: "success"})
}

func fetchManagerInfo(c *gin.Context) {
	token := c.GetHeader(SessionName)
	session := GlobalSessionManager.GetSession(token)
	c.JSON(http.StatusOK, JsonResult{Code: 0, Message: "suceess", Data: session.GetAttr("manager")})
}