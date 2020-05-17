package dao

import (
	"context"
	"log"

	"github.com/BalusChen/IKHNAIE_API/model"
)

func AddProduct(ctx context.Context, product *model.Product) error {
	err := ikhnaieDB.Where("qsid = (?)", product.QSID).Assign(product.ToMap()).FirstOrCreate(product).Error
	if err != nil {
		log.Printf("[RegisterUser] insert to db failed, err: %v\n", err)
		return err
	}
	return nil
}

func GetProductByID(ctx context.Context, foodID int64) (*model.Product, error) {
	product := &model.Product{}
	err := ikhnaieDB.Where("id = (?)", foodID).Find(product).Error
	if err != nil {
		log.Printf("[GetProductsByID] select from db failed, err: %v", err)
		return nil, err
	}
	return product, nil
}

func GetProductsByUserID(ctx context.Context, ownerID string) ([]model.Product, error) {
	var products []model.Product
	err := ikhnaieDB.Where("owner_id = (?)", ownerID).Find(&products).Error
	if err != nil {
		log.Printf("[GetProductsByUserID] select from db failed, err: %v", err)
		return nil, err
	}
	return products, nil
}
