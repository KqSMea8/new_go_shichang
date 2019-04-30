package model

import "time"

type ReportNotify struct {
	DateStr    string    `gorm:"column:date_str" json:"date_str"`
	HandleFlag string    `gorm:"column:handle_flag" json:"handle_flag"`
	ID         int64     `gorm:"column:id" json:"id"`
	WeightDate time.Time `gorm:"column:weight_date" json:"weight_date"`
}

// TableName sets the insert table name for this struct type
func (r *ReportNotify) TableName() string {
	return "report_notify"
}
