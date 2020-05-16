package product

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/BalusChen/IKHNAIE_API/constant"
	"github.com/BalusChen/IKHNAIE_API/dao"
	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/gin-gonic/gin"
)

const (
	imageUrlTpl = "http://localhost:9877/ikhnaie/v1/assets/images/product/%d.jpg"
)

type productInfo struct {
	Name          string    `form:"name" binding:"required" json:"name"`                                               // 农产品名
	OwnerID       string    `form:"owner_id" binding:"required" json:"owner_id"`                                       // 所属人 ID
	Specification string    `form:"specification" json:"specification"`                                                // 规格
	Region        string    `form:"region" binding:"required" json:"region"`                                           // 产地
	MFRSName      string    `form:"mfrs_name" binding:"required" json:"mfrs_name"`                                     // 生产商名
	MFGDate       time.Time `form:"mfg_date" time_format:"2006-01-02" time_utc:"1" binding:"required" json:"mfg_date"` // 生产日期
	EXPDate       time.Time `form:"exp_date" time_format:"2006-01-02" time_utc:"1" binding:"required" json:"exp_date"` // 保质期
	QSID          string    `form:"qsid" binding:"required" json:"qsid"`                                               // 生产许可证编号
	LOT           string    `form:"lot" binding:"required" json:"lot"`                                                 // 生产批次号
	Description   string    `form:"description" json:"description"`                                                    // 产品描述
	ImageUrl      string    `form:"image_url" json:"image_url"`                                                                // 图片
}

func GetInformation(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: get product information")
}

func List(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ownerID, found := ctx.GetQuery("owner_id")
	if !found {
		log.Printf("[ProductList] invalid params: %v", ctx.Params)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	productModels, err := dao.GetProductsByUserID(ctx, ownerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}

	// FIXME: bad code.
	products := make([]productInfo, len(productModels))
	for i := 0; i < len(productModels); i++ {
		products[i].Name = productModels[i].Name
		products[i].OwnerID = productModels[i].OwnerID
		products[i].Specification = productModels[i].Specification
		products[i].Region = productModels[i].Region
		products[i].MFRSName = productModels[i].MFRSName
		products[i].MFGDate = productModels[i].MFGDate
		products[i].EXPDate = productModels[i].EXPDate
		products[i].QSID = productModels[i].QSID
		products[i].LOT = productModels[i].LOT
		products[i].Description = productModels[i].Description
		products[i].ImageUrl = fmt.Sprintf(imageUrlTpl, productModels[i].ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products":    products,
		"status_code": constant.StatusCode_OK,
		"status_msg":  constant.StatusMsg_OK,
	})
}

func AddProduct(ctx *gin.Context) {
	// 允许跨域
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var product productInfo
	if err := ctx.ShouldBind(&product); err != nil {
		log.Printf("[AddProduct] invalid params, err: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	_, header, err := ctx.Request.FormFile("image")
	if err != nil {
		log.Printf("[AddProduct] read image failed, err: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": constant.StatusCode_InvalidParams,
			"status_msg":  constant.StatusMsg_InvalidParams,
		})
		return
	}

	productModel := &model.Product{
		Name:          product.Name,
		OwnerID:       product.OwnerID,
		Specification: product.Specification,
		Region:        product.Region,
		MFRSName:      product.MFRSName,
		MFGDate:       product.MFGDate,
		EXPDate:       product.EXPDate,
		QSID:          product.QSID,
		LOT:           product.LOT,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	if err := dao.AddProduct(ctx, productModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"status_msg":  constant.StatusMsg_ServerInternalError,
		})
		return
	}

	err = ctx.SaveUploadedFile(header, fmt.Sprintf("./assets/images/product/%d%s", productModel.ID, path.Ext(header.Filename)))
	if err != nil {
		log.Printf("[AddProduct] save uploaded file failed, err: %v", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  constant.StatusMsg_OK,
	})
}
