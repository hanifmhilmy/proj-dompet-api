package model

import "time"

type (
	Account struct {
		id          int64
		email       string
		name        string
		createdTime time.Time
		createdBy   int64
		updateTime  time.Time
		updateBy    int64
	}

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
)

const (
	UserStatusDeleted = 0
	UserStatusActive  = 1
	UserStatusPending = 2

	UserActionBySystem = 0
)

// NewUser create new user account data
func NewUser(ac AccountData) *Account {
	return &Account{
		id:          ac.ID,
		email:       ac.Email,
		name:        ac.Name,
		createdTime: ac.CreatedTime,
		createdBy:   ac.CreatedBy,
		updateTime:  ac.UpdateTime,
		updateBy:    ac.UpdateBy,
	}
}

// GetIdentifier get account user id
func (a *Account) GetIdentifier() int64 {
	return a.id
}