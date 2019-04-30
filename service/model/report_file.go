package model

import (
	"github.com/guregu/null"
	"time"
)

type ReportFile struct {
	BatchID    null.String `gorm:"column:batch_id" json:"batch_id"`
	DayStr     string      `gorm:"column:day_str" json:"day_str"`
	FileName   null.String `gorm:"column:file_name" json:"file_name"`
	ID         int64       `gorm:"column:id" json:"id"`
	ReportDate time.Time   `gorm:"column:report_date" json:"report_date"`
	ReportType string      `gorm:"column:report_type" json:"report_type"`
}

// TableName sets the insert table name for this struct type
func (r *ReportFile) TableName() string {
	return "report_file"
}
