package model

import (
	"github.com/guregu/null"
	"time"
)

type BasicDailyData struct {
	BatchID         string      `gorm:"column:batch_id" json:"batch_id"`
	CarNumber       string      `gorm:"column:car_number" json:"car_number"`
	DayStr          string      `gorm:"column:day_str" json:"day_str"`
	DriverName      null.String `gorm:"column:driver_name" json:"driver_name"`
	GrossWeight     null.String `gorm:"column:gross_weight" json:"gross_weight"`
	GroupName       string      `gorm:"column:group_name" json:"group_name"`
	ID              int64       `gorm:"column:id" json:"id"`
	ImportDate      time.Time   `gorm:"column:import_date" json:"import_date"`
	NetWeight       string      `gorm:"column:net_weight" json:"net_weight"`
	NumberFlag      string      `gorm:"column:number_flag" json:"number_flag"`
	RowIndex        null.Int    `gorm:"column:row_index" json:"row_index"`
	ShopNumber      string      `gorm:"column:shop_number" json:"shop_number"`
	TareWeight      null.String `gorm:"column:tare_weight" json:"tare_weight"`
	TimeStr         string      `gorm:"column:time_str" json:"time_str"`
	Vendor          string      `gorm:"column:vendor" json:"vendor"`
	VendorNetWeight null.String `gorm:"column:vendor_net_weight" json:"vendor_net_weight"`
	VendorSubName   null.String `gorm:"column:vendor_sub_name" json:"vendor_sub_name"`
	WeightDate      time.Time   `gorm:"column:weight_date" json:"weight_date"`
}

// TableName sets the insert table name for this struct type
func (b *BasicDailyData) TableName() string {
	return "basic_daily_data"
}
