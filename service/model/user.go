package model

type User struct {
	ID           int64  `gorm:"column:id" json:"id"`
	IsSupperUser string `gorm:"column:is_supper_user" json:"is_supper_user"`
	IsValid      string `gorm:"column:is_valid" json:"is_valid"`
	Password     string `gorm:"column:password" json:"password"`
	UserName     string `gorm:"column:user_name" json:"user_name"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
