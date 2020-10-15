package model

import "time"

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
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// SignUpDetails exposed signup model struct
	SignUpDetails struct {
		Email    string `db:"email" json:"email"`
		Name     string `db:"name" json:"name"`
		Username string `db:"username" json:"username"`
		Password string `db:"password" json:"password"`
	}

	// AccessDetails detail of extracted token data
	AccessDetails struct {
		AccUUID     string
		RefreshUUID string
		UserID      int64
	}
)

const (
	UserStatusDeleted = 0
	UserStatusActive  = 1
	UserStatusPending = 2

	UserActionBySystem = 0

	UserNotFound = 0

	LoggedInSuccess  = "Success login"
	LoggedOutSuccess = "Success logout"
	SignUpSuccess    = "User Created!"
	UserUnauthorized = "Unauthorized"

	RedisResetPassKey = "pwd:t_%d"
)
