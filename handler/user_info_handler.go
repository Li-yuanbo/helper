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

func RegisterUser(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.RegisterUserReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("register user request:", req)
	//调用service
	resp, err := service.RegisterUser(req, c)
	if err != nil {
		log.Println("register user err: ", err)
		ErrResponse(c, http.StatusInternalServerError, resp.Res.Code, resp.Res.Msg)
		return
	}
	//返回成功
	SucResponse(c, resp)
}

func GetUserInfos(c *gin.Context) {
	//读取request body
	body, err := c.GetRawData()
	if err != nil {
		log.Println("get request body err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10001, ERR_CODE[10001])
		return
	}
	//通过json将request body转换为struct
	var req model.GetUserInfosReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("request body unmarshal err: ", err)
		ErrResponse(c, http.StatusInternalServerError, 10002, ERR_CODE[10002])
		return
	}
	log.Println("register user request:", req)
	resp, err := service.GetUserInfos(c, &req)
	if err != nil {
		log.Println("[service] get users err: ", err)
		ErrResponse(c, http.StatusInternalServerError, resp.Res.Code, resp.Res.Msg)
		return
	}
	SucResponse(c, resp)
}
