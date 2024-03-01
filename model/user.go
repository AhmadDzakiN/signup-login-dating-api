package model

type UserRegisterPayload struct {
	Username        string `json:"username" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Fullname        string `json:"fullname" binding:"required"`
	Gender          string `json:"gender" binding:"required"`
	Location        string `json:"location" binding:"required"`
	Age             int    `json:"age" binding:"required"`
}

type UserLoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Fullname string `gorm:"not null"`
	Gender   string `gorm:"not null"`
	Location string `gorm:"not null"`
	Age      int    `gorm:"not null"`
}

type GetUserResponse struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Location string `json:"location"`
	Age      int    `json:"age"`
}
