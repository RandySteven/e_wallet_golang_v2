package entities

type QueryCondition struct {
	SortedBy  string `form:"sortBy,default=date"`
	Sort      string `form:"sort,default=asc"`
	Limit     string `form:"limit,default=10"`
	Page      string `form:"page"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
