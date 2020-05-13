package model

type Product struct {
	ID int64 `gorm:"id"`	// 自增主键
	ProductName string `gorm:"product_name"`	// 产品名
}
