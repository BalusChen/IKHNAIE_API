package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

type userInfo2 struct {
	UserName     string    `json:"user_name"`
	UserID       string    `json:"user_id"`
	Status       int32     `json:"status"`
	Organization string    `json:"organization"`
	RegisterTime time.Time `json:"register_time"`
	LastLogin    time.Time `json:"last_login"`
}

func Info(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO: get user from db

	ctx.JSON(http.StatusOK, gin.H{
		"user":        1,
		"status_code": constant.StatusCode_MethodONotImplemented,
		"status_msg":  constant.StatusMsg_MethodNotImplemented,
	})
}

func List(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var status []int32
	statusStr, ok := ctx.GetQuery("status")
	if !ok {
		status = []int32{model.UserStatus_Active, model.UserStatus_Inactive}
	} else {
		oneStatus, err := strconv.ParseInt(statusStr, 10, 32)
		if err != nil || (oneStatus != model.UserStatus_Inactive && oneStatus != model.UserStatus_Active) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  constant.StatusMsg_InvalidParams,
			})
			return
		}

		status = []int32{int32(oneStatus)}
	}

	users, err := dao.GetUsersByStatus(ctx, status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}

	userInfos := make([]userInfo2, len(users))
	for i := 0; i < len(users); i++ {
		userInfos[i].UserName = users[i].UserName
		userInfos[i].UserID = users[i].UserID
		userInfos[i].Status = users[i].Status
		userInfos[i].Organization = users[i].Organization
		userInfos[i].RegisterTime = users[i].RegisterTime
		userInfos[i].LastLogin = users[i].LastLogin
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users":       userInfos,
		"status_code": http.StatusOK,
		"status_msg":  constant.StatusMsg_OK,
	})
}
