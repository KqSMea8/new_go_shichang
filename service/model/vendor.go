package model

import "github.com/guregu/null"

type Vendor struct {
	ContactName     null.String `gorm:"column:contact_name" json:"contact_name"`
	CostAdvance     int64       `gorm:"column:cost_advance" json:"cost_advance"`
	CostAlert       int64       `gorm:"column:cost_alert" json:"cost_alert"`
	ID              int64       `gorm:"column:id" json:"id"`
	IsValid         string      `gorm:"column:is_valid" json:"is_valid"`
	NeedCostAdvance string      `gorm:"column:need_cost_advance" json:"need_cost_advance"`
	Tel             null.String `gorm:"column:tel" json:"tel"`
	VendorName      string      `gorm:"column:vendor_name" json:"vendor_name"`
}

// TableName sets the insert table name for this struct type
func (v *Vendor) TableName() string {
	return "vendor"
}
