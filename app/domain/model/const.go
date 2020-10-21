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

	// BalanceType

	BalanceTypeDebit  = 11
	BalanceTypeKredit = 22
)
