package model

const (
	UserStatusDeleted = 0
	UserStatusActive  = 1
	UserStatusPending = 2

	ActionBySystem = 0

	UserNotFound = 0

	LoggedInSuccess  = "Success login"
	LoggedOutSuccess = "Success logout"
	SignUpSuccess    = "User Created!"
	UserUnauthorized = "Unauthorized"

	RedisResetPassKey = "pwd:t_%d"

	InvalidCategory = 0
	ParentCategory  = 0

	StatusCategoryActive  = 1
	StatusCategoryDeleted = -1
	BalanceStatusActive   = 1

	QueryCategoryList = `SELECT category_id, value, parent_id, create_time, create_by, update_time, update_by from category where status = 1 and parent_id = $1`

	// QuerySaveBalance is the query for create the acc balance
	QuerySaveBalance = `INSERT INTO balance(user_id, last_values, name, color, status, create_time, create_by, update_time, update_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
)
