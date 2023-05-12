package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Products struct {
	ID         string `json:"ID" gorm:"primary_key"`
	Name       string `json:"Name"`
	Price      int32  `json:"Price"`
	Quantity   int8   `json:"Quantity"`
	IsDedele   bool   `json:"IsDedele" gorm:"default:false"`
	Details    string `json:"Details"`
	gorm.Model        //TODO: implement some fields by default of gorm
}

func CreateProduct(db *gorm.DB, product *Products) (err error) {
	fmt.Println(product)
	err = db.Create(product).Error
	if err != nil {
		return err
	}

	return nil
}

func GetProducts(db *gorm.DB, product *[]Products) (err error) {
	err = db.Find(product).Error
	if err != nil {
		return err
	}

	return nil
}

func GetProductById(db *gorm.DB, product *Products, id string) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}

	return nil
}
