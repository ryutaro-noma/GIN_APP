package model

type User struct {
	UserId            string `gorm:"unique;not null"`
	UserInformationId string `gorm:"not null"`
	UserName          string `gorm:"not null"`
	Email             string `gorm:"not null"`
}
