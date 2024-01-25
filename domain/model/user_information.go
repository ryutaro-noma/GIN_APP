package model

type UserInformation struct {
	Id        int    `gorm:"unique;not null"`
	LastName  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	Sex       string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	PostCode  string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Remarks   string `gorm:"not null"`
	User      User   `gorm:"not null"`
}
