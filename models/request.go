package models

type Users struct {
	Id       string `json:"id"       gorm:"column:id;primaryKey"`
	Name     string `json:"name"     gorm:"column:name"`
	Email    string `json:"email"    gorm:"column:email;uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
}

func (Users) TableName() string { return "users" }

type Login struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
