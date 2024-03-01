package model

type UserRegisterPayload struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Fullname        string `json:"fullname"`
}

type UserLoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Fullname string
}
