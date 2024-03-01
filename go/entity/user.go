package entity

type User struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Email        string `gorm:"column:email;not null" json:"email"`
	Password     string `gorm:"column:password" json:"password"`
	NamaPengguna string `gorm:"column:nama_pengguna" json:"nama_pengguna"`
	UUID         string `gorm:"column:uuid" json:"uuid"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
