package model

import (
	"errors"
	"time"
)

type (
	// AccountData exposed account data struct
	AccountData struct {
		ID          int64     `db:"user_id" json:"id"`
		Email       string    `db:"email" json:"email"`
		Name        string    `db:"name" json:"name"`
		CreatedTime time.Time `db:"create_time" json:"created_time"`
		CreatedBy   int64     `db:"create_by" json:"created_by"`
		UpdateTime  time.Time `db:"update_time" json:"update_time"`
		UpdateBy    int64     `db:"update_by" json:"update_by"`
	}

	// LoginDetails exposed login detail struct
	LoginDetails struct {
		Username string `json:"username" validate:"required,email,max=100"`
		Password string `json:"password" validate:"required,min=6,max=20"`
	}

	// SignUpDetails exposed signup model struct
	SignUpDetails struct {
		Email    string `db:"email" json:"email" validate:"required,email,max=100"`
		Name     string `db:"name" json:"name" validate:"min=4,max=100"`
		Username string `db:"username" json:"username" validate:"required,email,max=100"`
		Password string `db:"password" json:"password" validate:"required,min=6,max=20"`
	}

	// AccessDetails detail of extracted token data
	AccessDetails struct {
		AccUUID     string
		RefreshUUID string
		UserID      int64
	}
)

var (
	ErrNotLogin = errors.New("error user not login")
)
