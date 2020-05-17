package model

import "time"

type Product struct {
	ID            int64     `gorm:"primary_key" json:"id"`                     // 农产品 ID（唯一标识）
	Name          string    `gorm:"column:name" json:"name"`                   // 农产品名
	OwnerID       string    `gorm:"column:owner_id" json:"owner_id"`           // 所属人 ID
	Specification string    `gorm:"column:specification" json:"specification"` // 规格
	Region        string    `gorm:"column:region" json:"region"`               // 产地
	MFRSName      string    `gorm:"column:mfrs_name" json:"mfrs_name"`         // 生产商名
	MFGDate       time.Time `gorm:"column:mfg_date" json:"mfg_date"`           // 生产日期
	EXPDate       time.Time `gorm:"column:exp_date" json:"exp_date"`           // 保质期
	QSID          string    `gorm:"column:qsid" json:"qsid"`                   // 生产许可证编号
	LOT           string    `gorm:"column:lot" json:"lot"`                     // 生产批次号
	Description   string    `gorm:"column:description" json:"description"`     // 产品描述
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`     // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`     // 更新时间
}

func (*Product) TableName() string {
	return "Product"
}

func (product *Product) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            product.ID,
		"name":          product.Name,
		"owner_id":      product.OwnerID,
		"specification": product.Specification,
		"region":        product.Region,
		"mfrs_name":     product.MFRSName,
		"mfg_date":      product.MFGDate,
		"exp_date":      product.EXPDate,
		"qsid":          product.QSID,
		"lot":           product.LOT,
		"description":   product.Description,
		"create_time":   product.CreateTime,
		"update_time":   product.UpdateTime,
	}
}
