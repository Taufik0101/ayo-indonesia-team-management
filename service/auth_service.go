package service

import (
	"errors"
	"fmt"
	"gin-ayo/config"
	"gin-ayo/database/models"
	"gin-ayo/dto"
	"gin-ayo/pkg/utils"
	repositories "gin-ayo/repository"
	"strings"
)

type (
	LoginResponse struct {
		AccessToken *string      `json:"access_token"`
		User        *models.User `json:"user"`
	}
)

type AuthService interface {
	Login(input dto.LoginUser) (*LoginResponse, error)
}

type authService struct {
	userRepository repositories.UserRepositoryInterface
}

func (a authService) Login(input dto.LoginUser) (*LoginResponse, error) {
	findUser, err := a.userRepository.FindOne(
		map[string]any{
			"LOWER(email) = ?": strings.ToLower(input.Email),
		},
		nil,
		nil,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if findUser == nil {
		return nil, errors.New("user not found")
	}

	err = findUser.ComparePasswords(input.Password)

	if err != nil {
		return nil, errors.New("incorrect credential")
	}

	accessTokenPayload := dto.AccessTokenPayload{
		UserId:   findUser.ID.String(),
		UserType: string(findUser.Role),
	}

	jwtAccessTokenExpiration := config.GetEnv("JWT_ACCESS_TOKEN_EXPIRATION", "24h")
	jwtSecret := config.GetEnv("JWT_SECRET", "")
	accessToken, err := utils.JwtGenerate(accessTokenPayload, jwtAccessTokenExpiration, jwtSecret)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to generate accessToken due to %s", err.Error()))
	}

	return &LoginResponse{
		AccessToken: accessToken,
		User:        findUser,
	}, nil
}

func NewAuthService(userRepository repositories.UserRepositoryInterface) AuthService {
	return &authService{userRepository: userRepository}
}
