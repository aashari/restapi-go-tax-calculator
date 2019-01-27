package models

import (
	"github.com/jinzhu/gorm"
)

type Taxes struct {
	gorm.Model
	Code           string `gorm:"column:code" json:"code"`
	Name           string `gorm:"column:name" json:"name"`
	IsRefundable   bool   `gorm:"column:is_refundable" json:"is_refundable"`
	Configurations []TaxConfigurations
}

func (Taxes) TableName() string {
	return "taxes"
}
