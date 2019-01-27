package models

import (
	"github.com/jinzhu/gorm"
)

type TaxConfigurations struct {
	gorm.Model
	Tax      Taxes
	TaxId    int64    `gorm:"column:tax_id" json:"tax_id"`
	MinPrice *float64 `gorm:"column:min_price" json:"min_price"`
	MaxPrice *float64 `gorm:"column:max_price" json:"max_price"`
	Type     string   `gorm:"column:type" json:"type"`
	Value    float64  `gorm:"column:value" json:"value"`
	Priority int64    `gorm:"column:priority" json:"priority"`
}

func (TaxConfigurations) TableName() string {
	return "tax_configurations"
}
