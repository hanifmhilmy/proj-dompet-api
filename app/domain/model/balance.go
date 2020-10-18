package model

type (
	// Balance is a model that use to create/add/edit balance account property
	Balance struct {
		UID    int64
		Name   string
		Color  string
		Values float64
	}
)
