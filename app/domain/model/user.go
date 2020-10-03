package model

type (
	User struct {
		ID    int64  `db:"user_id"`
		Email string `db:"email"`
	}
)
