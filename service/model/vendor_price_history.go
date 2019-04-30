package model

import "time"

type VendorPriceHistory struct {
	BatchID    string    `gorm:"column:batch_id" json:"batch_id"`
	GoodsName  string    `gorm:"column:goods_name" json:"goods_name"`
	ID         int64     `gorm:"column:id" json:"id"`
	UnitPrice  float32   `gorm:"column:unit_price" json:"unit_price"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	Vendor     string    `gorm:"column:vendor" json:"vendor"`
}

// TableName sets the insert table name for this struct type
func (v *VendorPriceHistory) TableName() string {
	return "vendor_price_history"
}
