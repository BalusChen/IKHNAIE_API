package user

import (
	"log"
	"net/http"
	"time"

	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

type userInfo struct {
	UserName     string `form:"username" json:"username" binding:"required"`
	UserID       string `form:"user_id" json:"user_id" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Organization string `form:"organization" json:"organization" binding:"required"`
}

func Register(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var user userInfo
	// 检查参数是否正确
	if err := ctx.ShouldBind(&user); err != nil {
		log.Printf("[UserRegister] invalid params: %v", ctx.Params)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  StatusMsg_BadRequest,
		})
		return
	}

	// 检查用户是否已经存在
	exist, err := dao.GetUser(ctx, user.UserName, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  StatusMsg_ServerInternalError,
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": StatusCode_UserExist,
			"status_msg":  StatusMsg_UserExist,
		})
		return
	}

	// 将用户数据存入数据库
	userModel := &model.User{
		Type:         model.UserType_Normal,
		Status:       model.UserStatus_Inactive,
		UserName:     user.UserName,
		UserID:       user.UserID,
		Password:     user.Password,
		Organization: user.Organization,
		RegisterTime: time.Now(),
		LastLogin:    time.Now(), // FIXME: register and login?
	}
	if err := dao.RegisterUser(ctx, userModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  StatusMsg_ServerInternalError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode_OK,
		"status_msg":  StatusMsg_RegisterOK,
	})
}

func Login(ctx *gin.Context) {
	// 允许跨域
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
			"status_msg":  StatusMsg_WrongPassword,
		})
		return
	}

	_ = dao.UpdateUserLastLoginTime(ctx, userModel.UserID, time.Now())

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode_OK,
		"status_msg":  StatusMsg_LoginOK,
	})
}

func Check(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": StatusCode_MethodONotImplemented,
		"status_msg":  StatusMsg_MethodNotImplemented,
	})
}
