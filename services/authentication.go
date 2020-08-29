package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/twinj/uuid"

	"go-jwt-todo/helpers"
)

type AuthenticationService struct {
	Redis *redis.Client
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint
}

func (s *AuthenticationService) CreateAuth(userId uint) (*TokenDetails, error) {
	td, buildErr := s.buildToken(userId)
	if buildErr != nil {
		return nil, buildErr
	}

	saveErr := s.saveToken(userId, td)
	if saveErr != nil {
		return nil, saveErr
	}

	return td, nil
}

func (s *AuthenticationService) buildToken(userId uint) (*TokenDetails, error) {
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	// access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// refresh token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthenticationService) saveToken(userId uint, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := s.Redis.Set(td.AccessUuid, strconv.Itoa(int(userId)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := s.Redis.Set(td.RefreshUuid, strconv.Itoa(int(userId)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func (s *AuthenticationService) DeleteAuth(uuid string) error {
	_, err := s.Redis.Del(uuid).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthenticationService) VerifyToken(token string) (*AccessDetails, error) {

	jwtToken, err := helpers.ParseJwt(token, os.Getenv("ACCESS_SECRET"))
	if err != nil {
		return nil, err
	}

	claims, err := helpers.MapClaims(jwtToken)
	if err != nil {
		return nil, err
	}

	accessUuid, ok := claims["access_uuid"].(string)
	if !ok {
		return nil, fmt.Errorf("Access uuid invalid")
	}

	accessDetails, err := s.fetchAccessDetails(accessUuid)
	if err != nil {
		return nil, err
	}

	return accessDetails, nil
}

func (s *AuthenticationService) fetchAccessDetails(accessUuid string) (*AccessDetails, error) {
	result, err := s.Redis.Get(accessUuid).Result()
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return nil, err
	}
	return &AccessDetails{
		AccessUuid: accessUuid,
		UserId:     uint(userId),
	}, nil
}

func (s *AuthenticationService) RefreshAuth(refreshToken string) (*TokenDetails, error) {
	jwtToken, err := helpers.ParseJwt(refreshToken, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		return nil, err
	}

	claims, err := helpers.MapClaims(jwtToken)
	if err != nil {
		return nil, err
	}

	refreshUuid, ok := claims["refresh_uuid"].(string)
	if !ok {
		return nil, fmt.Errorf("Incorrect refresh token")
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return nil, err
	}

	if err := s.DeleteAuth(refreshUuid); err != nil {
		return nil, fmt.Errorf("Unable to remove refresh token")
	}

	tokenDetails, err := s.CreateAuth(uint(userId))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}
