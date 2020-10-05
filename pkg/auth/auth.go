package auth

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type (
	AuthInterface interface {
		CreateToken(uid int64) (token *Token, err error)
		VerifyToken(tokenString string) (*jwt.Token, error)
		ExtractTokenMetadata(token *jwt.Token) (*accessDetails, error)
	}

	auth struct {
		as  string
		rs  string
		opt Options
	}

	accessDetails struct {
		uuid   string
		userID int64
	}

	Options struct {
		AccessExpire  int64
		RefreshExpire int64
	}

	Token struct {
		accessToken   string
		refreshToken  string
		uuidAccess    string
		uuidRefresh   string
		expireAccess  int64
		expireRefresh int64
	}
)

const (
	prefix = "[Auth Package] "
)

func NewAuth(accessSecret, refreshSecret string, opt Options) AuthInterface {
	return &auth{
		as:  accessSecret,
		rs:  refreshSecret,
		opt: opt,
	}
}

func (a *auth) CreateToken(uid int64) (token *Token, err error) {
	uidAccess := uuid.NewV4().String()
	token = &Token{
		uuidAccess:    uidAccess,
		uuidRefresh:   uidAccess + "++" + strconv.Itoa(int(uid)),
		expireAccess:  time.Now().Add(time.Minute * time.Duration(a.opt.AccessExpire)).Unix(),
		expireRefresh: time.Now().Add(time.Hour * 24 * time.Duration(a.opt.RefreshExpire)).Unix(),
	}
	accessTokenClaims := jwt.MapClaims{
		"authorized":  true,
		"access_uuid": uidAccess,
		"user_id":     uid,
		"exp":         token.expireAccess,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	token.accessToken, err = at.SignedString(a.as)
	if err != nil {
		log.Println(prefix, "Fail at creating access token, err -> ", err)
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"refresh_uuid": token.uuidRefresh,
		"user_id":      uid,
		"exp":          token.expireRefresh,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	token.refreshToken, err = rt.SignedString(a.rs)
	if err != nil {
		log.Println(prefix, "Fail at creating refresh token, err -> ", err)
		return nil, err
	}

	return token, nil
}

// Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
func (a *auth) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.as), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return nil, err
	}
	return token, nil
}

// Extract token metadata
func (a *auth) ExtractTokenMetadata(token *jwt.Token) (*accessDetails, error) {
	var err error

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &accessDetails{
			uuid:   accessUUID,
			userID: userID,
		}, nil
	}
	return nil, err
}

// Refresh token data
func (a *auth) Refresh(refreshToken string) error {
	return nil
}
