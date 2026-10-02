package domain

import "context"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (string, error)
}
