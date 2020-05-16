package user

import (
	"log"
	"net/http"
	"time"

	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

const (
	accessRightOperator_Grant  = 1
	accessRightOperator_Revoke = 2
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
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	// 检查用户是否已经存在
	exist, err := dao.GetUser(ctx, user.UserName, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_UserExist,
			"status_msg":  constant.StatusMsg_UserExist,
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
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constant.StatusCode_OK,
		"status_msg":  constant.StatusMsg_RegisterOK,
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
			"status_code": constant.StatusCode_UserNotFound,
			"status_msg":  constant.StatusMsg_UserNotFound,
		})
		return
	}

	if password != userModel.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_WrongPassword,
			"status_msg":  constant.StatusMsg_WrongPassword,
		})
		return
	}

	_ = dao.UpdateUserLastLoginTime(ctx, userModel.UserID, time.Now())

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constant.StatusCode_OK,
		"status_msg":  constant.StatusMsg_LoginOK,
	})
}

func GrantAccessRight(ctx *gin.Context) {
	operateAccessRight(ctx, accessRightOperator_Grant)
}

func RevokeAccessRight(ctx *gin.Context) {
	operateAccessRight(ctx, accessRightOperator_Revoke)
}

func operateAccessRight(ctx *gin.Context, operator int32) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	userId, found := ctx.GetQuery("user_id")
	if !found {
		log.Printf("[operateAccessRight] invalid params: %v", ctx.Params)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	var err error
	if operator == accessRightOperator_Grant {
		err = dao.UpdateUserStatus(ctx, userId, model.UserStatus_Active)
	} else {
		err = dao.UpdateUserStatus(ctx, userId, model.UserStatus_Inactive)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constant.StatusCode_OK,
		"status_msg":  constant.StatusMsg_OK,
	})
}

func Check(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": constant.StatusCode_MethodONotImplemented,
		"status_msg":  constant.StatusMsg_MethodNotImplemented,
	})
}
