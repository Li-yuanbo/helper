package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"helper/model"
	"helper/service"
	. "helper/utils"
	"log"
	"net/http"
)

func RegisterOrganization(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.RegisterOrganizationReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("register user request:", req)
	resp, err := service.RegisterOrganization(c, &req)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, resp.ResStatus.Code, resp.ResStatus.Msg)
		return
	}
	resp.ResStatus = NewRes(10000, "SUCCESS")
	SucResponse(c, resp)
}

func LoginOrg(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.LoginOrgReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("register user request:", req)
	resp, err := service.LoginOrg(c, &req)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, resp.ResStatus.Code, resp.ResStatus.Msg)
		return
	}
	SucResponse(c, resp)
}

func UnLoginOrg(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.UnLoginOrgReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("register org request:", req)
	resp, err := service.UnLoginOrg(c, &req)
	if err != nil {
		ErrResponse(c, http.StatusInternalServerError, resp.ResStatus.Code, resp.ResStatus.Msg)
		return
	}
	SucResponse(c, resp)
}
