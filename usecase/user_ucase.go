package usecase

import (
	"context"
	"errors"
	"expense-manager-golang/domain"
	"expense-manager-golang/pkg/jwt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{userRepo: ur}
}

func (uc *UserUsecase) Register(ctx context.Context, user *domain.User) error {
	//Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//change plain text password with hashed password
	user.Password = string(hashedPassword)

	//save to database
	return uc.userRepo.Create(ctx, user)
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	//find user by email
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Wrong email or password!")
	}

	//compare submitted password with hash password in database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Wrong email or password!")
	}

	//generate jwt token if credentials are valid
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.GenerateToken(user.ID, secret)
	if err != nil {
		return "", errors.New("Fail to create token")
	}

	return token, nil
}
