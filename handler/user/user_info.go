package user

import (
	"log"
	"net/http"
	"time"

	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

const (
	baseStatusCode           = 50000
	StatusCode_OK            = 200
	StatusCode_UserNotFound  = baseStatusCode + 1
	StatusCode_WrongPassword = baseStatusCode + 2
	StatusMsg_LoginOK        = "登陆成功"
	StatusMsg_UserNotFound   = "该用户不存在"
	StatusMsg_WrongPasswrod  = "密码错误"
	StatusMsg_RegisterOK     = "注册成功"
	StatusMsg_UserExist      = "该用户已存在"
)

type userInfo struct {
	UserName string `form:"username" json:"username" binding:"required"`
	UserID   string `form:"user_id" json:"user_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetInformation(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.String(http.StatusOK, "TODO: get user information")
}

func Check(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.String(http.StatusOK, "TODO: check user status")
}

func Register(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var user userInfo
	if err := ctx.ShouldBind(&user); err != nil {
		log.Printf("[UserRegister] invalid params: %v", ctx.Params)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}

	// 检查用户是否已经存在
	exist, err := dao.GetUser(ctx, user.UserName, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg": "服务器内部错误",
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": StatusMsg_UserExist,
			"status_msg": "该用户已存在",
		})
		return
	}

	userModel := &model.User{
		UserName:     user.UserName,
		UserID:       user.UserID,
		Password:     user.Password,
		RegisterTime: time.Now(),
		LastLogin:    time.Now(), // FIXME: register and login?
	}
	if err := dao.RegisterUser(ctx, userModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  "服务器内部错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode_OK,
		"status_msg":  StatusMsg_RegisterOK,
	})
}

func Login(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")

	log.Printf("[UserLogin] userName: %q, password: %q\n", userName, password)

	userModel := dao.GetUserByUserName(ctx, userName)
	if userModel == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": StatusCode_UserNotFound,
			"status_msg":  StatusMsg_UserNotFound,
		})
		return
	}

	if password != userModel.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": StatusCode_WrongPassword,
			"status_msg":  StatusMsg_WrongPasswrod,
		})
		return
	}

	_ = dao.UpdateUserLastLoginTime(ctx, userModel.UserID, time.Now())

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode_OK,
		"status_msg":  StatusMsg_LoginOK,
	})
}
