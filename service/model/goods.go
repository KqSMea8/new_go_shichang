package model

import "github.com/guregu/null"

type Goods struct {
	GoodsAlias null.String `gorm:"column:goods_alias" json:"goods_alias"`
	GoodsName  null.String `gorm:"column:goods_name" json:"goods_name"`
	GroupName  null.String `gorm:"column:group_name" json:"group_name"`
	ID         int64       `gorm:"column:id" json:"id"`
	OrderIndex null.Int    `gorm:"column:order_index" json:"order_index"`
}

// TableName sets the insert table name for this struct type
func (g *Goods) TableName() string {
	return "goods"
}
