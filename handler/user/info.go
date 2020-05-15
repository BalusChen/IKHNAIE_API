package user

import (
	"net/http"
	"strconv"

	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

const (
	statusCode_Base                  = 50000
	StatusCode_OK                    = 200
	StatusCode_MethodONotImplemented = statusCode_Base + 1
	StatusCode_UserNotFound          = statusCode_Base + 101
	StatusCode_UserExist             = statusCode_Base + 102
	StatusCode_WrongPassword         = statusCode_Base + 103
	StatusCode_InactiveUser          = statusCode_Base + 104

	StatusMsg_OK                   = "成功"
	StatusMsg_BadRequest           = "参数错误"
	StatusMsg_ServerInternalError  = "服务器内部错误"
	StatusMsg_MethodNotImplemented = "该接口尚未实现哦"
	StatusMsg_LoginOK              = "登陆成功"
	StatusMsg_UserNotFound         = "该用户不存在"
	StatusMsg_WrongPassword        = "密码错误"
	StatusMsg_RegisterOK           = "注册成功"
	StatusMsg_UserExist            = "该用户已存在"
	StatusMsg_InactiveUser         = "该用户尚未被准入，请联系管理员"
)

func Info(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO: get user from db

	ctx.JSON(http.StatusOK, gin.H{
		"user":        1,
		"status_code": StatusCode_MethodONotImplemented,
		"status_msg":  StatusMsg_MethodNotImplemented,
	})
}

func List(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	status, err := strconv.ParseInt(ctx.Query("status"), 10, 32)
	if err != nil || (status != model.UserStatus_Inactive && status != model.UserStatus_Active) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  StatusMsg_BadRequest,
		})
		return
	}

	users, err := dao.GetUsersByStatus(ctx, int32(status))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  StatusMsg_ServerInternalError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users":       users,
		"status_code": http.StatusOK,
		"status_msg":  StatusMsg_OK,
	})
}
