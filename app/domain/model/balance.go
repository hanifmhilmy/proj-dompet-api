package model

type (
	// Balance is a model that use to create/add/edit balance account property
	Balance struct {
		ID         int64   `json:"id"`
		UID        int64   `json:"uid"`
		Name       string  `json:"name" validate:"required,min=3,max=50,alphanum"`
		Color      string  `json:"color"`
		Status     int     `json:"status"`
		Values     float64 `json:"values" validate:"required,numeric"`
		CurrencyID int32   `json:"currency_id"`
	}

	// BalanceData is a model that use to create/add/edit balance account property
	BalanceData struct {
		ID         int64   `json:"id" db:"balance_id"`
		Name       string  `json:"name" db:"name"`
		Color      string  `json:"color" db:"color"`
		Status     int     `json:"status" db:"status"`
		Values     float64 `json:"values" db:"last_values"`
		CurrencyID int32   `json:"currency_id" db:"currency_id"`
	}
)
